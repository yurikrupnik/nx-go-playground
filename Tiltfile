load('ext://helm_resource', 'helm_resource', 'helm_repo')

helm_repo('bitnami', 'https://charts.bitnami.com/bitnami')
helm_repo('devtron', 'https://helm.devtron.ai')
# helm_repo('hashicorp', 'https://helm.releases.hashicorp.com')
# helm_repo('external-secrets', 'https://charts.external-secrets.io')
# helm_repo('bitnami', 'https://charts.bitnami.com/bitnami')
load('ext://helm_remote', 'helm_remote')
#helm_remote('mysql',
 #           repo_name='bitnami',
  #          namespace='mysql',
   #         create_namespace="true",
    #        repo_url='https://charts.bitnami.com/bitnami')
# helm_resource('vault', 'hashicorp/vault')
# helm_resource('kube-prometheus-stack', 'devtron/kube-prometheus-stack')
helm_resource('secrets', 'external-secrets/external-secrets')
helm_resource('cert-manager', 'bitnami/cert-manager')

helm_resource('prometheus', 'bitnami/kube-prometheus')

helm_remote('grafana',
            repo_name='grafana',
            namespace='grafana',
            create_namespace="true",
            repo_url='https://charts.bitnami.com/bitnami')
# helm_remote('prometheus',
#            repo_name='kube-prometheus',
#            namespace='prometheus',
#            create_namespace="true",
#            repo_url='https://charts.bitnami.com/bitnami')

# helm_resource('grafana', 'bitnami/grafana')
k8s_resource("prometheus", port_forwards="9090:9090")
k8s_resource("grafana", port_forwards="3000:3000")

# helm_resource('mongo', 'bitnami/mongodb-sharded')
helm_resource('redis', 'bitnami/redis')

# k8s_yaml(helm('./charts/main-chart', name="main-chart"))


# k8s_yaml(local('helm template --set key1=val1,key2=val2 ./charts/main-chart'))
# watch_file('/charts/main-chart')

local_resource('yarn', cmd='yarn install', deps=['package.json', 'yarn.lock'])

include('./apps/users/api/Tiltfile')
include('./apps/users/client/Tiltfile')

#local_resource(
 # 'build-users-api',
 # cmd='yarn nx run users-api:build',
 # deps=['./apps/users/api', './libs/go'],
 # resource_deps=["yarn"],
 # env={"GOOS":"linux","GOARCH":"amd64"},
 # ignore=["dist/apps/users/api", "node_modules"]
#)

# docker_build(
 # "yurikrupnik/users-api",
 # ".",
 # target="go-builder",
 # build_args={"DIST_PATH":"dist/apps/users/api"},
 # only=["dist/apps/users/api"]
#)
k8s_yaml(kustomize('k8s/base'))

# k8s_resource("users-api", port_forwards="5001:8080")
# ports to container port that runs as container env var - both ways
# k8s_resource("users-api", port_forwards="5001:8080")
# k8s_resource(workload='users-api', port_forwards="5001:8080")


# GOOS=linux GOARCH=amd64
