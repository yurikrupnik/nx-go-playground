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
