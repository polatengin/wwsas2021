# WorldWide Software Architecture Summit 2021

[WorldWide Software Architecture Summit 2021](https://geekle.us/software_architecture) event dates are _August 3rd, 2021_ and _August 4th, 2021_

As a speaker in the event, I prepared this repo as a guideline to build a solution with following requirements;

* Built with microservices approach in mind
* Built with several components, such as, a React frontend, a .Net 5 backend api, a Go backend api, a NodeJs backend api
* All the components must be containerized with Docker
* All components must live inside of a Kubernetes cluster
* Easy scalability using Kubernetes deployments
* Published to Azure
* No vendor lock, must be publishable to other cloud providers
* Fully automated using GitHub Actions

I'd like to see you in my session ([Building and Automating a solution with Microservices approaches using .Net, Go, Node, Kubernetes and GitHub Actions](https://geekle.us/software_architecture)), so we can discuss the solution even further.

Index

* [Development Environment](#development-environment)
* [Infrastructure as Code](#infrastructure-as-code)

## Development Environment

For the project, we're using DevContainers to define the development environment. GitHub Codespaces are using DevContainers under-the-hood, so, one stone, two birds. We'll have both local and remote environments super easily, with the help of DevContainers.

You can find more detailed explanation of DevContainers here and here

Let's start building the development environment

* Create [devcontainer.json](./.devcontainer/devcontainer.json) file under [.devcontainer](./.devcontainer/) folder at the root of the project

  We can define;

  name of the development environment

  ```json
  "name": "Development Environment",
  ```

  settings that are gonna applied to Visual Studio Code

  ```json
  "settings": {
    "terminal.integrated.profiles.linux": {
      "bash": {
        "path": "/bin/bash"
      }
    },
    "workbench.iconTheme": "vscode-icons"
  },
  ```

  extensions that are gonna installed to Visual Studio

  ```json
  "extensions": [
    "ms-dotnettools.csharp",
    "golang.go",
    "ms-vscode.vscode-node-azure-pack",
    "durablefunctionsmonitor.durablefunctionsmonitor",
    "ms-azuretools.vscode-docker",
    "editorconfig.editorconfig",
    "vscode-icons-team.vscode-icons",
    "humao.rest-client"
  ]
  ```

  forwarded ports from inside of the DevContainer to host machine

  ```json
  "forwardPorts": [ 5000 ],
  ```

  bash commands that is gonna run after DevContainer created

  PS: _In this example, we're starting login process of the GitHub CLI tool. So we ensure that the development environment has the GitHub CLI logged in after it's created_

  ```json
  "postCreateCommand": "gh auth login --web",
  ```

* Create [Dockerfile](./.devcontainer/Dockerfile) under the [.devcontainer](./.devcontainer/) folder

  We can start with the image of the framework we used in one of the _main_ projects, such as, [GoLang Docker Image](https://hub.docker.com/_/golang/), [Node Docker Image](https://hub.docker.com/_/node), [.Net SDK Docker Image](https://hub.docker.com/_/microsoft-dotnet-sdk/)

  In this solution, we're gonna start from [GoLang Docker Image](https://hub.docker.com/_/golang/)

  ```docker
  FROM golang:1.16.5-buster
  ```

  PS: _At the time of writing this guideline, latest go version is 1.16.5, there may be newer versions when you read this_

  In this [Dockerfile](./.devcontainer/Dockerfile), we're gonna install and configure all the tools, libraries, sdks that we'll use through-out the development, such as;

  * [Go SDK](https://github.com/polatengin/wwsas2021/blob/main/.devcontainer/Dockerfile#L1)
  * [Node SDK](https://github.com/polatengin/wwsas2021/blob/main/.devcontainer/Dockerfile#L24)
  * [.Net 5 SDK](https://github.com/polatengin/wwsas2021/blob/main/.devcontainer/Dockerfile#L28)
  * [Azure CLI](https://github.com/polatengin/wwsas2021/blob/main/.devcontainer/Dockerfile#L37)
  * [GitHub CLI](https://github.com/polatengin/wwsas2021/blob/main/.devcontainer/Dockerfile#L40)
  * [Docker CLI](https://github.com/polatengin/wwsas2021/blob/main/.devcontainer/Dockerfile#L17)
  * [Helm](https://github.com/polatengin/wwsas2021/blob/main/.devcontainer/Dockerfile#L52)
  * [Terraform](https://github.com/polatengin/wwsas2021/blob/main/.devcontainer/Dockerfile#L58)
  * [kubectl](https://github.com/polatengin/wwsas2021/blob/main/.devcontainer/Dockerfile#L46)

* Create [.editorconfig](./.editorconfig) file at the root of the solution

```ini
root = true

[*]
charset = utf-8
indent_style = space
indent_size = 2
insert_final_newline = true
trim_trailing_whitespace = true
```

* Create [.gitignore](./.gitignore) file at the root of the solution and ignore unneeded files from the git system, for example,

```ìni
bin/
obj/
Properties/

node_modules/
dist/
package-lock.json
bundle.css
```

## Infrastructure as Code

[./iac/setup.sh](./iac/setup.sh) script file includes all the Bash code to do below checks, provision below resources on Azure and complete the setup by executing below;

* Check if Azure CLI is exists, if not, install Azure CLI
* Check if GitHub CLI is exists, if not, install GitHub CLI and proceed to login
* If Azure CLI is not loggedin yet, proceed to login to Azure
* Create a Resource Group on Azure
* Create an Azure Container Registry (ACR), to hold project specific Docker Images
* Create an Azure Kubernetes Services (AKS) linked to the Azure Container Registry (ACR), to have a compute power to run project Docker Images
* Reset kubectl config and set current context to Azure Kubernetes Services (AKS) instance
* Create new Service Principal and set it as secret to the GitHub CLI
* Install nginx into the Azure Kubernetes Services (AKS) instance as Ingress Controller to handle incoming traffic
* Apply ingress.yml config onto the Ingress Controller to associate endpoints to services

```yaml
- path: /api/product/(.*)
  pathType: Prefix
  backend:
    service:
      name: product-service
      port:
        number: 80
- path: /api/campaign/(.*)
  pathType: Prefix
  backend:
    service:
      name: campaign-service
      port:
        number: 80
- path: /api/user/(.*)
  pathType: Prefix
  backend:
    service:
      name: user-service
      port:
        number: 80
- path: /(.*)
  pathType: Prefix
  backend:
    service:
      name: frontend-service
      port:
        number: 80
```

## Projects

There are 4 projects in the solution;

* api-campaing: NodeJs project that handles /api/campaign/* requests
* api-product: GoLang project that handles /api/product/* requests
* api-user: .Net 5 project that handles /api/user/* requests
* web-frontend: React project that renders frontend
