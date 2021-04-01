podTemplate(
cloud: 'kubernetes',
containers: [
    containerTemplate(name: 'docker', image: 'docker:19.03.1-dind', ttyEnabled: true, command: 'cat'),
    containerTemplate(name: 'argocd', image: 'argoproj/argocd', ttyEnabled: true, command: 'cat'),
    ],
volumes: [
    hostPathVolume(
        mountPath: '/var/run/docker.sock', 
        hostPath: '/var/run/docker.sock'
    )
]){    
    node(POD_LABEL){
        stage('build'){
            container('docker'){
                checkout scm
                sh 'docker build --build-arg BUILD_VERSION=${BRANCH_NAME} . -t app-svc:latest'
            }
        }
    }
}
