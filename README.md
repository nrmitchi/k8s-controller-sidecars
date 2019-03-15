# Kubernetes Custom Controller - Sidecar Shutdown

Kubernetes (cron)jobs sidecar terminator.
Originally forked from https://github.com/nrmitchi/k8s-controller-sidecars .

## What is this?

This is a custom Kubernetes controller for the purpose of watching running pods, and sending a SIGTERM to sidecar containers when the "main" application container has exited (and the sidecars are the only non-terminated containers).

This is a response to https://github.com/kubernetes/kubernetes/issues/25908.

## Usage

1. Deploy the controller into your cluster

```sh
kubectl apply -f manifest.yml
```

1. Add the `lemonade.com/sidecars` annotation to your pods, with a comma-seperated list of sidecar container names.

Example:

```yaml
---
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: test-job
spec:
  schedule: "* * * * *"
  jobTemplate:
    spec:
      template:
        metadata:
          annotations:
            lemonade.com/sidecars: logging
        spec:
          restartPolicy: Never
          containers:
            - name: test-job
              image: ubuntu:latest
              command: ["sleep", "5"]
            - name: logging
              image: fluentd
```
