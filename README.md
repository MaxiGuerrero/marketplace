# Marketplace monorepo by Maxi

This is my personal project, where i use all my knowladge as a Developer and Operator. (Full stack Dev/Ops).

![Marketplace](https://drive.google.com/uc?export=view&id=1uIv4jQsAbbrJipk6CIWIrDKXdCVtbE28)

As developer:

I designed a basic marketplace using a "microservices" architecture.
Into each API, i use the code "hexagonal" + Vertical sliceing architecture, because it provide an clean organization, implementation abstraction and easly for make unit test in my code.

As Operator

I designed a solution of how to connect all those services and which infrastructure we need to implement this solution.

I used docker to make the service's images and docker compose to implement all services in a VM or localhost machine.

Also, i used the kubernetes implementation to up all services into a Kubernetes cluster either in cloud enviroment or in localhost machine using minikube.

As DevOps automater

Using github action to make CI automations, i made workflows for each API services.

The jobs of the workflow are:

1. Validate code with an linter.
2. Run unit tests.
3. If 1 and 2 it's ok, run create docker image of the service and upload into the repository registry.

![Workflow](https://drive.google.com/uc?export=view&id=1C3JI1vQgqH4pOlZiuojcZ0pV7bJwVHI7)
