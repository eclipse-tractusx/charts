# Tractus-X Helm Charts

To have all Tractus-X sub-product Helm Charts in one place, a central Tractus-X Helm Repository is build. The central
Helm Repository is split into two repositories:

- Dev
- Stable

## Helm Repository URL

Both repositories will be hosted via GitHub Pages within this
repository ([eclipse-tractus-x/charts](https://github.com/eclipse-tractusx/charts)) and will be accessible under URL

- https://eclipse-tractusx.github.io/charts/dev for Dev
- https://eclipse-tractusx.github.io/charts/stable for Stable

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
$ helm search repo tractusx-dev/portal
NAME                    CHART VERSION   APP VERSION     DESCRIPTION                            
tractusx-dev/portal         0.8.0           0.8.0           Helm chart for Catena-X Portal frontend
tractusx-dev/portal-backend 0.8.0           0.8.0           Helm chart for Catena-X Portal backend
$ helm install tractusx-dev/portal
[...]
```

### Stable Repository

```shell
$ helm repo add tractusx https://eclipse-tractusx.github.io/charts/dev
$ helm search repo tractusx/portal
NAME                    CHART VERSION   APP VERSION     DESCRIPTION                            
tractusx/portal         0.8.0           0.8.0           Helm chart for Catena-X Portal frontend
tractusx/portal-backend 0.8.0           0.8.0           Helm chart for Catena-X Portal backend
$ helm install tractusx/portal
[...]
```
