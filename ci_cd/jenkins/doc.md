## Luanch instantce
- m7i-flex.large
    - 0.1232$ / 1 tiếng => 3 tiếng = 11.000 VNĐ
- ổ cứng 20 GB

## Cài jenkins
https://www.jenkins.io/doc/book/installing/linux/#debianubuntu

- cài đúng thứ tự, java trước, jenkin sau

### cài java
- https://www.jenkins.io/doc/book/installing/linux/#debian-java
```bash
sudo apt update
sudo apt install fontconfig openjdk-21-jre
java -version
```
### cài jenkins
- https://www.jenkins.io/doc/book/installing/linux/#debian-stable
```bash
sudo wget -O /etc/apt/keyrings/jenkins-keyring.asc \
  https://pkg.jenkins.io/debian-stable/jenkins.io-2026.key
echo "deb [signed-by=/etc/apt/keyrings/jenkins-keyring.asc]" \
  https://pkg.jenkins.io/debian-stable binary/ | sudo tee \
  /etc/apt/sources.list.d/jenkins.list > /dev/null
sudo apt update
sudo apt install jenkins

# kiểm tra
jenkins --version
```

## Thiết lập riêng cho jenkins

- cài docker
- thêm jenkins vào nhóm docker để sử dụng lệnh không cần phải sudo nữa
    ```bash
    sudo usermod -aG docker jenkins
    ```
- stop jenkins
    ```bash
    sudo systemctl stop jenkins
    sudo systemctl status jenkins
    ```

- đổi editor mặc định của jenkins
    ```bash
    sudo update-alternatives --config editor
    # chọn số 3
    ```

- Setup cấu hình
    ```bash
    sudo systemctl edit jenkins
    ```
```conf
[Service]
Environment="JAVA_OPTS=-Djava.awt.headless=true -Djava.net.preferIPv4Stack=true -Djava.io.tmpdir=/var/cache/jenkins/tmp/"
Environment="JENKINS_OPTS=--pluginroot=/var/cache/jenkins/plugins"
```

- Tạo folder đã setup, và set user
```bash
sudo mkdir /var/cache/jenkins/tmp
sudo chown -R jenkins:jenkins /var/cache/jenkins/tmp

sudo mkdir /var/cache/jenkins/plugins
sudo chown -R jenkins:jenkins /var/cache/jenkins/plugins

# chown: change own là thay đổi chủ nhân
# -R: đệ quy, quét toàn bộ file, thư mục con bên trong

# kiểm tra
sudo su
cd /var/cache/jenkins
ls -la

# drwxr-xr-x  2 jenkins jenkins 4096 Nov  7 07:52 plugins
# drwxr-xr-x  2 jenkins jenkins 4096 Nov  7 07:52 tmp
```
- start lại jenkins
    ```bash
    sudo systemctl start jenkins
    sudo systemctl status jenkins
    ```

## cài nginx
```bash
sudo apt update && sudo apt install nginx
```

## script jenkins

```groovy
pipeline {
    agent any

    stages {
        stage('Hello') {
            steps {
                echo 'Hello World'
            }
        }
        
        stage('Checkout Code') {
            steps {
                git branch: 'main', credentialsId: 'github-login', url: 'https://github.com/vulebaolong/go_gin_be_jenkins'
                sh 'ls -la'
            }
        }
        
        stage('Build Docker Image') {
            steps {
                sh 'docker build -t go_gin_be_img .'
            }
        }
        
        stage('Deploy Container Move Image') {
            steps {
                withCredentials([
                    sshUserPrivateKey(credentialsId:'ssh-private', keyFileVariable:'SSH_KEY',usernameVariable:'SSH_USER'),
                    string(credentialsId:'ip-server-run', variable:'SSH_IP')
                ]) {
                    sh '''
                        docker save go_gin_be_img | gzip | ssh -o StrictHostKeyChecking=no -i ${SSH_KEY} ${SSH_USER}@${SSH_IP} "gunzip | docker load"
                        
                        ssh -o StrictHostKeyChecking=no -i ${SSH_KEY} ${SSH_USER}@${SSH_IP} "
                            cd /home/ubuntu/devops_05
    
                            docker compose up -d --force-recreate go_gin_be
                        "
                    '''
                }
            }
        }
    }
}

```

## Phân quyền tạo tài khoản

- 1. cài plungin: Role-based Authorization Strategy
- 2. Setting -> Security -> Authorization
    - chuyển Anyone can do anything
    - thành: Role-based Strategy (plugin mới cài)
- 3. tạo tài khoản cho member
- 4. tạo quyền và gắn quyền cho tài khoản

- global
- item
- agent - nhiều máy xử lý build