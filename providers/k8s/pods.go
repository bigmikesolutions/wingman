package k8s

import (
	"github.com/bigmikesolutions/wingman/pkg/iam"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

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
	return p.pods.Items
}
