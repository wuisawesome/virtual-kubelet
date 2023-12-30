package grpc

import (
	"context"
	"encoding/json"
	"io"
	"time"
	"os"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	dto "github.com/prometheus/client_model/go"
	"github.com/virtual-kubelet/virtual-kubelet/log"
	"github.com/virtual-kubelet/virtual-kubelet/node/api"
	"github.com/virtual-kubelet/virtual-kubelet/node/api/statsv1alpha1"
	v1 "k8s.io/api/core/v1"
)


type GrpcProvider struct { //nolint:golint
	client PodProviderClient
	// nodeName           string
	// operatingSystem    string
	// internalIP         string
	// daemonEndpointPort int32
	// config             MockConfig
	// startTime          time.Time
	// notifier           func(*v1.Pod)
}

// GrpcConfig contains a mock virtual-kubelet's configurable parameters.
type GrpcConfig struct { //nolint:golint
	ServerAddr        string            `json:"server_addr,omitempty"`
}

func newGrpcProvider(config GrpcConfig) (*GrpcProvider, error) {
	for i := 0; i < 5; i++ {
		grpcChannel, err := grpc.Dial(config.ServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err == nil {
			return &GrpcProvider{
				client: NewPodProviderClient(grpcChannel),
			}, nil
		}
		if i == 4 {
			return nil, err
		}
		time.Sleep(5 * time.Second)
	}
	return nil, nil
}

func loadGrpcConfig(configPath string) (config GrpcConfig, err error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	if err = json.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("Error unmarshalling string %s. %s", data, err)
	}

	return config, nil
}

func NewGrpcProvider(configPath string) (*GrpcProvider, error) {
	config, err := loadGrpcConfig(configPath)
	if err != nil {
		return nil, err
	}

	return newGrpcProvider(config)
}


func (provider *GrpcProvider) ConfigureNode(ctx context.Context, node *v1.Node) {
	logger := log.GetLogger(ctx)
	logger.Info("ConfigureNode")

	node_json, err := json.Marshal(node)
	if err != nil {
		logger.Fatalf("Couldn't marshal node %v. %s", node, err)
		return
	}

	request := ConfigureNodeRequest{
		CoreV1NodeJson: string(node_json),
	}
	reply, err := provider.client.ConfigureNode(ctx, &request)
	if err != nil {
		logger.Fatalf("Request failed %v. %s", request, err)
		return
	}

	err = json.Unmarshal([]byte(reply.GetCoreV1NodeJson()), node)
	if err != nil {
		logger.Fatalf("Couldn't deserialize reply %v. %s", request, err)
		return
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
