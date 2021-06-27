PROJECT_NAME="wwsas2021"
LOCATION="northeurope"

AZURE_CLI_EXISTS=$(command -v "az")
if [[ -z ${AZURE_CLI_EXISTS} ]]; then
  echo "Installing Azure CLI..."

  curl -sL https://aka.ms/InstallAzureCLIDeb | bash
fi
