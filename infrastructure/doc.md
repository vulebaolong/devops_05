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
## Ansible
```bash
docker run --name ubuntu-ansible -d -it -v ./ansible:/root/ansible ubuntu

docker exec -it ubuntu-ansible bash
```

### install
- truy cập: https://docs.ansible.com/
- header: resource
    - Ansible package documentation
- menu bên trái:
    - Installation, Upgrade & Configuration
        - Installation Guide
        - Installing Ansible on specific operating systems
        - Installing Ansible on Ubuntu: https://docs.ansible.com/projects/ansible/latest/installation_guide/installation_distros.html#installing-ansible-on-ubuntu

```bash
apt update
apt install software-properties-common
add-apt-repository --yes --update ppa:ansible/ansible
apt install ansible

# Geographic area: 5. Asia
# Time zone: 29. Ho_Chi_Minh
```

### setting

- Ansible getting started
    - Getting started with Ansible
    - Building an inventory: https://docs.ansible.com/projects/ansible/latest/getting_started/get_started_inventory.html#building-an-inventory

- tạo file inventory
```bash
ansible-inventory -i inventory.ini --list

ansible myhosts -m ping -i inventory.ini
# [ERROR]: Task failed: Failed to connect to the host via ssh: Host key verification failed.

# tạo file ansible.cfg

# chạy lại
ansible myhosts -m ping -i inventory.ini
# [ERROR]: Task failed: Failed to connect to the host via ssh: Warning: Permanently added '13.212.6.43' (ED25519) to the list of known hosts.

# file inventory.ini
# [myhosts:vars]
# ansible_user=ubuntu
# ansible_ssh_private_key_file=./devops.pem

# chạy lại
ansible myhosts -m ping -i inventory.ini
# [WARNING]: Host '18.143.179.34' is using the discovered Python interpreter at '/usr/bin/python3.12', but future installation of another Python interpreter could cause a different interpreter to be discovered. See https://docs.ansible.com/ansible-core/2.21/reference_appendices/interpreter_discovery.html for more information.

# file inventory.ini
# ansible_python_interpreter=/usr/bin/python3

# 13.212.6.43 | SUCCESS => {
#     "changed": false,
#     "ping": "pong"
# }
# 18.143.179.34 | SUCCESS => {
#     "changed": false,
#     "ping": "pong"
# }
```

### playbook
https://docs.ansible.com/projects/ansible/latest/getting_started/get_started_playbook.html#creating-a-playbook

- run playbook
```bash
ansible-playbook -i inventory.ini nginx.yaml
```