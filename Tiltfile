
#k8s_yaml(local('helm template --set key1=val1,key2=val2 ./charts/main-chart'))
#watch_file('/charts/main-chart')

local_resource(
  "make k-b-a",
  #cmd="make k-b-a",
  cmd="ls",
  allow_parallel = True,
  trigger_mode=TRIGGER_MODE_MANUAL,
  # only=["/k8s/base/helm/values/"],
  deps=["/k8s/base/helm/values/"],
  labels=["makefile", "helm", "manual"],
)
local_resource(
  "make k-b-d",
  cmd="ls",
  #cmd="make k-b-d",
  allow_parallel = True,
  trigger_mode=TRIGGER_MODE_MANUAL,
  # only=["/k8s/base/helm/values/"],
  deps=["/k8s/base/helm/values/"],
  labels=["makefile", "helm", "manual"],
)

local_resource('yarn', cmd='yarn install', deps=['package.json', 'yarn.lock'], labels=['npm'])

#local_resource('vault-templates',
#    cmd="helm template vault hashicorp/vault -n vault -f ./k8s/base/helm/values/vault-values.yaml > ./k8s/base/helm/manifests/vault1.yaml",
#    deps=["/k8s/base/helm/vault-values.yaml"],
#    labels=['manifests', "helm", "generator"]
#)
#local_resource('vault-templates',
#    cmd="helm template vault hashicorp/vault -n vault -f ./k8s/base/helm/values/vault-values.yaml > ./k8s/base/helm/manifests/vault.yaml",
#    deps=["/k8s/base/helm/vault-values.yaml"],
#    labels=['manifests', "helm", "generator"]
#)
#local_resource('consul-template',
#    cmd="helm template consul hashicorp/consul -n consul -f ./k8s/base/helm/values/consul-values.yaml > ./k8s/base/helm/manifests/consul.yaml",
#    deps=["/k8s/base/helm/consul-values.yaml"],
#    labels=['manifests', "helm", "generator", "self-test"]
#)

# using kubctl aplly command
#local_resource('redis-template',
 #   cmd="helm template redis bitnami/redis -n redis -f ./k8s/base/helm/values/redis-values.yaml> ./k8s/base/helm/manifests/redis.yaml",
 #   deps=["/k8s/base/helm/values/redis-values.yaml"],
 #   labels=['manifests', "helm", "generator", "self-test"]
#)
#local_resource('prometheus-template',
#    cmd="helm template prometheus bitnami/kube-prometheus -n prometheus -f k8s/base/helm/values/prometheus-values.yaml > ./k8s/base/helm/manifests/prometheus.yaml",
#    deps=["/k8s/base/helm/values/prometheus-values.yaml"],
#    labels=['manifests', "helm", "generator", "self-test1"]
#)
#local_resource('grafana-template',
#    cmd="helm template grafana bitnami/grafana -n grafana -f k8s/base/helm/values/grafana-values.yaml --output-dir ./k8s/base/helm/manifests/",
 #   deps=["/k8s/base/helm/values/grafana-values.yaml"],
 #   labels=['manifests', "helm", "generator", "self-test1"]
#)
#local_resource('grafana',
#    cmd='helm template grafana bitnami/grafana -n grafana -f ./k8s/base/helm/grafana-values.yaml > ./k8s/base/manifests/grafana.yaml',
#    deps=["/k8s/base/helm/grafana-values.yaml"],
#    labels=['manifests']
#)
include('./k8s/base/helm/Tiltfile')

include('./apps/users/api/Tiltfile')
include('./apps/users/client/Tiltfile')
include('./apps/infra/my-kube-controller/Tiltfile')

k8s_yaml(kustomize('k8s/base'))
# k8s_yaml(kustomize('k8s/base/helm/manifests'))

# k8s_resource("users-api", port_forwards="5001:8080")
# ports to container port that runs as container env var - both ways
# k8s_resource("users-api", port_forwards="5001:8080")
# k8s_resource(workload='users-api', port_forwards="5001:8080")


# GOOS=linux GOARCH=amd64
