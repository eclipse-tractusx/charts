# Tractus-X - Stable Helm Repository

Helm charts to deploy Tractus-X products. For differences between the DEV and Stable Helm repository, please refer
to [documentation](../README.md#availability).

# Add Tractus-X Helm Repository

```shell
$ helm repo add tractusx https://eclipse-tractusx.github.io/charts/stable
$ helm search repo tractusx/portal
NAME                    CHART VERSION   APP VERSION     DESCRIPTION                            
tractusx/portal         0.8.0           0.8.0           Helm chart for Catena-X Portal frontend
tractusx/portal-backend 0.8.0           0.8.0           Helm chart for Catena-X Portal backend
$
```

Happy helming!
