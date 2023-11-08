package grpc

import (
	"context"

	corev1 "k8s.io/api/core/v1"
)


// CreatePod takes a Kubernetes Pod and deploys it within the provider.
func (*GrpcProvider) CreatePod(ctx context.Context, pod *corev1.Pod) error {
	return nil
}

// UpdatePod takes a Kubernetes Pod and updates it within the provider.
func (*GrpcProvider )UpdatePod(ctx context.Context, pod *corev1.Pod) error {
	return nil
}

// DeletePod takes a Kubernetes Pod and deletes it from the provider. Once a pod is deleted, the provider is
// expected to call the NotifyPods callback with a terminal pod status where all the containers are in a terminal
// state, as well as the pod. DeletePod may be called multiple times for the same pod.
func (*GrpcProvider) DeletePod(ctx context.Context, pod *corev1.Pod) error {
	return nil
}

// GetPod retrieves a pod by name from the provider (can be cached).
// The Pod returned is expected to be immutable, and may be accessed
// concurrently outside of the calling goroutine. Therefore it is recommended
// to return a version after DeepCopy.
func (*GrpcProvider) GetPod(ctx context.Context, namespace, name string) (*corev1.Pod, error) {
	return nil, nil
}

// GetPodStatus retrieves the status of a pod by name from the provider.
// The PodStatus returned is expected to be immutable, and may be accessed
// concurrently outside of the calling goroutine. Therefore it is recommended
// to return a version after DeepCopy.
func (*GrpcProvider) GetPodStatus(ctx context.Context, namespace, name string) (*corev1.PodStatus, error) {
	return nil, nil
}

// GetPods retrieves a list of all pods running on the provider (can be cached).
// The Pods returned are expected to be immutable, and may be accessed
// concurrently outside of the calling goroutine. Therefore it is recommended
// to return a version after DeepCopy.
func (*GrpcProvider) GetPods(context.Context) ([]*corev1.Pod, error) {
	return nil, nil
}
