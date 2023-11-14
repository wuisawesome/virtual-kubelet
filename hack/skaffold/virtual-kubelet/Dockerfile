# NOTE: This Dockerfile is designed to be built via `make skaffold/build`.


FROM golang:bookworm

ENV APISERVER_CERT_LOCATION /vkubelet-mock-0-crt.pem
ENV APISERVER_KEY_LOCATION /vkubelet-mock-0-key.pem
ENV KUBELET_PORT 10250



RUN mkdir /src
ADD ./ /src/virtual-kubelet
WORKDIR /src/virtual-kubelet

# Clean away mach-o artifacts.
RUN make clean build

RUN cp bin/virtual-kubelet /virtual-kubelet

COPY hack/skaffold/virtual-kubelet/vkubelet_grpc_cfg.yaml /vkubelet_grpc_cfg.yml
# Copy the certificate for the HTTPS server.
COPY hack/skaffold/virtual-kubelet/vkubelet-mock-0-crt.pem /vkubelet-mock-0-crt.pem
# Copy the private key for the HTTPS server.
COPY hack/skaffold/virtual-kubelet/vkubelet-mock-0-key.pem /vkubelet-mock-0-key.pem



CMD ["/virtual-kubelet"]










































# ADD hack/skaffold/virtual-kubelet/ /

# # Copy the configuration file for the grpc provider.
# COPY /users/Alex/repos/virtualkubelet/virtual-kubelet/hack/skaffold/virtual-kubelet/vkubelet_grpc_cfg.yaml /vkubelet_grpc_cfg.yml
# # Copy the certificate for the HTTPS server.
# COPY /users/Alex/repos/virtualkubelet/virtual-kubelet/hack/skaffold/virtual-kubelet/vkubelet-mock-0-crt.pem /vkubelet-mock-0-crt.pem
# # Copy the private key for the HTTPS server.
# COPY /users/Alex/repos/virtualkubelet/virtual-kubelet/hack/skaffold/virtual-kubelet/vkubelet-mock-0-key.pem /vkubelet-mock-0-key.pem