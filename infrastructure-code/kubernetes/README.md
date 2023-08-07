# Installation

## Requeriments

1. Kubectl binary installed. https://kubernetes.io/docs/tasks/tools/
2. Have access on your kubernetes cluster.
3. Kustomize binary installed on your computer. https://kustomize.io/
4. Have installed an ingress controller on your kubernetes cluster.

## Steps to follow

1. Go to the directory infrastructure-code/kubernetes.
2. Configure route.ingress.yaml with your ingress controller annotations.
2. Run command:
        
        kubectl apply -k .