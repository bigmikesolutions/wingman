package k8s

import (
	"flag"
	"path/filepath"

	"github.com/bigmikesolutions/wingman/pkg/cqrs"

	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type Provider struct {
	client *kubernetes.Clientset
}

func NewProvider(cfg *cqrs.Config) error {
	var k8sCfgPath *string
	if home := homedir.HomeDir(); home != "" {
		k8sCfgPath = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		k8sCfgPath = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *k8sCfgPath)
	if err != nil {
		return errors.Wrap(err, "couldn't load k8s config")
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return errors.Wrap(err, "couldn't create k8s client")
	}
	cfg.AddQueryHandlers(
		NewPodsQueryHandler(client),
	)
	return nil
}
