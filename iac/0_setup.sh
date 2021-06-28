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
