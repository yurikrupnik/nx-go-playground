GCP_PROJECT:=$(gcloud config get-value project)

define get-secret
$(shell gcloud secrets versions access latest --secret=MONGO_URI --project=mussia14)
endef

affected-graph:
	npx nx affected:dep-graph
# NX end
