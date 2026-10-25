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

## chip

```bash
# kiểm tra chip server
uname -m
```
- linux/amd64 (EC2): Linux + chip kiến trúc x86_64
    - Mac đời cũ chip intel
- linux/arm64: Linux + chip kiến trúc ARM 64-bit
    - Mac đời mới: m1, m2, m3, m4, m5
- windows/amd64: Windows + chip kiến trúc x86_64

Lệnh build image
```bash
# linux/amd64
# linux/arm64
docker build --platform=linux/amd64
docker build --platform=linux/x86_64
```

Dockerfile
```dockerfile
FROM --plaform=linux/amd64 node:18-alpine
```

docker-compose.yml
```yml
services:
    js_nestjs:
        platform: linux/amd64
        build:
        context: contextPath
        dockerfile: Dockerfile
```

## nginx

https://nginx.org/

```bash
sudo apt update && sudo apt install nginx
nginx -v
```

- nginx.conf: cấu hình chung
- /conf.d: các file bên trong là cấu hình bổ sung cho cấu hình chung
- sites-enabled: chứa các file config của web site đang chạy
- sites-available: chưa các file config lưu trữ, muốn chạy cái nào thì link qua bên sites-enabled

- proxy_params: những cái config sài chung cho các config website khác
    - HTTP: forward header cho BE

- fastcgi_params, scgi_params, uwsgi_params: dành cho các giao thức khác
- snippets: cũng là sài chung giống `params`, ý nghĩa để sắp xếp gọn gàng
- modules-available, modules-enabled: quản lý import thư viên bên ngoài, cài thêm
- koi-utf, win-utf, koi-win: chuyển đổi mã hoá ký tự HTTP reponse gửi về cho client

### LINK qua sites-enabled
```bash
# bắt buộc sài đường dẫn tuyệt đối
sudo ln -s /etc/nginx/sites-available/my_web.conf /etc/nginx/sites-enabled
# ln: link
# -s: symbolink (tạo liên kết mềm)

sudo nginx -t
sudo systemctl reload nginx
```
