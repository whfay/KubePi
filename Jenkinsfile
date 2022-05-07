
pipeline {
    agent {
        node {
            label '165'
        }
    }
    environment { 
        IMAGE_NAME = 'metersphere'
        IMAGE_PREFIX = 'registry.kubeoperator.io:8083/kubeoperator'
    }
    stages {
        stage('拉取 git 仓库代码') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: '$tag']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/whfay/KubePi.git']]])
            }
        }
        stage('Docker build') {
            steps {
                echo 'build success'
            }
        }
        stage('Docker push') {
            steps {
                echo 'push success'
            }
        }
    }
}
