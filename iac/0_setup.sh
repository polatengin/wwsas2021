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

AVAILABLE_AKS_VERSION=`az aks get-versions --location "${LOCATION}" --query "orchestrators[-1].orchestratorVersion" --output tsv`
