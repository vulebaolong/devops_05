#!/usr/bin/env bash
set -euo pipefail
# Để script dừng ngay khi một lệnh lỗi.


# https://github.com/cri-o/packaging/blob/main/README.md#usage

KUBERNETES_VERSION=v1.36
CRIO_VERSION=v1.36


# Cài dependency để thêm repository
# Install the dependencies for adding repositories
# https://github.com/cri-o/packaging/blob/main/README.md#install-the-dependencies-for-adding-repositories
apt-get update
apt-get install -y software-properties-common curl

# add url và các key vào apt, để khi chạy apt sẽ gọi đúng chính chủ

# Add the Kubernetes repository: kubelet, kubeadm, kubectl
# https://github.com/cri-o/packaging/blob/main/README.md#add-the-kubernetes-repository-1
curl -fsSL https://pkgs.k8s.io/core:/stable:/$KUBERNETES_VERSION/deb/Release.key |
    gpg --batch --yes --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg

echo "deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/$KUBERNETES_VERSION/deb/ /" |
    tee /etc/apt/sources.list.d/kubernetes.list


# Add the CRI-O repository:CRI-O
# https://github.com/cri-o/packaging/blob/main/README.md#add-the-cri-o-repository-1
curl -fsSL https://download.opensuse.org/repositories/isv:/cri-o:/stable:/$CRIO_VERSION/deb/Release.key |
    gpg --batch --yes --dearmor -o /etc/apt/keyrings/cri-o-apt-keyring.gpg

echo "deb [signed-by=/etc/apt/keyrings/cri-o-apt-keyring.gpg] https://download.opensuse.org/repositories/isv:/cri-o:/stable:/$CRIO_VERSION/deb/ /" |
    tee /etc/apt/sources.list.d/cri-o.list


# Install the packages
# https://github.com/cri-o/packaging/blob/main/README.md#install-the-packages-1
# cri-o   → container runtime
# kubelet → nhận các chỉ thị từ trung tâm Kubernetes (Control Plane), sau đó ra lệnh cho cri-o thực hiện việc tạo, xóa hoặc kiểm tra sức khỏe của các container trên máy đó
# kubeadm → init hoặc join cluster
# kubectl → CLI quản lý Kubernetes
apt-get update
apt-get install -y cri-o kubelet kubeadm kubectl


# Nhằm tránh apt upgrade tự động nâng từng thành phần
# Kubernetes sai quy trình.
# https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/
apt-mark hold kubelet kubeadm kubectl

# Khởi động CRI-O
# https://github.com/cri-o/packaging/blob/main/README.md#start-cri-o
systemctl start crio.service

# README hướng dẫn chỉ start. Nhưng để CRI-O tự chạy lại khi EC2 reboot:
systemctl enable crio.service
systemctl status crio.service --no-pager

# Chuẩn bị server để tham gia Kubernetes cluster
# https://github.com/cri-o/packaging/blob/main/README.md#bootstrap-a-cluster
swapoff -a
modprobe br_netfilter
# sysctl -w net.ipv4.ip_forward=1 -> Không dùng lệnh này
# Vì chỉ được tạm thời ở 1 lần start
# Ghi vào file cấu hình để reboot EC2 vẫn giữ cấu hình.
# https://kubernetes.io/docs/setup/production-environment/container-runtimes/#prerequisite-ipv4-forwarding-optional
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.ipv4.ip_forward = 1
EOF

sudo sysctl --system


# https://kubernetes.io/docs/setup/production-environment/container-runtimes/#cgroup-drivers
# Kubelet và CRI-O phải dùng cgroup để thực thi các giới hạn:
# CPU
# RAM
# số lượng process
# I/O
#
# Vấn đề là cả hai phải dùng cùng một cgroup driver:
# kubelet → systemd
# CRI-O   → systemd
#
# CRI-O mặc định dùng systemd.
# Kubeadm từ Kubernetes 1.22 trở đi cũng mặc định
# cấu hình kubelet dùng systemd.