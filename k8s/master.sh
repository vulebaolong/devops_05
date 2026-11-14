#!/usr/bin/env bash
set -euo pipefail
# Để script dừng ngay khi một lệnh lỗi.


read -rp "Enter Control Plane Private IP: " CONTROL_PLANE_PRIVATE_IP # Private IP của EC2 Control Plane.
POD_CIDR="192.168.0.0/16" # Dải IP dùng cho Pod.  Giá trị này phải phù hợp với CNI sẽ cài sau kubeadm init.


# Khởi tạo Kubernetes Control Plane.
#
# --cri-socket:
# Chỉ định Kubernetes sử dụng CRI-O làm container runtime.
#
# --apiserver-advertise-address:
# Private IP mà Kubernetes API Server sử dụng.
#
# --pod-network-cidr:
# Dải IP dành cho các Pod.
kubeadm init \
    --cri-socket="unix:///var/run/crio/crio.sock" \
    --apiserver-advertise-address="${CONTROL_PLANE_PRIVATE_IP}" \
    --pod-network-cidr="${POD_CIDR}"


# Cấu hình kubeconfig cho user gọi sudo để chạy kubectl không cần sudo.
# Nếu chạy trực tiếp bằng root thì cấu hình cho root.
KUBECTL_USER="${SUDO_USER:-$(id -un)}"
KUBECTL_HOME="$(getent passwd "$KUBECTL_USER" | cut -d: -f6)"
KUBECTL_GROUP="$(id -gn "$KUBECTL_USER")"

install -d -m 700 -o "$KUBECTL_USER" -g "$KUBECTL_GROUP" "$KUBECTL_HOME/.kube"
install -m 600 -o "$KUBECTL_USER" -g "$KUBECTL_GROUP" \
    /etc/kubernetes/admin.conf "$KUBECTL_HOME/.kube/config"

# Dùng kubeconfig quản trị cho các lệnh kubectl trong script đang chạy bằng root.
export KUBECONFIG=/etc/kubernetes/admin.conf

# Cài Calico CNI v3.32.1.
curl -O https://raw.githubusercontent.com/projectcalico/calico/v3.32.1/manifests/calico.yaml

kubectl apply -f calico.yaml
