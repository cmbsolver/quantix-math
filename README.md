# quantix-math
A set of math utilities for solving the Cicada 3301 puzzle. 

## Getting Started

### Database

The project uses **SQLite** as its database engine. The database file (`quantix.db`) is automatically created in the root directory upon first run, based on the configuration in `settings/appsettings.json`.

### Deployment

The project includes a shell script to help with deployment using Podman:

#### `cont-only.sh`
Use this script to deploy the Quantix application within a Podman pod.
- **What it does**: Creates a pod named `quantix-pod`, pulls the Quantix image, and starts the container.
- **Ports**: Exposes the web application on `3301`.
- **Usage**:
  ```bash
  chmod +x cont-only.sh
  ./cont-only.sh
  ```

---
