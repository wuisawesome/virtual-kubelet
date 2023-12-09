package grpc

import (
	"context"
	"encoding/json"
	"fmt"


	"github.com/virtual-kubelet/virtual-kubelet/log"
	corev1 "k8s.io/api/core/v1"
)


// CreatePod takes a Kubernetes Pod and deploys it within the provider.
func (provider *GrpcProvider) CreatePod(ctx context.Context, pod *corev1.Pod) error {
	logger := log.GetLogger(ctx)
	logger.Info("CreatePod")

	key, err := buildKey(pod)
	if err != nil {
		return err
	}

	pod_json, err := json.Marshal(pod)
	if err != nil {
		logger.Errorf("Couldn't marshal pod %v. %s", pod, err)
		return err
	}

	request := CreatePodRequest{
		CoreV1PodJson: string(pod_json),
	}
	_, err = provider.client.CreatePod(ctx, &request)
	if err != nil {
		return err
	}
	provider.pods[key] = pod
	return nil
}

// UpdatePod takes a Kubernetes Pod and updates it within the provider.
func (*GrpcProvider )UpdatePod(ctx context.Context, pod *corev1.Pod) error {
	logger := log.GetLogger(ctx)
	logger.Error("UpdatePod")
	return nil
}

// DeletePod takes a Kubernetes Pod and deletes it from the provider. Once a pod is deleted, the provider is
// expected to call the NotifyPods callback with a terminal pod status where all the containers are in a terminal
// state, as well as the pod. DeletePod may be called multiple times for the same pod.
func (provider *GrpcProvider) DeletePod(ctx context.Context, pod *corev1.Pod) error {
	logger := log.GetLogger(ctx)
	logger.Info("DeletePod")

	pod_json, err := json.Marshal(pod)
	if err != nil {
		logger.Errorf("Couldn't marshal pod %v. %s", pod, err)
		return err
	}

	request := DeletePodRequest{
		CoreV1PodJson: string(pod_json),
	}
	_, err = provider.client.DeletePod(ctx, &request)
	if err != nil {
		logger.Errorf("Request failed %v. %s", request, err)
		return err
	}
	return nil
}

// GetPod retrieves a pod by name from the provider (can be cached).
// The Pod returned is expected to be immutable, and may be accessed
// concurrently outside of the calling goroutine. Therefore it is recommended
// to return a version after DeepCopy.
func (*GrpcProvider) GetPod(ctx context.Context, namespace, name string) (*corev1.Pod, error) {
	// This is only called for 2 reasons.
	// 1. To check whether to call Update or Create a pod.
	// 2. Ensure the pod exists before deleting it.
	logger := log.GetLogger(ctx)
	logger.Fatal("GetPod should never be called")
	return nil, nil
}

// GetPodStatus retrieves the status of a pod by name from the provider.
// The PodStatus returned is expected to be immutable, and may be accessed
// concurrently outside of the calling goroutine. Therefore it is recommended
// to return a version after DeepCopy.
func (*GrpcProvider) GetPodStatus(ctx context.Context, namespace, name string) (*corev1.PodStatus, error) {
	logger := log.GetLogger(ctx)
	logger.Error("GetPodStatus")
	return nil, nil
}

// GetPods retrieves a list of all pods running on the provider (can be cached).
// The Pods returned are expected to be immutable, and may be accessed
// concurrently outside of the calling goroutine. Therefore it is recommended
// to return a version after DeepCopy.
func (*GrpcProvider) GetPods(ctx context.Context) ([]*corev1.Pod, error) {
	logger := log.GetLogger(ctx)

	logger.Error("GetPods")
	return nil, nil
}


// buildKey is a helper for building the "key" for the providers pod store.
func buildKey(pod *corev1.Pod) (string, error) {
	if pod.ObjectMeta.Namespace == "" {
		return "", fmt.Errorf("pod namespace not found")
	}

	if pod.ObjectMeta.Name == "" {
		return "", fmt.Errorf("pod name not found")
	}

	return buildKeyFromNames(pod.ObjectMeta.Namespace, pod.ObjectMeta.Name)
}

func buildKeyFromNames(namespace string, name string) (string, error) {
	return fmt.Sprintf("%s-%s", namespace, name), nil
}
