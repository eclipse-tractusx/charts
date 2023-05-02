# Tractus-X - DEV Helm Repository

Helm charts to deploy Tractus-X products. For differences between the DEV and Stable Helm repository, please refer
to [documentation](../README.md#availability).

# Add Tractus-X Helm Repository

```shell
$ helm repo add tractusx-dev https://eclipse-tractusx.github.io/charts/dev
$ helm search repo tractusx-dev
NAME                                          	CHART VERSION	APP VERSION	DESCRIPTION
tractusx-dev/app-dashboard                    	1.0.5        	1.0.4      	A Helm chart for Kubernetes
tractusx-dev/autosetup                        	1.2.4        	1.2.4      	This service will help service provider to set ...
tractusx-dev/bpdm                             	2.0.0        	2.0.0      	A Helm chart for deploying the BPDM service
tractusx-dev/bpdm-gate                        	3.3.0        	3.2.0      	A Helm chart for deploying the BPDM gate service
...
```

Happy helming!
