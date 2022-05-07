pipeline {
    agent {
        node {
            label '165'
        }
    }
    environment { 
        IMAGE_NAME = 'kubeoperator/kubepi-server'
        IMAGE_PREFIX = 'registry.kubeoperator.io:8083'
    }
    stages {
        stage('拉取 git 仓库代码') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: '$branch']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/whfay/KubePi.git']]])
            }
        }
        stage('Docker build') {
            steps {
                sh '''make build_docker
                docker tag kubeoperator/kubepi-server:master $IMAGE_PREFIX/$IMAGE_NAME:$tag'''
            }
        }
        stage('Docker push') {
            steps {
                sh '''docker login $IMAGE_PREFIX -u admin -p admin123
                docker push $IMAGE_PREFIX/$IMAGE_NAME:$tag'''
            }
        }
    }
}
