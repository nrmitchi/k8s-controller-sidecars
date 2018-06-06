package main

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	core_v1 "k8s.io/api/core/v1"

	set "github.com/deckarep/golang-set"
)

// Handler interface contains the methods that are required
type Handler interface {
	Init() error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(objOld, objNew interface{})
}

// SidecarShutdownHandler is a sample implementation of Handler
type SidecarShutdownHandler struct{}

// Init handles any handler initialization
func (t *SidecarShutdownHandler) Init() error {
	log.Info("SidecarShutdownHandler.Init")
	return nil
}

// Send a shutdown signal to all containers in the Pod
func sendShutdownSignal(pod *core_v1.Pod) {
	log.Infof("It's going down, I'm yelling TIMBERRRRR for pod: %s", pod.Name)
}

// ObjectCreated is called when an object is created
func (t *SidecarShutdownHandler) ObjectCreated(obj interface{}) {
	log.Info("SidecarShutdownHandler.ObjectCreated")
	// assert the type to a Pod object to pull out relevant data
	pod := obj.(*core_v1.Pod)

	sidecarsString, exists := pod.Annotations["nrmitchi.com/sidecars"]

	if exists {
		log.Infof("    ResourceTrackable: true")
		log.Infof("    Sidecars: %s", sidecarsString)
	} else {
		log.Infof("    ResourceTrackable: false")

		return
	}

	sidecars := set.NewSet()

	for _, s := range strings.Split(sidecarsString, ",") {
		sidecars.Add(s)
	}

	log.Infof("    ResourceVersion: %s", pod.ObjectMeta.ResourceVersion)
	log.Infof("    NodeName: %s", pod.Spec.NodeName)
	log.Infof("    Phase: %s", pod.Status.Phase)

	allContainers := set.NewSet()
	runningContainers := set.NewSet()
	completedContainers := set.NewSet()

	for _, containerStatus := range pod.Status.ContainerStatuses {
		allContainers.Add(containerStatus.Name)

		if containerStatus.Ready {
			runningContainers.Add(containerStatus.Name)
		} else {
			if containerStatus.State.Terminated != nil && containerStatus.State.Terminated.Reason == "Completed" {
				completedContainers.Add(containerStatus.Name)
			}
		}
	}

	log.Infof("    all       : %s", allContainers)
	log.Infof("    running   : %s", runningContainers)
	log.Infof("    completed : %s", completedContainers)
	log.Infof("    sidecars  : %s", sidecars)

	// If we have accounted for all of the containers, and the sidecar containers are the only
	// ones still running, issue them each a shutdown command
	if runningContainers.Union(completedContainers).Equal(allContainers) {
		log.Infof("  We have all the containers")
		if runningContainers.Equal(sidecars) {
			log.Infof("    Sending shutdown signal to containers: %s, %s", pod.Name, sidecars)
			sendShutdownSignal(pod)
		}
	}
}

// ObjectDeleted is called when an object is deleted
func (t *SidecarShutdownHandler) ObjectDeleted(obj interface{}) {
	log.Info("SidecarShutdownHandler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated.
// Note that the controller in this repo will never call this function properly.
// It uses only ObjectCreated
func (t *SidecarShutdownHandler) ObjectUpdated(objOld, objNew interface{}) {
	log.Info("SidecarShutdownHandler.ObjectUpdated")
}
