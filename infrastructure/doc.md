## Cài terraform

```bash
docker run --name ubuntu-terraform -d -it -v ./terraform:/root/terraform ubuntu

docker exec -it ubuntu-terraform bash
```

https://developer.hashicorp.com/terraform/install#linux
- Binary download
- 386: tải

- mang file teraform được giải nén ra qua bên folder hệ thống

```bash
cd /root/terraform/terraform_1.17.2_linux_386
cp terraform /usr/local/bin
terraform --version
```

## AWS CLI

https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html

- tải: truy cập đường dẫn vào chrome
    - https://awscli.amazonaws.com/v2/install.sh 

```bash
bash install.sh --system

# LỖI
# aws-cli installer: error: missing required dependencies: curl unzip. Install them via your package manager and re-run.

# cài công cụ bị thiếu để fix lỗi
apt update && apt install -y curl unzip

# chạy lại
bash install.sh --system

# LỖI
# aws-cli installer: error: gpg not found; cannot verify the installer signature. Install gnupg and re-run, re-run with --skip-signature-verification, or set AWS_CLI_INSTALL_SKIP_SIGNATURE_VERIFICATION=true, to install without verification (not recommended).

# cài công cụ bị thiếu để fix lỗi
apt update && apt install -y gnupg

# chạy lại
bash install.sh --system
```

### cấu hình aws cli
```bash
aws configure
```

- lấy AWS Access Key ID [None]:
    - vào aws console
    - click user góc phải trên
    - sercurity credentials
    - Create Access Key

- nhập thông tin:
- AWS Access Key ID [None]: AKIA......SLM4
- AWS Secret Access Key [None]: foB3feMZZq........5D/gdM6z8
- Default region name [None]: ap-southeast-1
- Default output format [None]: json

- check
```bash
aws sts get-caller-identity

# LỖI
# aws: [ERROR]: An error occurred (Pager): Unable to redirect output to pager. Received the following error when opening pager:
# [Errno 2] No such file or directory: 'less'
# Learn more about configuring the output pager by running "aws help config-vars".

# Cài công cụ bị thiếu để fix lỗi
apt update && apt install -y less

# chạy lại
aws sts get-caller-identity

```

## Tạo project terraform

- https://developer.hashicorp.com/terraform/tutorials/aws-get-started/aws-create#the-terraform-block
- tạo file terraform.tf
```bash
# di chuyển vào đúng nơi chứ file main.tf
terraform fmt
terraform init
terraform validate
terraform apply
```