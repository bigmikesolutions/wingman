package k8s

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
)

func GetRestartCount(pod *v1.Pod) int {
	for _, containerStatus := range pod.Status.ContainerStatuses {
		if containerStatus.RestartCount > 0 {
			return int(containerStatus.RestartCount)
		}
	}
	return 0
}

func GetReadyInfo(pod *v1.Pod) string {
	ready := 0
	for _, containerStatus := range pod.Status.ContainerStatuses {
		if containerStatus.Ready {
			ready++
		}
	}
	return fmt.Sprintf("%d/%d", ready, len(pod.Status.ContainerStatuses))
}
