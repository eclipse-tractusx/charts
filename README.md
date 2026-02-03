# Tractus-X Helm Charts

To have all Tractus-X sub-product Helm Charts in one place, a central Tractus-X Helm Repository is build. The central
Helm Repository is split into two repositories:

- Dev
- Stable

## Helm Repository URL

Both repositories will be hosted via GitHub Pages within this
repository ([eclipse-tractus-x/charts](https://github.com/eclipse-tractusx/charts)) and will be accessible under URL

- [https://eclipse-tractusx.github.io/charts/dev](https://eclipse-tractusx.github.io/charts/dev) for Dev
- [https://eclipse-tractusx.github.io/charts/stable](https://eclipse-tractusx.github.io/charts/stable) for Stable

## Availability

### Dev Repository

Dev repository will contain all released Helm Charts of any Tractus-X sub-product. Only a certain number of released
Helm charts per Tractus-X sub-product might be kept due to clarity reasons.

The Dev Helm repository will be updated once a day.

### Stable repository

Stable repository will contain all Helm charts versions of Tractus-X sub-products associated with an official Tractus-X
release. Helm charts associated with Tractus-X versions which have reached its end of lifetime, will be removed from the
stable repository.

The Stable Helm repository will be updated when a new Tractus-X release or a patch update is released.

## Usage

### Dev Repository

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
$ helm install tractusx-dev/portal
[...]
```

### Stable Repository

```shell
$ helm repo add tractusx https://eclipse-tractusx.github.io/charts/stable
$ helm search repo tractusx
NAME                         	CHART VERSION	APP VERSION           	DESCRIPTION
tractusx/bpdm            	2.0.0        	2.0.0                 	A Helm chart for deploying the BPDM service
tractusx/bpdm-gate       	2.0.0        	2.0.0                 	A Helm chart for deploying the BPDM gate service
tractusx/daps-server     	1.7.2        	1.7.1                 	DAPS server helm-chart
tractusx/irs-helm        	4.0.0        	2.0.0                 	IRS Helm chart for Kubernetes
tractusx/portal          	0.6.0        	0.6.0                 	Helm chart for Catena-X Portal
$ helm install tractusx/portal
[...]
```

## Local Testing

This repository also needs maintenance. To assess the functionality after e.g., version bumps, you can run it  manually.

[act](https://github.com/nektos/act) ([installation guide](https://nektosact.com/installation/)) is a tool to run jobs within `./workflows` locally. Configuration of events can be stored in `./act`. We adjusted the workflow [helm-repo-index-DEV.yaml](./.github/workflows/helm-repo-index-DEV.yaml) to not interact with github in a writing way (commit / push). This implies to dispatch the workflow with `mode=local`.

```bash
act workflow_dispatch --job repoIndex -e ./.act/test-repoIndex-locally-event.json --bind
```

> [!warning]
>
> Using the `--bind` option act has write access to your local directory. 
> 

The act worfklow will copy the updated / created file `dev/index.yaml` to your local working file. You can check for changes, running `git diff dev/index.yaml`. There are two scenarios:

1. If there have been no new releases, only the created timestamps will be now - the rest remains the same.
2. If there have been new releases, there will be also version bumps or new entries (for new products).
