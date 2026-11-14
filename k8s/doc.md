

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
kubectl create secret generic nodepad-fe-env \
    --from-env-file=.env \
    --dry-run=client \
    -o yaml | kubectl apply -f -

# kiểm tra
kubectl get secret
```

## run
- đưa file `nodepad-be.yaml` lên server
```bash
kubectl apply -f nodepad-be.yaml

kubectl get pod
kubectl logs 
```