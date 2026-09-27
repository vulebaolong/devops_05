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
                sh 'docker build -t js_nextjs:latest ./buoi_3_4_5/js_nextjs'
            }
        }

        stage('Deploy Container') {
            steps {
                sh '''
                    docker rm -f js_nextjs-container || true
                    docker run -d --name js_nextjs-container -p 3002:3000 js_nextjs:latest
                '''
            }
        }
    }
}