pipeline {
    agent {
        node {
            label '165'
        }
    }
    environment {
        IMAGE_PREFIX = 'registry.kubeoperator.io:8083'
        IMAGE_REPO = 'kubeoperator'
        IMAGE_NAME = 'kubepi-server'
        REGISTRY_USER = 'admin'
        REGISTRY_PASSWORD = 'admin123'
    }
    stages {
        stage('拉取 git 仓库代码') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: '$branch']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/whfay/KubePi.git']]])
            }
        }
        stage('构建 docker 镜像') {
            steps {
                sh '''make build_docker
                docker tag kubeoperator/kubepi-server:master $IMAGE_PREFIX/$IMAGE_REPO/$IMAGE_NAME:$tag'''
            }
        }
        stage('推送 docker 镜像到镜像仓库') {
            steps {
                sh '''docker login $IMAGE_PREFIX -u $REGISTRY_USER -p $REGISTRY_PASSWORD
                docker push $IMAGE_PREFIX/$IMAGE_REPO/$IMAGE_NAME:$tag'''
            }
        }
        stage('部署服务') {
            steps {
                sshPublisher(publishers: [sshPublisherDesc(configName: '181', transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: "deploy_kubepi.sh $IMAGE_PREFIX $IMAGE_REPO $IMAGE_NAME $tag $host_port $container_port", execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: '', remoteDirectorySDF: false, removePrefix: '', sourceFiles: '')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: false)])
            }
        }
    }
}
