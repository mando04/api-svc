podTemplate(
cloud: 'kubernetes',
containers: [
    containerTemplate(name: 'docker', image: 'docker:19.03.1-dind', ttyEnabled: true, command: 'cat'),
    containerTemplate(name: 'argocd', image: 'argocd-cli:custom', ttyEnabled: true, command: 'cat'),
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
            environment {
                ENV = "dev"
            }
             withCredentials([usernamePassword(credentialsId: 'argocd-appsvc', passwordVariable: 'ARGOCD_TOKEN', usernameVariable: 'ARGOCD_USERNAME')]) {
                container('argocd'){
                    if (!BRANCH_NAME.contains('PR')){           
                        sh "argocd app create app-svc-${ENV} --repo=https://github.com/mando04/app-svc.git --path=deploy/helm/app-svc --dest-namespace=app --dest-server=https://kubernetes.docker.internal:6443 --insecure --auth-token=${ARGOCD_TOKEN} --revision=${BRANCH_NAME} --server=argocd-server.argocd --plaintext"
                        sh "argocd app sync app-svc-${ENV} --insecure --auth-token ${ARGOCD_TOKEN} --server=argocd-server.argocd --plaintext"
                    }
                }
            }
        }
    }
}
