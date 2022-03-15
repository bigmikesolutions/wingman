package k8s

import (
	"context"
	"flag"
	"path/filepath"

	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type Cluster struct {
	client *kubernetes.Clientset
}

func NewCluster() (*Cluster, error) {
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
	return &Cluster{
		client: client,
	}, nil
}

func (c *Cluster) GetPods(ctx context.Context, namespace string, opt metav1.ListOptions) (*v1.PodList, error) {
	return c.client.CoreV1().Pods(namespace).List(ctx, opt)
}
