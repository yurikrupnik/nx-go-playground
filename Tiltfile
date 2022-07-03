load('ext://helm_resource', 'helm_resource', 'helm_repo')

helm_repo('bitnami', 'https://charts.bitnami.com/bitnami', labels=["helm"])
helm_repo('devtron', 'https://helm.devtron.ai', labels=["helm"])
helm_repo('hashicorp', 'https://helm.releases.hashicorp.com', labels=["helm"])
# helm_repo('external-secrets', 'https://charts.external-secrets.io')
# helm_repo('bitnami', 'https://charts.bitnami.com/bitnami')
load('ext://helm_remote', 'helm_remote')
#helm_remote('mysql',
#            repo_name='bitnami',
#            namespace='mysql',
#            create_namespace="true",
#            repo_url='https://charts.bitnami.com/bitnami')
# helm_resource('vault', 'hashicorp/vault')
# helm_resource('kube-prometheus-stack', 'devtron/kube-prometheus-stack')


# Start of cert-manager
# different way - shows in helm charts
# but creates namespace and can pass values file
# helm_resource('cert-manager', 'bitnami/cert-manager')
#helm_remote('cert-manager',
#            repo_name='cert-manager',
#            namespace='cert-manager',
#            create_namespace="true",
#            repo_url='https://charts.bitnami.com/bitnami')
# End of cert-manager

# Start of external-secrets
# helm_resource('secrets', 'external-secrets/external-secrets')
#helm_remote('external-secrets',
#           repo_name='external-secrets',
#           namespace='external-secrets',
#           create_namespace="true",
#           repo_url='https://charts.external-secrets.io')
# End of external-secrets

# Prometheus start
# helm_resource('prometheus', 'bitnami/kube-prometheus')
# todo does not work
#helm_remote('prometheus',
#            repo_name='kube-prometheus',
#            namespace='kube-prometheus',
#            create_namespace="true",
#            repo_url='https://charts.bitnami.com/bitnami')
#k8s_resource("prometheus", port_forwards="9090:9090")
# Grafana end

# Grafana start
#helm_remote('grafana',
#            repo_name='grafana',
#            namespace='grafana',
#            create_namespace="true",
#            repo_url='https://charts.bitnami.com/bitnami')
# k8s_resource("grafana", port_forwards="3000:3000")
# Grafana end

# Redis start
# no namespace and values
# helm_resource('redis', 'bitnami/redis')
#helm_remote('redis',
#            repo_name='redis',
#            namespace='redis',
#            create_namespace="true",
#            repo_url='https://charts.bitnami.com/bitnami')
# k8s_resource("redis", port_forwards="3000:3000")
# Redis end
#helm_remote('kube-prometheus',
 #           repo_name='kube-prometheus',
  #          namespace='kube-prometheus',
   #         create_namespace="true",
    #        repo_url='https://charts.bitnami.com/bitnami')

# helm_resource('grafana', 'bitnami/grafana')
# k8s_resource("kube-prometheus", port_forwards="9090:9090")

# helm_resource('mongo', 'bitnami/mongodb-sharded')

# k8s_yaml(helm('./charts/main-chart', name="main-chart",namespace="aris"))


#k8s_yaml(local('helm template --set key1=val1,key2=val2 ./charts/main-chart'))
#watch_file('/charts/main-chart')

local_resource('yarn', cmd='yarn install', deps=['package.json', 'yarn.lock'], labels=['npm'])
local_resource('helm', cmd='helm -h', deps=["."], labels=['manifests'])

include('./apps/users/api/Tiltfile') # labels=["users", "api"]
#include('./apps/users/api/Tiltfile', labels: ["aris"])
include('./apps/users/client/Tiltfile') # labels=["users", "client"]
# include('./apps/infra/my-kube-controller/Tiltfile')

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
