
#k8s_yaml(local('helm template --set key1=val1,key2=val2 ./charts/main-chart'))
#watch_file('/charts/main-chart')

local_resource('yarn', cmd='yarn install', deps=['package.json', 'yarn.lock'], labels=['npm'])
local_resource('vault',
    #cmd='helm template vault hashicorp/vault -n vault -f ./k8s/base/helm/vault-values.yaml > ./k8s/base/helm/manifests/vault.yaml',
    # cmd='helm template vault hashicorp/vault -n vault --output-dir ./k8s/base/helm/manifests/',
    cmd="helm template vault hashicorp/vault -n vault --output-dir ./k8s/base/helm/manifests/",
    deps=["/k8s/base/helm/vault-values.yaml"],
    labels=['manifests', "helm", "generator"]
)
local_resource('consul',
    #cmd='helm template vault hashicorp/vault -n vault -f ./k8s/base/helm/vault-values.yaml > ./k8s/base/helm/manifests/vault.yaml',
    # cmd='helm template vault hashicorp/vault -n vault --output-dir ./k8s/base/helm/manifests/',
    cmd="helm template consul hashicorp/consul -n consul --output-dir ./k8s/base/helm/manifests/",
    deps=["/k8s/base/helm/consul-values.yaml"],
    labels=['manifests', "helm", "generator"]
)
local_resource('redis',
    #cmd='helm template vault hashicorp/vault -n vault -f ./k8s/base/helm/vault-values.yaml > ./k8s/base/helm/manifests/vault.yaml',
    # cmd='helm template vault bitnami/vault -n vault --output-dir ./k8s/base/helm/manifests/',
    cmd="helm template redis bitnami/redis -n redis --output-dir ./k8s/base/helm/manifests/",
    deps=["/k8s/base/helm/redis-values.yaml"],
    labels=['manifests', "helm", "generator"]
)
#local_resource('grafana',
#    cmd='helm template grafana bitnami/grafana -n grafana -f ./k8s/base/helm/grafana-values.yaml > ./k8s/base/manifests/grafana.yaml',
#    deps=["/k8s/base/helm/grafana-values.yaml"],
#    labels=['manifests']
#)
# include('./k8s/base/helm/Tiltfile')

# include('./apps/users/api/Tiltfile')
include('./apps/users/client/Tiltfile')
include('./apps/infra/my-kube-controller/Tiltfile')

k8s_yaml(kustomize('k8s/base'))

# k8s_resource("users-api", port_forwards="5001:8080")
# ports to container port that runs as container env var - both ways
# k8s_resource("users-api", port_forwards="5001:8080")
# k8s_resource(workload='users-api', port_forwards="5001:8080")


# GOOS=linux GOARCH=amd64
