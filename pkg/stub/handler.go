package stub

import (
	"context"
	"time"

	"github.com/Percona-Lab/percona-server-mongodb-operator/pkg/apis/psmdb/v1alpha1"
	pkgSdk "github.com/Percona-Lab/percona-server-mongodb-operator/pkg/sdk"

	motPkg "github.com/percona/mongodb-orchestration-tools/pkg"
	podk8s "github.com/percona/mongodb-orchestration-tools/pkg/pod/k8s"
	watchdog "github.com/percona/mongodb-orchestration-tools/watchdog"
	wdConfig "github.com/percona/mongodb-orchestration-tools/watchdog/config"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
)

var ReplsetInitWait = 10 * time.Second

const minPersistentVolumeClaims = 1

func NewHandler(client pkgSdk.Client) sdk.Handler {
	return &Handler{
		client:       client,
		startedAt:    time.Now(),
		watchdogQuit: make(chan bool, 1),
	}
}

type Handler struct {
	client       pkgSdk.Client
	pods         *podk8s.Pods
	watchdog     *watchdog.Watchdog
	watchdogQuit chan bool
	startedAt    time.Time
}

func (h *Handler) ensureWatchdog(m *v1alpha1.PerconaServerMongoDB, usersSecret *corev1.Secret) error {
	if h.watchdog != nil {
		return nil
	}

	// Start the watchdog if it has not been started
	h.watchdog = watchdog.New(&wdConfig.Config{
		ServiceName:    m.Name,
		Username:       string(usersSecret.Data[motPkg.EnvMongoDBClusterAdminUser]),
		Password:       string(usersSecret.Data[motPkg.EnvMongoDBClusterAdminPassword]),
		APIPoll:        5 * time.Second,
		ReplsetPoll:    5 * time.Second,
		ReplsetTimeout: 3 * time.Second,
	}, &h.watchdogQuit, h.pods)
	go h.watchdog.Run()

	return nil
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.PerconaServerMongoDB:
		psmdb := o

		// Ignore the delete event since the garbage collector will clean up all secondary resources for the CR
		// All secondary resources must have the CR set as their OwnerReference for this to be the case
		if event.Deleted {
			logrus.Infof("Received deleted event for %s", psmdb.Name)
			close(h.watchdogQuit)
			h.watchdog = nil
			return nil
		}

		// Create the mongodb internal auth key if it doesn't exist
		err := h.client.Create(newPSMDBMongoKeySecret(o))
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				logrus.Errorf("failed to create psmdb auth key: %v", err)
				return err
			}
		} else {
			logrus.Info("created mongodb auth key secret")
		}

		// Load MongoDB system users/passwords from secret
		usersSecret, err := getPSMDBSecret(psmdb, h.client, psmdb.Spec.Secrets.Users)
		if err != nil {
			logrus.Errorf("failed to load psmdb user secrets: %v", err)
			return err
		}

		// Ensure all replica sets exist. When sharding is supported this
		// loop will create the cluster shards and config server replset
		for _, replset := range psmdb.Spec.Replsets {
			// Update the PSMDB status
			podList, err := h.updateStatus(psmdb, replset, usersSecret)
			if err != nil {
				logrus.Errorf("failed to update psmdb status for replset %s: %v", replset.Name, err)
				return err
			}

			// Ensure replset exists
			status, err := h.ensureReplset(psmdb, podList, replset, usersSecret)
			if err != nil {
				if err == ErrNoRunningMongodContainers {
					logrus.Debugf("no running mongod containers for replset %s, skipping replset initiation", replset.Name)
					continue
				}
				logrus.Errorf("failed to ensure replset %s: %v", replset.Name, err)
				return err
			}
			if !status.Initialized {
				continue
			}

			// Remove PVCs left-behind from scaling down
			err = persistentVolumeClaimReaper(psmdb, h.client, podList, replset, status)
			if err != nil {
				logrus.Errorf("failed to run persistent volume claim reaper for replset %s: %v", replset.Name, err)
				return err
			}
		}
	}
	return nil
}
