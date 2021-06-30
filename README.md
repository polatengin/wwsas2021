# WorldWide Software Architecture Summit 2021

[WorldWide Software Architecture Summit 2021](https://geekle.us/software_architecture) event dates are _August 3rd, 2021_ and _August 4th, 2021_

As a speaker in the event, I prepared this repo as a guideline to build a solution with following requirements;

* Built with microservices approach in mind
* Built with several components, such as, a React frontend, a .Net 5 backend api, a Go backend api, a NodeJs backend api
* All the components must be containerized with Docker
* Fully automated using GitHub Actions
* Published to Azure
* No vendor lock, must be publishable to other cloud providers
* All components must live inside of a Kubernetes cluster
* Easy scalability using Kubernetes deployments

I'd like to see you in my session ([Building and Automating a solution with Microservices approaches using .Net, Go, Node, Kubernetes and GitHub Actions](https://geekle.us/software_architecture)), so we can discuss the solution even further.

## Development Environment

For the project, we're using DevContainers to define the development environment. GitHub Codespaces are using DevContainers under-the-hood, so, one stone, two birds. We'll have both local and remote environments super easily, with the help of DevContainers.

You can find more detailed explanation of DevContainers here and here

Let's start building the development environment

Create devcontainer.json file under .devcontainer folder at the root of the project

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

  _In this example, we're starting login process of the GitHub CLI tool. So we ensure that the development environment has the GitHub CLI logged in after it's created_

  ```json
  "postCreateCommand": "gh auth login --web",
  ```
