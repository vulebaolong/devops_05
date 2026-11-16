

- tạo sg chung tên `sg_k8s`
    - tạo trước 1 sg rỗng (không thiết lập gì hết chỉ đặt tên và description)
    - vào lại sg mới tạo đó
    - inbound thêm 1 cái rule
        - Type: All traffic
        - Source: chọn sg mới tạo
    - gắn `sg_k8s` vào tất cả máy
    - mỗi máy sẽ có sg riêng của máy đó và `sg_k8s`

- mua máy control plane
    - lưu ý mọi máy đều phải cùng `Availability Zone`
    - gắn security group: `sg_k8s`

- mang 2 file `common.sh`, `master.sh` lên Controll plane

- EC2 control plane: chạy 2 file
    - chạy trước `common.sh`
    - chạy sau `master.sh`
- EC2 node (cung cấp thêm tài nguyên): 
    - chỉ chạy `common.sh`

```bash
vim common.sh
vim master.sh

sudo bash common.sh
sudo bash master.sh
# sẽ bị hỏi "Enter Control Plane Private IP:"
# điền địa chỉ IP PRIVATE
```

# lệnh k8s
```bash
kubectl get node
# NAME              STATUS   ROLES           AGE   VERSION
# ip-172-31-38-66   Ready    control-plane   14m   v1.36.6

kubectl get pod
# chỉ lấy ra những pod mình tự start

kubectl get pod -A
# lấy ra tất cả pod kể cả pod hệ thống của k8s
```

## sau khi đã chuẩn bị xong 2 máy control plane và node
- đảm bảo 2 máy control plane và node đều gắn sg_k8s
- gắn máy node -> control plane

- control plane
```bash
kubeadm token create --print-join-command
# kubeadm join 172.31.38.66:6443 --token 98ba2t.926fchl5mnxvfmvc --discovery-token-ca-cert-hash sha256:82b3a905f8732e7a23e0bda545798267711a05df8ab94b80afa813921aa30488 
```
- node:
    - mang lệnh được tạo ra ở control plane qua máy node chạy
    - thêm sudo ở đầu

- bên máy control plane đổi label cho đẹp
```bash
kubectl get node
# NAME               STATUS   ROLES           AGE   VERSION
# ip-172-31-33-217   Ready    <none>          57s   v1.36.6
# ip-172-31-38-66    Ready    control-plane   31m   v1.36.6

kubectl label node [NAME] node-role.kubernetes.io/worker=worker
kubectl label node ip-172-31-33-217 node-role.kubernetes.io/worker=worker
```

## Cài Metric Server
- để k8s đọc được thông tin máy để scale
- mang file `metrics-server.yaml` lên control plane
```bash
kubectl apply -f metrics-server.yaml

# check
kubectl get pod -A
```
- sau khi cài metric thành công sẽ chạy được 2 cái lệnh
```bash
kubectl top node
kubectl top pod
```

- mang env fe và be lên server
- nạp env vào k8s
```bash
kubectl create secret generic nodepad-be-env \
    --from-env-file=.env \
    --dry-run=client \
    -o yaml | kubectl apply -f -

# kiểm tra
kubectl get secret
kubectl describe secret
```

## run
- đưa file `nodepad-be.yaml` lên server
```bash
kubectl apply -f nodepad-be.yaml

kubectl get pod
kubectl logs 

# show chi tiết pod đang chạy ở máy nào
kubectl get pod -o wide

```

### tạo server db
- gắn sg_k8s
- cài docker
```bash
docker run --name database -e POSTGRES_PASSWORD=12345 -d -p 5432:5432 postgres

docker exec -it database bash
psql -U postgres
CREATE DATABASE nodepad;
\c nodepad

# lấy file migration trong source BE trong
# k8s/nodepad/nodepad-be/ent/migrate/migrations/20260728083226_init.sql

\dt
```

### lệnh check xem giữa các EC2 nói chuyện được với nhau không
```bash
nc -zv [IP_PRIVATE] port
nc -zv 172.31.33.62 5432
```

chạy lại pod
```bash
kubectl rollout restart deployment/nodepad-be
```

### run source fe
- mang file `nodepad-fe.yaml` lên server
```bash
kubectl apply -f nodepad-fe.yaml
```

## GẮN DOMAIN cài nginx ingress controller
- mang 2 file ssl lên server
```bash
kubectl create secret tls vulebaolong.com \
    --namespace default \
    --cert=vulebaolong.com.cer \
    --key=vulebaolong.com.key \
    --dry-run=client \
    -o yaml | kubectl apply -f -

# kiểm tra
kubectl get secret
kubectl describe secret
```

### cài helm
> Công cụ cài đặt và quản lý ứng dụng trên k8s

https://helm.sh/docs/intro/install/#from-script
```bash
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-4 | bash

helm version
```

