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
