# Tractus-X - DEV Helm Repository

Helm charts to deploy Tractus-X products. For differences between the DEV and Stable Helm repository, please refer
to [documentation](../README.md#availability).

# Add Tractus-X Helm Repository

```shell
$ helm repo add tractusx-dev https://eclipse-tractusx.github.io/charts/dev
$ helm search repo tractusx-dev/portal
NAME                    CHART VERSION   APP VERSION     DESCRIPTION                            
tractusx-dev/portal         0.8.0           0.8.0           Helm chart for Catena-X Portal frontend
tractusx-dev/portal-backend 0.8.0           0.8.0           Helm chart for Catena-X Portal backend
$
```

Happy helming!
