#!/bin/bash

kubectl apply -f ./k8s/granitedb-deployment.yaml
kubectl apply -f ./k8s/granitedb-pvc.yaml
kubectl apply -f ./k8s/granite-deployment.yaml
kubectl apply -f ./k8s/granite-config.yaml
kubectl apply -f ./k8s/granite-secret.yaml
