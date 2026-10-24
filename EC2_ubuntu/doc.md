```bash
# Cặp RSA
ssh-keygen -t rsa -b 4096 -C "devops_demo"

# Cặp ed25519 (ngắn hơn, hiện đại hơn, bảo mật hơn)
ssh-keygen -t ed25519 -C "devops_demo"


# Enter file in which to save the key (/Users/vulebaolong/.ssh/id_ed25519):
# Dán đường dẫn folder muốn lưu key, và tên file

# Enter passphrase (empty for no passphrase)
# Lớp bảo vệ thứ 2, Enter bỏ qua

# nạp user ubuntu vào group docker
# để sử dụng lệnh docker không cần thêm sudo
sudo usermod -aG docker ubuntu
sudo adduser ubuntu docker
```