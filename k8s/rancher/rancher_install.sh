#!/bin/bash

set -e
set -o pipefail

# https://ranchermanager.docs.rancher.com/getting-started/quick-start-guides/deploy-rancher-manager/helm-cli#_install_k3s_product_name_on_linux
curl -sfL https://get.k3s.io | INSTALL_K3S_CHANNEL=v1.36 sh -s - server --cluster-init

# https://ranchermanager.docs.rancher.com/getting-started/quick-start-guides/deploy-rancher-manager/helm-cli#_save_the_kubeconfig_to_your_workstation
mkdir -p ~/.kube
sudo cp /etc/rancher/k3s/k3s.yaml ~/.kube/config
sudo chown ubuntu:ubuntu ~/.kube/config
chmod 600 ~/.kube/config
export KUBECONFIG=$HOME/.kube/config

# Lưu cấu hình vĩnh viễn, để lần sau SSH vào EC2 không cần export lại:
echo 'export KUBECONFIG=$HOME/.kube/config' >>~/.bashrc
source ~/.bashrc

# cài helm
curl -fsSL https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

# Install Rancher with Helm
# https://ranchermanager.docs.rancher.com/getting-started/quick-start-guides/deploy-rancher-manager/helm-cli#_install_rancher_with_helm

helm repo add rancher-latest https://releases.rancher.com/server-charts/latest

kubectl create namespace cattle-system

kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.21.3/cert-manager.crds.yaml

helm repo add jetstack https://charts.jetstack.io

helm repo update

helm install cert-manager jetstack/cert-manager \
    --namespace cert-manager \
    --create-namespace \
    --wait \
    --timeout 10m

# Cài Rancher
helm install rancher rancher-latest/rancher \
    --namespace cattle-system \
    --set hostname=rancher.vulebaolong.com \
    --set replicas=1 \
    --set bootstrapPassword=matkhau

# Chờ Deployment Rancher sẵn sàng; lỗi hoặc quá 10 phút thì set -e dừng script.
kubectl rollout status deployment/rancher -n cattle-system --timeout=10m

# Hiện Rancher đang dùng certificate do cert-manager quản lý, nên trước tiên chuyển Helm sang chế độ secret:
helm upgrade rancher rancher-latest/rancher \
  --namespace cattle-system \
  --reuse-values \
  --set ingress.tls.source=secret 

# kiểm tra
helm get values rancher -n cattle-system

# Vì Rancher đã chuyển sang ingress.tls.source=secret, xóa cái cũ
kubectl delete certificate tls-rancher-ingress -n cattle-system

# kiểm tra
kubectl get certificate -n cattle-system


# Tạo thư mục nếu chưa có.
ssl_dir="$HOME/ssl"
mkdir -p "$ssl_dir"

echo "Dán nội dung CERT, nhấn Enter rồi Ctrl+D để kết thúc:"
tee "$ssl_dir/tls.crt" > /dev/null

echo "Dán nội dung KEY, nhấn Enter rồi Ctrl+D để kết thúc:"
tee "$ssl_dir/tls.key" > /dev/null

kubectl create secret tls tls-rancher-ingress \
  --namespace cattle-system \
  --cert="$ssl_dir/tls.crt" \
  --key="$ssl_dir/tls.key" \
  --dry-run=client \
  -o yaml | kubectl apply -f -

# Kiểm tra
kubectl get ingress rancher -n cattle-system -o yaml
