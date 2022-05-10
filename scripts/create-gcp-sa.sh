# replace this with your own gcp project id and the name of the service account
# that will be created.
PROJECT_ID=$(gcloud config get-value project)
NEW_SA_NAME=test-service-account-name
SA="${NEW_SA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com"
gcloud iam service-accounts create $NEW_SA_NAME --project $PROJECT_ID
ROLE="roles/cloudsql.admin"
gcloud projects add-iam-policy-binding --role="$ROLE" $PROJECT_ID --member "serviceAccount:$SA"

#echo $NEW_SA_NAME
#OUTPUT=$(ls -1)
#echo "${OUTPUT}"
echo $PROJECT_ID
#echo $PROJECT_ID
# create service account
#SA="${NEW_SA_NAME}@${PROJECT_ID}.iam.gserviceaccount.com"
#gcloud iam service-accounts create $NEW_SA_NAME --project $PROJECT_ID
#
## enable cloud API
#SERVICE="sqladmin.googleapis.com"
#gcloud services enable $SERVICE --project $PROJECT_ID
#
## grant access to cloud API
#ROLE="roles/cloudsql.admin"
#gcloud projects add-iam-policy-binding --role="$ROLE" $PROJECT_ID --member "serviceAccount:$SA"
#
## create service account keyfile
#gcloud iam service-accounts keys create creds.json --project $PROJECT_ID --iam-account $SA
