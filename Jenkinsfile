podTemplate(
cloud: 'kubernetes',
containers: [
    containerTemplate(name: 'golang', image: 'golang:1.8.0', ttyEnabled: true, command: 'cat'),
    containerTemplate(name: 'docker', image: 'docker:19.03.1-dind', ttyEnabled: true, command: 'cat')
    ],
volumes: [
    hostPathVolume(
        mountPath: '/var/run/docker.sock', 
        hostPath: '/var/run/docker.sock'
    )
]){    
    node(POD_LABEL){
        stage('build'){
            container('golang'){
                checkout scm
                sh 'docker build . -t app-svc:latest'
            }
        }
    }
}
