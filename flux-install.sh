#!/bin/sh

set -e

helm repo add fluxcd https://charts.fluxcd.io

kubectl apply -f https://raw.githubusercontent.com/fluxcd/helm-operator/master/deploy/crds.yaml

kubectl create namespace flux

helm upgrade -i flux fluxcd/flux \
  --set git.url=git@github.com/seadiaz/adoption \
  --namespace flux \
  --set helm.versions=v3 \
  --set git.path=/deployment \
  --set git.readonly=false
  