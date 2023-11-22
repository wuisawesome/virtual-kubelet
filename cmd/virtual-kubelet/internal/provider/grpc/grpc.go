package grpc

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"google.golang.org/grpc"

	dto "github.com/prometheus/client_model/go"
	"github.com/virtual-kubelet/virtual-kubelet/log"
	"github.com/virtual-kubelet/virtual-kubelet/node/api"
	"github.com/virtual-kubelet/virtual-kubelet/node/api/statsv1alpha1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)


type GrpcProvider struct { //nolint:golint
	client PodProviderClient
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
type GrpcConfig struct { //nolint:golint
	ServerAddr        string            `json:"server_addr,omitempty"`
}

func newGrpcProvider(configPath *GrpcConfig) (*GrpcProvider, error) {
	grpcChannel, err := grpc.Dial(configPath.ServerAddr)
	if err != nil {
		return nil, err
	}
	return &GrpcProvider{
		client: NewContainerProviderClient(grpcChannel),
	}, nil
}

func loadGrpcConfig(configPath string) (config *GrpcConfig, err error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}
	configMap := map[string]GrpcConfig{}
	err = json.Unmarshal(data, &configMap)
	return config, err
}

func NewGrpcProvider(configPath string) (*GrpcProvider, error) {
	config, err := loadGrpcConfig(configPath)
	if err != nil {
		return nil, err
	}

	return newGrpcProvider(config)
}


func (*GrpcProvider) ConfigureNode(ctx context.Context, node *v1.Node) {
	logger := log.GetLogger(ctx)
	logger.Error("ConfigureNode")
	node.Status.Capacity = v1.ResourceList{
		"cpu":    resource.MustParse("99"),
		"memory": resource.MustParse("100Gi"),
		"pods":   resource.MustParse("1000"),
	}
}


// GetContainerLogs retrieves the logs of a container by name from the provider.
func (*GrpcProvider) GetContainerLogs(ctx context.Context, namespace, podName, containerName string, opts api.ContainerLogOpts) (io.ReadCloser, error) {
	logger := log.GetLogger(ctx)
	logger.Error("GetContainerLogs")
	return nil, nil
}

// RunInContainer executes a command in a container in the pod, copying data
// between in/out/err and the container's stdin/stdout/stderr.
func (*GrpcProvider) RunInContainer(ctx context.Context, namespace, podName, containerName string, cmd []string, attach api.AttachIO) error {
	logger := log.GetLogger(ctx)
	logger.Error("RunInContainer")
	return nil
}

// AttachToContainer attaches to the executing process of a container in the pod, copying data
// between in/out/err and the container's stdin/stdout/stderr.
func (*GrpcProvider) AttachToContainer(ctx context.Context, namespace, podName, containerName string, attach api.AttachIO) error {
	logger := log.GetLogger(ctx)
	logger.Error("RunInContainer")
	return nil
}

// GetStatsSummary gets the stats for the node, including running pods
func (*GrpcProvider) GetStatsSummary(ctx context.Context) (*statsv1alpha1.Summary, error) {
	logger := log.GetLogger(ctx)
	logger.Error("GetStatsSummary")
	return nil, nil
}

// GetMetricsResource gets the metrics for the node, including running pods
func (*GrpcProvider) GetMetricsResource(ctx context.Context) ([]*dto.MetricFamily, error) {
	logger := log.GetLogger(ctx)
	logger.Error("GetMetricsResource")
	return nil, nil
}

// PortForward forwards a local port to a port on the pod
func (*GrpcProvider) PortForward(ctx context.Context, namespace, pod string, port int32, stream io.ReadWriteCloser) error {
	logger := log.GetLogger(ctx)
	logger.Error("PortForward")
	return nil
}
