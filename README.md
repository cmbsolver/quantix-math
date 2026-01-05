# quantix-math
A set of math utilities for solving the Cicada 3301 puzzle. 

## Getting Started

### Scripts

The project includes two shell scripts to help with environment setup and deployment using Podman:

#### `create-dev-db.sh`
Use this script to set up a clean PostgreSQL instance for local development.
- **What it does**: Cleans up existing containers/volumes, pulls the latest Postgres image, and starts a new container named `postgres` with the password `quantixpw`.
- **Port**: Exposes PostgreSQL on `5432`.
- **Usage**:
  ```bash
  chmod +x create-dev-db.sh
  ./create-dev-db.sh
  ```

#### `cont-only.sh`
Use this script to deploy the entire Quantix stack (App + DB) within a Podman pod.
- **What it does**: Creates a pod named `quantix-pod`, pulls the Quantix and Postgres images, and starts both containers within the same network namespace.
- **Ports**: Exposes the web application on `3301` and the database on `5432`.
- **Usage**:
  ```bash
  chmod +x cont-only.sh
  ./cont-only.sh
  ```

---
