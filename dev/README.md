# Tractus-X - DEV Helm Repository

Helm charts to deploy Tractus-X products. For differences between the DEV and Stable Helm repository, please refer
to [documentation](../README.md#availability).

# Add Tractus-X Helm Repository

```shell
$ helm repo add tractusx-dev https://eclipse-tractusx.github.io/charts/dev
$ helm search repo tractusx-dev
NAME                         	CHART VERSION	APP VERSION           	DESCRIPTION
tractusx-dev/autosetup       	1.0.1        	0.0.2                 	A Helm chart for Kubernetes
tractusx-dev/bpdm            	2.0.0        	2.0.0                 	A Helm chart for deploying the BPDM service
tractusx-dev/bpdm-gate       	2.0.0        	2.0.0                 	A Helm chart for deploying the BPDM gate service
tractusx-dev/daps-server     	1.7.2        	1.7.1                 	DAPS server helm-chart
tractusx-dev/dft-backend     	1.7.0        	1.7.0                 	A Helm chart for DFT application
tractusx-dev/dft-frontend    	1.7.0        	1.7.0                 	A Helm chart for Kubernetes
tractusx-dev/irs             	1.1.1        	1.1.0                 	IRS Helm chart for Kubernetes for Catena-X
tractusx-dev/irs-edc-consumer	1.0.1        	0.1.2                 	IRS Helm chart for the EDC consumer
tractusx-dev/irs-helm        	4.0.0        	2.0.0                 	IRS Helm chart for Kubernetes
tractusx-dev/portal          	0.6.0        	0.6.0                 	Helm chart for Catena-X Portal
tractusx-dev/registry        	0.2.4        	0.2.0-M4-multi-tenancy	Tractus-X Digital Twin Registry Helm Chart
tractusx-dev/sdfactory       	1.1.0        	1.1.0                 	A Helm chart for SDFactory
tractusx-dev/semantic-hub    	0.1.3        	0.1.0-M2              	Helm Chart for the Catena-X Semantic Hub Applic...
$
```

Happy helming!
