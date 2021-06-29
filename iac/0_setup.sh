PROJECT_NAME="wwsas2021"
LOCATION="northeurope"

AZURE_CLI_EXISTS=$(command -v "az")
if [[ -z ${AZURE_CLI_EXISTS} ]]; then
  echo "Installing Azure CLI..."

  curl -sL https://aka.ms/InstallAzureCLIDeb | bash
fi

GITHUB_CLI_EXISTS=$(command -v "gh")
if [[ -z ${GITHUB_CLI_EXISTS} ]]; then
  echo "Installing GitHub CLI..."

  curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | gpg --dearmor -o /usr/share/keyrings/githubcli-archive-keyring.gpg
  echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | tee /etc/apt/sources.list.d/github-cli.list > /dev/null
  apt update
  apt install gh

  gh auth login --web
fi

LOGIN_COUNT=`az account list --query "length([])"`

if [[ ${LOGIN_COUNT} -eq 0 ]]
then
  az login --output none
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

echo "Installing nginx as the Ingress Controller into the Kubernetes"

kubectl create namespace ingress-basic

helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx

helm install nginx-ingress ingress-nginx/ingress-nginx \
  --namespace ingress-basic \
  --set controller.replicaCount=2

echo "Configuring nginx"

kubectl apply -f ./1_ingress.yml

echo "Setting a GitHub Secret for GitHub Actions"

AZURE_CREDENTIALS=$(az ad sp create-for-rbac --name wwsas2021 --sdk-auth)

gh secret set "AZURE_CREDENTIALS" -b "${AZURE_CREDENTIALS}"
