package k8s

import (
	"time"

	"github.com/bigmikesolutions/wingman/pkg/iam"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type PodInfo struct {
	Name     string        `json:"name"`
	Ready    string        `json:"ready"`
	Status   string        `json:"status"`
	Restarts int           `json:"restarts"`
	Age      time.Duration `json:"age"`
}

type ResourcePods struct {
	pods   *v1.PodList
	client *kubernetes.Clientset
}

func (p ResourcePods) GetID() iam.ResourceID {
	return ""
}

func (p ResourcePods) GetType() iam.ResourceType {
	return iam.ResourceType(p.pods.Kind)
}

func (p ResourcePods) GetProviderID() iam.ProviderID {
	return ProviderName
}

func (p ResourcePods) GetName() string {
	return ""
}

func (p ResourcePods) Execute(action iam.ResourceAction) error {
	// TODO implement
	return errors.New("not implemented")
}

func (p ResourcePods) GetInfo() interface{} {
	info := make([]PodInfo, 0)
	for _, pod := range p.pods.Items {
		info = append(info, PodInfo{
			Name:     pod.Name,
			Ready:    GetReadyInfo(&pod),
			Status:   string(pod.Status.Phase),
			Restarts: GetRestartCount(&pod),
			Age:      time.Now().Sub(pod.Status.StartTime.Time),
		})
	}
	return info
}
