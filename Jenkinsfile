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
                sh 'docker build --build-arg BUILD_VERSION=${BRANCH_NAME} . -t app-svc:${BRANCH_NAME}'
            }
        }
        stage('deploy'){
            container('argocd'){
                if (!BRANCH_NAME.contains('PR')){
                    withCredentials([usernamePassword(credentialsId: 'argocd-appsvc', passwordVariable: 'ARGOCD_TOKEN', usernameVariable: 'ARGOCD_USERNAME')]) {
                        def env = 'dev'
                        sh 'argocd app create app-svc --repo https://github.com/mando04/app-svc.git --path deploy/helm/app-svc --dest-namespace app --dest-server https://kubernetes.docker.internal:6443 --insecure --grpc-web --auth-token ${ARGOCD_TOKEN} --revision ${BRANCH_NAME}'
                        sh 'argocd app sync app-svc --insecure --grpc-web --auth-token ${ARGOCD_TOKEN}'
                    }
                }
            }
        }
    }
}
