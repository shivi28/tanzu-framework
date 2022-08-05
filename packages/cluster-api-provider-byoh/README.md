# cluster-api-provider-byoh Package

This package provides consistent deployment and day 2 operations of
"self-managed" and EKS Kubernetes clusters on byoh using
[cluster-api-provider-byoh](https://github.com/kubernetes-sigs/cluster-api-provider-byoh).

## Components

* capa-controller-manager

## Configuration

The following configuration values can be set to customize the installation.

| Value                              | Required/Optional | Description                                                                 |
|------------------------------------|-------------------|-----------------------------------------------------------------------------|
| `capaControllerManager.httpProxy`  | Optional          | Configures the HTTP_PROXY environment variable on capa-controller-manager.  |
| `capaControllerManager.httpsProxy` | Optional          | Configures the HTTPS_PROXY environment variable on capa-controller-manager. |
| `capaControllerManager.noProxy`    | Optional          | Configures the NO_PROXY environment variable on capa-controller-manager.    |

## Usage Example

To learn more about cluster-api-provider-byoh visit
<https://github.com/vmware-tanzu/cluster-api-provider-byoh>.
