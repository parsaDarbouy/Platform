# Platform

This platform is designed to generate values for Kubernetes projects. It helps in creating configuration files for Kubernetes deployments, services, and other resources.

## Getting Started

### Prerequisites

- Go 1.20 or later

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/parsaDarbouy/Platform.git
    cd Platform
    ```

### Usage

To run the `main.go` file and generate the values file:

1. Navigate to the `cmd` directory:
    ```sh
    cd cmd
    ```

2. Run the `main.go` file:
    ```sh
    go run main.go
    ```

This will create or overwrite a file named `val_back.yaml` with the generated values.

### Project Structure

- `main.go`: The main entry point of the application.
- `values.go`: Contains the `MakeValues` function to generate values for Kubernetes resources.

### Example

The `main.go` file generates a YAML file with the following values:

- `replicas`: 1
- `repository`: `docker.io/redis`
- `tag`: `latest`
- `pullPolicy`: `IfNotPresent`
- `containerPort`: 6379
- `serviceType`: `ClusterIP`
- `port`: 6379
- `targetPort`: 6379

The generated file will be named `val_back.yaml`.