#!/bin/bash

# Define names
POD_NAME="quantix-pod"
APP_NAME="quantix-app"
APP_IMAGE="quay.io/cmbsolver/quantix"

echo "Cleaning up existing pod and containers..."
podman pod rm -f $POD_NAME 2>/dev/null

echo "Pulling latest images..."
podman pull $APP_IMAGE

echo "Creating the pod..."
# Exposing 3301 for the web app
podman pod create --name $POD_NAME -p 3301:3301

echo "Starting the Quantix application container in the pod..."
podman run -d \
  --name $APP_NAME \
  --pod $POD_NAME \
  $APP_IMAGE

echo "Deployment complete."
podman pod ps
podman ps --filter pod=$POD_NAME