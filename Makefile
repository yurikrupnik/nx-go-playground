GCP_PROJECT:=$(gcloud config get-value project)
#NEW_SA_NAME=test-service-account-name


define get-secret
$(shell gcloud secrets versions access latest --secret=MONGO_URI --project=mussia14)
endef

get-argocd-secret-yaml:
	kubectl get secret argocd-initial-admin-secret -n argocd -o jsonpath='{.data.password}' | base64 -d
# create service account
csa:
		SA="${NEW_SA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com"
		gcloud iam service-accounts create $NEW_SA_NAME --project $PROJECT_ID
		# enable cloud API
    SERVICE="sqladmin.googleapis.com"
#    gcloud services enable $SERVICE --project $GCP_PROJECT
    # grant access to cloud API
    ROLE="roles/cloudsql.admin"
    gcloud projects add-iam-policy-binding --role="$ROLE" $PROJECT_ID --member "serviceAccount:$SA"

    # create service account keyfile
#    gcloud iam service-accounts keys create creds.json --project $PROJECT_ID --iam-account $SA
#		kubectl create secret generic gcp-creds -n crossplane-system --from-file=creds=./creds.json

minikube-gcp:
	SA="${NEW_SA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com"
	gcloud iam service-accounts create $NEW_SA_NAME --project $PROJECT_ID
	ROLE="roles/cloudsql.admin"
#  gcloud projects add-iam-policy-binding --role="$ROLE" $PROJECT_ID --member "serviceAccount:$SA"
#  gcloud iam service-accounts keys create creds.json --project $PROJECT_ID --iam-account $SA

minikube-up:
	minikube start
	minikube addons enable gcp-auth
	minikube addons enable ingress

#	kubectl create namespace crossplane-system
#	helm repo add crossplane-stable https://charts.crossplane.io/stable
#	helm repo update
#	helm install crossplane --namespace crossplane-system crossplane-stable/crossplane
	#kubectl create secret generic gcp-creds -n crossplane-system --from-file=creds=./creds.json



	# Patches
# does create a secret but does not work with pulling images - todo check
#  kubectl create secret docker-registry docker-registry-key --docker-server=europe-west1-docker.pkg.dev --docker-username=oauth3accesstoken --docker-password="$(gcloud auth print-access-token)" --docker-email=krupnik.yuri@gmail.com
	#kubectl patch serviceaccount default -p '{"imagePullSecrets": [{"name": "docker-registry-key"}]}'
	#kubectl patch serviceaccount default -p '{"imagePullSecrets": [{"name": "gcp-creds"}]}'
	#kubectl patch serviceaccount default -p '{"imagePullSecrets": [{"name": "gcr-access-token"}]}'
	# settings Kubecost
	#kubectl create namespace kubecost
#  helm repo add kubecost https://kubecost.github.io/cost-analyzer/
#  helm install kubecost kubecost/cost-analyzer --namespace kubecost --set kubecostToken="a3J1cG5pay55dXJpQGdtYWlsLmNvbQ==xm343yadf98"
	#kubectl port-forward --namespace kubecost deployment/kubecost-cost-analyzer 9090

	kustomize build k8s/base | kubectl apply -f -

# in cluster
	#docker login -u oauth2accesstoken --password=ya29.A0ARrdaM-9yadNtYAEScIdz0RxjbfSkw79tAg6ZmIlfzkQe1PXnFUcyKsLNqtsw7DgrO-zLuRkkLLWGmhVPSTMqTqncYNLH9kXV8keW1hZ8dt04IOY17jnxJJyCGLSkhU0DpzyxyVSB7Zb2V-Jb5wgHww08sLEGM0ugwi1sQ https://europe-west1-docker.pkg.dev
	#cat .docker/config.json | base64
#	 update then the secrets file with dockerfonfig base64

minikube-down:
	kustomize build k8s/base | kubectl delete -f -
	minikube stop

minikube-delete:
	minikube delete

kind-up:
	#minikube start
	#kubectl create namespace argocd
	kind create cluster
	#k9s
# k8s stop running workspace
kind-down:
	#minikube stop
	kind delete cluster

# k8s delete workspace
kube-reset:
	minikube delete
	minikube start
affected-graph:
	npx nx affected:dep-graph
# NX end
