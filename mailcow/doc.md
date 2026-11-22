## luanch server
- 8GB: m7i-flex.large
- 50GB

## cài đặt mailcow

- https://docs.mailcow.email/getstarted/prerequisite-system/

- Kiểm tra
```bash
ss -tlpn | grep -E -w '25|80|110|143|443|465|587|993|995|4190'
```

- vào sercurity group mở port
- https://docs.mailcow.email/getstarted/prerequisite-system/#incoming-ports

### Cài docker
### Quick Installation

```bash
sudo apt update
sudo apt install -y git openssl curl gawk coreutils grep jq
```
### Installing mailcow
https://docs.mailcow.email/getstarted/install/#installing-mailcow

```bash
# chạy trước
sudo -i

# chạy sau
umask 0022
cd /opt
git clone https://github.com/mailcow/mailcow-dockerized
cd mailcow-dockerized

sudo bash generate_config.sh

# Mail server hostname (FQDN) - this is not your mail domain, but your mail servers hostname: mail.vulebaolong.com
# Timezone [Etc/UTC]: Asia/Ho_Chi_Minh
# Choose the Branch with it's number [1/2/3]: 1
```

## Starting mailcow
### tải image về
```bash
docker compose pull
```

- trước khi run bằng `docker compose up -d`, map DNS bên cloudflare
- để khi start, thì service của mailcow sẽ xác thực chứng chỉ luôn
- nếu start trước, map DNS, thì cần restart lại service docker compose `acme-mailcow`

### map domain DNS
- vào cloudflare
    - mail.vulebaolong.com
        - Type: A
        - Name: mail
        - IPv4 Address: IP_PUBLIC
        - Proxy status: `OFF` -> DNS ONLY
        - TTL: Auto
    - @vulebaolong.com
        - Type: MX
        - Name: @
        - Mail Server: mail.vulebaolong.com
        - Priority: 10
        - TTL: Auto

### run
```bash
docker compose up -d
```

## tao domain
- Email -> Configuation -> Domain (tab) -> + Add domain
- Tên miền: vulebaolong.com
- Click: Thêm tên miền và khởi động lại SOGo
- Click: DNS để kiểm tra xem đã xử lý
    - mail.vulebaolong.com - A
    - vulebaolong.com - MX

## tạo mail box
- Email -> Configuation -> Mailboxes (tab) -> + Add mailbox

## mặc định AWS chặn cổng 25 
- khi gửi email ra ngoài, để tránh spam
- nếu muốn mở thì xin phép AWS, đợi duyệt
- Brevo: xử lý bị aws chặn

## Brevo
- xác thực sđt

### Authorize IP address
- Setting -> Security -> Authorized IPs -> Click Authorize IP Addresses
- điền địa chỉ IP PUBLIC

### SMTP keys
- Setting -> SMTP & API -> Click Generate SMTP key

### Xác thực domain
- Setting -> Senders, Domains & Dedicated IPs -> Domains -> Click Add a Domain
- step 1:
    - Domain name: vulebaolong.com
- step 2:
    - Branded subdomain: brevo.vulebaolong.com
- step 3:
    - Automatic
    - Đăng nhập Cloudflare
    - Authorize

##  chuyển gửi email tới 25 qua bên brevo
- thay vì gửi chính thống tới server gmail 25
- thì bây giờ gửi qua bên brevo    

### tạo định tuyến



