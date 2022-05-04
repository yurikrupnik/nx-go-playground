GCP_PROJECT:=$(gcloud config get-value project)

define get-secret
$(shell gcloud secrets versions access latest --secret=MONGO_URI --project=mussia14)
endef

ku:
	#minikube start
	#kubectl create namespace argocd
	kind create cluster
	#k9s
# k8s stop running workspace
kd:
	minikube stop
	kind delete cluster

# k8s delete workspace
kube-reset:
	minikube delete
	minikube start
affected-graph:
	npx nx affected:dep-graph
# NX end
