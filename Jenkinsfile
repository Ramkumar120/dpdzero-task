pipeline {
    agent any

    stages {
        stage('clean work space') {
            steps {
                cleanWs()
            }
        }
        
        stage('Git Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/Ramkumar120/dpdzero-task.git'
            }
        }
        
        stage('Build Docker Image') {
            steps {
                sh '''
                docker compose down
                docker compose build
                '''
            }
        }
        
        stage('Run Containers') {
            steps {
                sh 'docker compose up -d'
            }
        }
    }
}
