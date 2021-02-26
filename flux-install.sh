#!/bin/sh

set -e

helm repo add fluxcd https://charts.fluxcd.io

kubectl apply -f https://raw.githubusercontent.com/fluxcd/helm-operator/master/deploy/crds.yaml

kubectl create namespace flux

helm upgrade -i flux fluxcd/flux \
  --namespace flux \
  --set helm.versions=v3 \
  --set git.ciSkip=true \
  --set git.path=deployment \
  --set git.url=ssh://git@github.com/seadiaz/adoption \
  --set syncGarbageCollection.enabled=true
  
kubectl -n flux logs deployment/flux | grep identity.pub | cut -d '"' -f2