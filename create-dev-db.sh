echo "Cleaning up existing pod and containers..."
podman pod rm -f quantix-pod 2>/dev/null

echo "cleaning up the containers"
podman stop postgres

podman rm -f postgres

podman rmi -f docker.io/library/postgres

podman volume rm --all
podman volume prune --force

echo "Pulling the images"
podman pull docker.io/library/postgres

echo "Bringing up the db containers"
podman run --name postgres -e POSTGRES_PASSWORD=quantixpw -dt -p 5432:5432 docker.io/library/postgres

echo "seeing processes"
podman ps -a