# NOTE: This Dockerfile is designed to be built via `make skaffold/build`.
FROM golang:bookworm

SHELL ["/bin/bash", "-c"]


ENV APISERVER_CERT_LOCATION /vkubelet-mock-0-crt.pem
ENV APISERVER_KEY_LOCATION /vkubelet-mock-0-key.pem
ENV KUBELET_PORT 10250

RUN echo $' \n\
deb https://ftp.debian.org/debian/ bookworm contrib main non-free non-free-firmware \n\
deb https://ftp.debian.org/debian/ bookworm-updates contrib main non-free non-free-firmware  \n\
deb https://ftp.debian.org/debian/ bookworm-proposed-updates contrib main non-free non-free-firmware  \n\
deb https://ftp.debian.org/debian/ bookworm-backports contrib main non-free non-free-firmware  \n\
deb https://security.debian.org/debian-security/ bookworm-security contrib main non-free non-free-firmware \n\
' > /etc/apt/sources.list

RUN apt-get update && apt-get install -y unzip protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

ENV PATH="${PATH}:/go/bin"


RUN find /go/bin

RUN mkdir /src
ADD ./ /src/virtual-kubelet
WORKDIR /src/virtual-kubelet

# Clean away mach-o artifacts.
RUN make clean build

RUN cp bin/virtual-kubelet /virtual-kubelet

COPY hack/skaffold/virtual-kubelet/vkubelet_grpc_cfg.yaml /vkubelet_grpc_cfg.yaml
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