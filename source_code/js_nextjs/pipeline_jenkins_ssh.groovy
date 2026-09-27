pipeline {
    agent any

    stages {
        stage('Checkout Code') {
            steps {
                git branch: 'main', 
                    credentialsId: 'github-login', 
                    url: 'https://github.com/vulebaolong/devops_document.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker build --no-cache -t js_nextjs:latest ./buoi_3_4_5/js_nextjs'
            }
        }

        stage('Deploy Container') {
            steps {
                sshagent(['ssh-private']) {
                    sh '''
                        # 1. Stream nén trực tiếp Docker Image sang Server đích và load
                        docker save js_nextjs:latest | gzip | ssh -o StrictHostKeyChecking=no ubuntu@13.214.29.76 "gunzip | docker load"

                        # 2. Re-deploy container trên server đích
                        ssh -o StrictHostKeyChecking=no ubuntu@13.214.29.76 "
                            docker rm -f js_nextjs || true &&
                            docker run -d --name js_nextjs -p 3002:3000 js_nextjs:latest &&
                            docker image prune -f
                        "
                    '''
                }
            }
        }
    }
}
