# Tractus-X - Stable Helm Repository

Helm charts to deploy Tractus-X products. For differences between the DEV and Stable Helm repository, please refer
to [documentation](../README.md#availability).

Attention:  
The Stable Helm Repository hasn't been initialized yet. Command output below is for demonstration only.

# Add Tractus-X Helm Repository

```shell
$ helm repo add tractusx https://eclipse-tractusx.github.io/charts/stable
$ helm search repo tractusx
NAME                         	CHART VERSION	APP VERSION           	DESCRIPTION
tractusx/bpdm            	2.0.0        	2.0.0                 	A Helm chart for deploying the BPDM service
tractusx/bpdm-gate       	2.0.0        	2.0.0                 	A Helm chart for deploying the BPDM gate service
tractusx/daps-server     	1.7.2        	1.7.1                 	DAPS server helm-chart
tractusx/irs-helm        	4.0.0        	2.0.0                 	IRS Helm chart for Kubernetes
tractusx/portal          	0.6.0        	0.6.0                 	Helm chart for Catena-X Portal
$
```

Happy helming!