- để ép nginx chỉ chạy trên control plane
- gắn label=ingres-node=true cho control plane
```bash
kubectl label node [TEN_NODE_CONTROL_PLANE] ingress-node=true --overwrite
kubectl label node ip-172-31-37-81 ingress-node=true --overwrite

# lấy tên TEN_NODE_CONTROL_PLANE
kubectl get node
kubectl get node --show-labels
```

https://docs.nginx.com/nginx-ingress-controller/install/helm/open-source/#oci-registry
```bash
helm install nginx-ingress \
    oci://ghcr.io/nginx/charts/nginx-ingress \
    --version 2.8.1 \
    --namespace nginx-ingress \
    --create-namespace \
    --values /home/ubuntu/nginx-ingress.yaml

# check
kubectl get pods -A -o wide
```

## cài rancher
- tối thiếu: `m7i-flex.large`
    - 8GB Ram
    - 40GB ổ cứng

https://ranchermanager.docs.rancher.com/v2.15/getting-started/quick-start-guides/deploy-rancher-manager/helm-cli#_install_k3s_product_name_on_linux
```bash
curl -sfL https://get.k3s.io | INSTALL_K3S_CHANNEL=v1.36 sh -s - server --cluster-init

# kiểm tra
sudo k3s kubectl get node
```

- cấu hình
```bash
mkdir -p ~/.kube
sudo cp /etc/rancher/k3s/k3s.yaml ~/.kube/config
sudo chown ubuntu:ubuntu ~/.kube/config
chmod 600 ~/.kube/config
export KUBECONFIG=$HOME/.kube/config
```

Lưu cấu hình vĩnh viễn
Để lần sau SSH vào EC2 không cần export lại:
```bash
echo 'export KUBECONFIG=$HOME/.kube/config' >> ~/.bashrc
source ~/.bashrc

# kiểm tra
kubectl get nodes
```

https://ranchermanager.docs.rancher.com/v2.15/getting-started/quick-start-guides/deploy-rancher-manager/helm-cli#_install_rancher_with_helm
```bash
helm repo add rancher-latest https://releases.rancher.com/server-charts/latest

kubectl create namespace cattle-system

kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.21.3/cert-manager.crds.yaml

helm repo add jetstack https://charts.jetstack.io

helm repo update

helm install cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace
```

```bash
helm install rancher rancher-latest/rancher \
  --namespace cattle-system \
  --set hostname=rancher.vulebaolong.com \
  --set replicas=1 \
  --set bootstrapPassword=matkhau
```

- chạy và đợi cho đến khi rancher khởi động thành công thì mới cài tiếp
```bash
kubectl get pod -n cattle-system -w
```

## map domain và gắn ssl

- Hiện Rancher đang dùng certificate do cert-manager quản lý, nên trước tiên chuyển Helm sang chế độ secret:
```bash
helm upgrade rancher rancher-latest/rancher \
  --namespace cattle-system \
  --reuse-values \
  --set ingress.tls.source=secret 

# kiểm tra
helm get values rancher -n cattle-system
# ingress:
#   tls:
#     source: secret
```

- Vì Rancher đã chuyển sang ingress.tls.source=secret, xóa cái cũ
```bash
kubectl delete certificate tls-rancher-ingress -n cattle-system

# kiểm tra
kubectl get certificate -n cattle-system
# No resources found in cattle-system namespace.
```

- mang 2 file ssl lên server
- Tạo Kubernetes TLS Secret
- thư mục ~/ssl, chạy:
```bash
kubectl create secret tls tls-rancher-ingress \
  --namespace cattle-system \
  --cert=./vulebaolong.com.cer \
  --key=./vulebaolong.com.key \
  --dry-run=client \
  -o yaml | kubectl apply -f -

```


## kết nối rancher với cluster
- setting trước:
    - Rancher UI
    - Global Setting
    - tìm agent-tls-mode: `Strict` chuyển thành `System Store`

- import
    - Rancher UI
    - Cluster Management
    - Import Exitsting
    - Generic
    - Cluster Name: `nodepad`
    - Create
    - copy lệnh cmd đầu tiên, mang qua cluster (control plane) chạy



## test
### tăng nhẹ
```bash
# Rancher: Service Discovery -> Services -> Cluster IP
# để lấy clusterIP
kubectl get svc nodepad-be-service

while true; do
  curl -s -o /dev/null --max-time 5 \
    -H 'Content-Type: application/json' \
    -d '{"email":"admin@gmail.com","password":"12345"}' \
    http://<CLUSTER-IP>:8888/api/auth/login
done
```

### tăng mạnh
```bash
# Rancher: Service Discovery -> Services -> Cluster IP
kubectl get svc nodepad-be-service

#  -P 100 là 100 request chạy đồng thời; có thể tăng lên 200. Nhấn Ctrl+C để dừng.
seq 1000000 | xargs -P 100 -I '{}' curl -s -o /dev/null \
  --max-time 5 \
  -H 'Content-Type: application/json' \
  -d '{"email":"admin@gmail.com","password":"12345"}' \
  http://<CLUSTER-IP>:8888/api/auth/login
```
