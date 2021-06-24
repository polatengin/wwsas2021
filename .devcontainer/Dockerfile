FROM golang:1.16.5-buster

# Avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Configure apt and install packages
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils dialog libicu63 libssl1.1 2>&1 \
    #
    # Verify wget, jq, unzip, git and process tools installed
    && apt-get -y install git iproute2 procps wget unzip jq \
    #
    # Install Docker CE CLI
    && apt-get install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common lsb-release \
    && curl -fsSL https://download.docker.com/linux/$(lsb_release -is | tr '[:upper:]' '[:lower:]')/gpg | (OUT=$(apt-key add - 2>&1) || echo $OUT) \
    && add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/$(lsb_release -is | tr '[:upper:]' '[:lower:]') $(lsb_release -cs) stable" \
    && apt-get update \
    && apt-get install -y docker-ce-cli \
    #
    # Install Node
    && curl -sL https://deb.nodesource.com/setup_15.x | bash - \
    && apt-get install -y nodejs \
    #
    # Install DotNet 5
    && wget https://packages.microsoft.com/config/debian/10/packages-microsoft-prod.deb -O packages-microsoft-prod.deb \
    && dpkg -i packages-microsoft-prod.deb \
    && apt-get update \
    && apt-get install -y dotnet-sdk-5.0 \
    #
    # Install global packages
    && npm install npm-check-updates -g \
    #
    # Install Azure CLI
    && curl -sL https://aka.ms/InstallAzureCLIDeb | bash \
    #
    # Install kubectl
    && curl -sL https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - \
    && echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | tee -a /etc/apt/sources.list.d/kubernetes.list \
    && apt-get update \
    && apt-get install -y kubectl \
    #
    # Install Helm
    && curl -sL https://git.io/get_helm.sh | bash \
    #
    # Install Terraform
    && CURR_VER=$(curl -s https://checkpoint-api.hashicorp.com/v1/check/terraform | jq -r -M '.current_version') \
    && wget https://releases.hashicorp.com/terraform/${CURR_VER}/terraform_${CURR_VER}_linux_amd64.zip \
    && unzip terraform_${CURR_VER}_linux_amd64.zip -d /usr/local/bin/ \
    && rm -rf terraform_${CURR_VER}_linux_amd64.zip \
    # Clean up
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

# Switch back to dialog for any ad-hoc use of apt-get
ENV DEBIAN_FRONTEND=

CMD ["bash"]
