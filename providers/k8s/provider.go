package k8s

import (
	"flag"
	"path/filepath"

	"github.com/bigmikesolutions/wingman/pkg/iam"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

const (
	ProviderName iam.ProviderID = "k8s"
)

type Provider struct {
	client *kubernetes.Clientset
}

func NewProvider() (iam.ResourceProvider, error) {
	var k8sCfgPath *string
	if home := homedir.HomeDir(); home != "" {
		k8sCfgPath = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		k8sCfgPath = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *k8sCfgPath)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't load k8s config")
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't create k8s client")
	}
	return &Provider{
		client: client,
	}, nil
}

func (c *Provider) Provide(req *iam.GetResourceRequest) (iam.Resource, error) {
	if len(req.Path) == 0 {
		return nil, errors.New("path is required")
	}
	switch req.Path[0] {
	case "pods":
		pods, err := c.client.CoreV1().Pods("default").List(req.Ctx, metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
		return ResourcePods{
			pods:   pods,
			client: c.client,
		}, nil
	}
	return nil, errors.Errorf("not implemented resource: %s", req.Path[0])
}
