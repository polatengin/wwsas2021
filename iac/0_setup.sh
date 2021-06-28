PROJECT_NAME="wwsas2021"
LOCATION="northeurope"

AZURE_CLI_EXISTS=$(command -v "az")
if [[ -z ${AZURE_CLI_EXISTS} ]]; then
  echo "Installing Azure CLI..."

  curl -sL https://aka.ms/InstallAzureCLIDeb | bash
fi

LOGIN_COUNT=`az account list --query "length([])"`

if [[ ${LOGIN_COUNT} -eq 0 ]]
then
  az login
fi

echo "Creating Resource Group..."

az group create --name "${PROJECT_NAME}-rg" --location "${LOCATION}" --output none

echo "Creating Azure Container Registry..."

az acr create --name "${PROJECT_NAME}acr" --resource-group "${PROJECT_NAME}-rg" --sku Basic --admin-enabled true --output none

echo "Getting ACR_ID..."

ACR_ID=`az acr show --name "${PROJECT_NAME}acr" --resource-group "${PROJECT_NAME}-rg" --query "id" --output tsv`

echo "Creating Azure Kubernetes Service..."

AVAILABLE_AKS_VERSION=`az aks get-versions --location "${LOCATION}" --query "orchestrators[-1].orchestratorVersion" --output tsv`

az aks create --name "${PROJECT_NAME}-aks" --resource-group "${PROJECT_NAME}-rg" --attach-acr "$ACR_ID" --kubernetes-version "${AVAILABLE_AKS_VERSION}" --no-ssh-key --output none

echo "Setting Azure Kubernetes Service Credentials..."

kubectl config unset users
kubectl config unset current-context
kubectl config unset contexts
kubectl config unset clusters

az aks get-credentials --name "${PROJECT_NAME}-aks" --resource-group "${PROJECT_NAME}-rg" --output none

ACR_LOGIN_SERVER=`az acr show --name "${PROJECT_NAME}acr" --resource-group "${PROJECT_NAME}-rg" --query "loginServer" --output tsv`

DOCKER_LOGGEDIN=`cat ~/.docker/config.json | jq ".auths" | grep "${ACR_LOGIN_SERVER}" | sed "s/://" | sed 's/"//g' | sed "s/{//" | tr -d "[:space:]"`

if [ "${DOCKER_LOGGEDIN}" == "" ]
then
  echo "Logging-in to Azure Container Registry through Docker CLI"
  az acr login --name "${PROJECT_NAME}acr"
fi

echo "Creating a Namespace in Kubernetes for this project"

kubectl create namespace "project" --dry-run=client -o yaml | kubectl apply -f -
