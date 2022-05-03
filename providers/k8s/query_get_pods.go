package k8s

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/bigmikesolutions/wingman/pkg/iam/access"

	"github.com/bigmikesolutions/wingman/providers/actions"

	v1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/bigmikesolutions/wingman/pkg/provider"

	"github.com/bigmikesolutions/wingman/pkg/cqrs"
	"k8s.io/client-go/kubernetes"
)

type K8sGetPodsQueryResult struct {
	Pods []*PodInfo
}

type PodInfo struct {
	Name     string        `json:"name"`
	Ready    string        `json:"ready"`
	Status   string        `json:"status"`
	Restarts int           `json:"restarts"`
	Age      time.Duration `json:"age"`
}

type PodsQueryHandler struct {
	client *kubernetes.Clientset
}

func NewPodsQueryHandler(client *kubernetes.Clientset) cqrs.QueryHandler {
	return &PodsQueryHandler{
		client: client,
	}
}

func (h PodsQueryHandler) GetType() cqrs.QueryType {
	return provider.GetProviderGetResourceQueryType(ProviderName, PodsResourceType)
}

func (h PodsQueryHandler) Handle(ctx context.Context, q cqrs.Query) (interface{}, error) {
	queryBus := cqrs.GetQueryBus(ctx)
	if queryBus == nil {
		return nil, errors.New("query bus not found in context")
	} else if noAccessErr := actions.HasResourceAccess(
		ctx,
		queryBus,
		ProviderName,
		PodsResourceType,
		nil,
		access.AccessTypeRead,
	); noAccessErr != nil {
		return nil, noAccessErr
	}
	query := q.(*provider.QueryGetResource)
	pods, err := h.client.CoreV1().
		Pods(GetNamespace(query.Query)).
		List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return newPodsQueryResult(pods), nil
}

func newPodsQueryResult(list *v1.PodList) *K8sGetPodsQueryResult {
	pods := make([]*PodInfo, 0)
	for _, pod := range list.Items {
		pods = append(pods, &PodInfo{
			Name:     pod.Name,
			Ready:    GetReadyInfo(&pod),
			Status:   string(pod.Status.Phase),
			Restarts: GetRestartCount(&pod),
			Age:      time.Now().Sub(pod.Status.StartTime.Time),
		})
	}
	return &K8sGetPodsQueryResult{
		Pods: pods,
	}
}
