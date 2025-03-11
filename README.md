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

To run the [main.go](http://_vscodecontentref_/1) file and generate the values file:

1. Navigate to the [cmd](http://_vscodecontentref_/2) directory:
    ```sh
    cd cmd
    ```

2. Run the [main.go](http://_vscodecontentref_/3) file:
    ```sh
    go run main.go
    ```

This will create or overwrite a file named [val_back.yaml](http://_vscodecontentref_/4) with the generated values.

### Project Structure

- [main.go](http://_vscodecontentref_/5): The main entry point of the application.
- [values.go](http://_vscodecontentref_/6): Contains the `MakeValues` function to generate values for Kubernetes resources.

### Example

The [main.go](http://_vscodecontentref_/7) file generates a YAML file with the following values:

- [replicas](http://_vscodecontentref_/8): 1
- [repository](http://_vscodecontentref_/9): `docker.io/redis`
- [tag](http://_vscodecontentref_/10): `latest`
- [pullPolicy](http://_vscodecontentref_/11): `IfNotPresent`
- [containerPort](http://_vscodecontentref_/12): 6379
- [serviceType](http://_vscodecontentref_/13): `ClusterIP`
- [port](http://_vscodecontentref_/14): 6379
- [targetPort](http://_vscodecontentref_/15): 6379

The generated file will be named [val_back.yaml](http://_vscodecontentref_/16).