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
        stage('build and release'){
            container('docker'){
                checkout scm
                def BUILD_VERSION = new Date().format("y.M.d")+"-${BRANCH_NAME}-${BUILD_NUMBER}".replace("/","")
                sh "docker build --build-arg BUILD_VERSION=${BUILD_VERSION} . -t app-svc:${BUILD_VERSION}"
                writeYaml file: 'version.yml', data: ['version': BUILD_VERSION], overwrite: true
            }
        }
        stage('deploy'){
             withCredentials([usernamePassword(credentialsId: 'argocd-appsvc', passwordVariable: 'ARGOCD_TOKEN', usernameVariable: 'ARGOCD_USERNAME')]) {
                container('argocd'){
                    def ENV="dev"
                    if (!BRANCH_NAME.contains('PR')){           
                        sh "argocd app create app-svc-${ENV} --repo=https://github.com/mando04/app-svc.git --path=deploy/helm/app-svc --dest-namespace=app --dest-server=https://kubernetes.docker.internal:6443 --insecure --auth-token=${ARGOCD_TOKEN} --revision=${BRANCH_NAME} --server=argocd-server.argocd --plaintext --helm-set=image.tag=${BRANCH_NAME} --upsert"
                        sh "argocd app sync app-svc-${ENV} --insecure --auth-token ${ARGOCD_TOKEN} --server=argocd-server.argocd --plaintext"
                    }
                }
            }
        }
    }
}
