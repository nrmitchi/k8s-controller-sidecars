> The base of this repo was originally forked from github.com/trstringer/k8s-controller-core-resource

# Kubernetes Custom Controller - Sidecar Shutdown


## What is this?

This is a custom Kubernetes controller for the purpose of watching running pods, and sending a SIGTERM to sidecar containers when the "main" application container has exited (and the sidecars are the only non-terminated containers). 

This is a response to https://github.com/kubernetes/kubernetes/issues/25908.

## Usage

1. Deploy the controller into your cluster.
1. Add the `nrmitchi.com/sidecars` annotation to your pods, with a comma-seperated list of sidecar container names. 
