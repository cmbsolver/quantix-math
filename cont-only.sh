#!/bin/bash

# Define names
POD_NAME="quantix-pod"
DB_NAME="postgres-db"
APP_NAME="quantix-app"
APP_IMAGE="quay.io/cmbsolver/quantix"
DB_IMAGE="docker.io/library/postgres"

echo "Cleaning up existing pod and containers..."
podman pod rm -f $POD_NAME 2>/dev/null

echo "Pulling latest images..."
podman pull $APP_IMAGE
podman pull $DB_IMAGE

echo "Creating the pod..."
# Exposing 3301 for the web app and 5432 for DB access from host if needed
podman pod create --name $POD_NAME -p 3301:3301 -p 5432:5432

echo "Starting the Postgres container in the pod..."
podman run -d \
  --name $DB_NAME \
  --pod $POD_NAME \
  -e POSTGRES_PASSWORD=quantixpw \
  $DB_IMAGE

echo "Letting the db fully set up"
sleep 2m

echo "Starting the Quantix application container in the pod..."
# Since they share the pod, the app connects to DB via localhost:5432
podman run -d \
  --name $APP_NAME \
  --pod $POD_NAME \
  $APP_IMAGE

echo "Deployment complete."
podman pod ps
podman ps --filter pod=$POD_NAME