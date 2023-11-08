package grpc

import (
	"context"
	"io"

	dto "github.com/prometheus/client_model/go"
	"github.com/virtual-kubelet/virtual-kubelet/node/api"
	"github.com/virtual-kubelet/virtual-kubelet/node/api/statsv1alpha1"
	v1 "k8s.io/api/core/v1"
)


type GrpcProvider struct { //nolint:golint
	// nodeName           string
	// operatingSystem    string
	// internalIP         string
	// daemonEndpointPort int32
	// pods               map[string]*v1.Pod
	// config             MockConfig
	// startTime          time.Time
	// notifier           func(*v1.Pod)
}

// GrpcConfig contains a mock virtual-kubelet's configurable parameters.
// type GrpcConfig struct { //nolint:golint
// 	CPU        string            `json:"cpu,omitempty"`
// 	Memory     string            `json:"memory,omitempty"`
// 	Pods       string            `json:"pods,omitempty"`
// 	Others     map[string]string `json:"others,omitempty"`
// 	ProviderID string            `json:"providerID,omitempty"`
// }

func NewGrpcProvider(configPath string) (*GrpcProvider, error){
	return &GrpcProvider{}, nil
}


func (*GrpcProvider) ConfigureNode(context.Context, *v1.Node) {

}


// GetContainerLogs retrieves the logs of a container by name from the provider.
func (*GrpcProvider) GetContainerLogs(ctx context.Context, namespace, podName, containerName string, opts api.ContainerLogOpts) (io.ReadCloser, error) {
	return nil, nil
}

// RunInContainer executes a command in a container in the pod, copying data
// between in/out/err and the container's stdin/stdout/stderr.
func (*GrpcProvider) RunInContainer(ctx context.Context, namespace, podName, containerName string, cmd []string, attach api.AttachIO) error {
	return nil
}

// AttachToContainer attaches to the executing process of a container in the pod, copying data
// between in/out/err and the container's stdin/stdout/stderr.
func (*GrpcProvider) AttachToContainer(ctx context.Context, namespace, podName, containerName string, attach api.AttachIO) error {
	return nil
}

// GetStatsSummary gets the stats for the node, including running pods
func (*GrpcProvider) GetStatsSummary(context.Context) (*statsv1alpha1.Summary, error) {
	return nil, nil
}

// GetMetricsResource gets the metrics for the node, including running pods
func (*GrpcProvider) GetMetricsResource(context.Context) ([]*dto.MetricFamily, error) {
	return nil, nil
}

// PortForward forwards a local port to a port on the pod
func (*GrpcProvider) PortForward(ctx context.Context, namespace, pod string, port int32, stream io.ReadWriteCloser) error {

	return nil
}
