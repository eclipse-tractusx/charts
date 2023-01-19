# Decentralized Digital Twin Registry Setup
This document contains instructions how to setup the `dec-dtr` [Helm](https://helm.sh/) chart.
The chart contains deployments for two EDC instances, a DAPS instance and one registry. It relies on the EDC [tractusx-connector](https://github.com/catenax-ng/product-edc/tree/develop/charts/tractusx-connector) Helm chart.

Please be aware, that you need a Kubernetes cluster to deploy the setup using Helm. You can use [minikube](https://minikube.sigs.k8s.io/docs/start/), [K3s](https://k3s.io/), or [MicroK8s](https://microk8s.io/) to deploy a Kubernetes cluster on your local machine.

## Architecture

```
                                   ┌───────────────────┐
                                   │                   │
                                   │    Backend App    │
                                   │                   │
                                   └─────────┬─────────┘
                                             │
                                             │
                                             │
                              ┌──────────────┴────────────┐
                              │                           │
                              │                           │
                              │                           │
                              │                           │
                     ┌────────┴─────────┐                 │
                     │     Adapter      │                 │
                     ├──────────────────┤       ┌─────────┴────────┐
                     │                  │       │                  │
        ┌────────────┤ EDC Controlplane ├───────┤   EDC Dataplane  │
        │            │    (Consumer)    │       │    (Consumer)    │
┌───────┴───────┐    │                  │       │                  │
│               │    └──────────┬───────┘       └─────────┬────────┘
│     DAPS      │               │                         │
│               │               │                         │
└───────┬───────┘    ┌──────────┴───────┐       ┌─────────┴────────┐
        │            │                  │       │                  │
        └────────────┤ EDC Controlplane ├───────┤   EDC Dataplane  │
                     │    (Provider)    │       │    (Provider)    │
                     │                  │       │                  │
                     └──────────────────┘       └─────────┬────────┘
                                                          │
                                                          │
                                                          │
                                  ┌───────────────────┐   │
                                  │                   │   │
                                  │     Registry      ├───┘
                                  │                   │
                                  └───────────────────┘
```

## Setup
To start the deployment run:
```
helm install decentral-registry-setup -n dec-reg . --create-namespace
```

This will install all necessary modules (as described above) on you cluster in the namespace `decentral-registry-setup`.

In the next step we want to create a new asset in the provider EDC controlplane. Make sure that you can access the `plato-edc-controlplane` service from outside you cluster (depending on your K8s provider you need to forward a port). Run the following curl request:
```
curl -0 -v -X POST 'http://<your-k8s-host>/data/assets' \
-H "Expect:" \
-H 'Content-Type: application/json; charset=utf-8' \
-H 'X-Api-Key: <password>' \
--data-binary @- << EOF
{
    "asset": {
        "properties": {
            "asset:prop:id": "registry-id",
            "asset:prop:description": "Digital Twin Registry 5"
        }
    },
    "dataAddress": {
        "properties": {
            "type": "HttpData",
            "baseUrl": "http://cx-registry-setup-registry-svc:8080",
            "proxyPath": true,
            "proxyBody": true,
            "proxyMethod": true
        }
    }
}
EOF
```

In the next step, create a policy for you data offering:
```
curl -0 -v -X POST 'http://<your-k8s-host>/data/policydefinitions' \
-H "Expect:" \
-H 'Content-Type: application/json; charset=utf-8' \
-H 'X-Api-Key: <password>' \
--data-binary @- << EOF
{
    "id": "registry-id-policy",
    "policy": {
        "prohibitions": [],
        "obligations": [],
        "permissions": [
            {
                "edctype": "dataspaceconnector:permission",
                "action": {
                    "type": "USE"
                },
                "constraints": []
            }
        ]
    }
}
EOF
```

Finally, we need to create a contract definition:
```
curl -0 -v -X POST 'http://<your-k8s-host>/data/contractdefinitions' \
-H 'Content-Type: application/json; charset=utf-8' \
-H 'X-Api-Key: <password>' \
--data-binary @- << EOF
{
    "id": "registry-id-contract",
    "criteria": [
        {
            "operandLeft": "asset:prop:id",
            "operator": "=",
            "operandRight": "registry-id"
        }
    ],
    "accessPolicyId": "registry-id-policy",
    "contractPolicyId": "registry-id-policy"
}
EOF
```

With that, we now need to make two calls to receive actual data from the registry. First of all we need to acquire an EDC token that later allows us to fetch data using the EDC dataplane, which functions as a proxy.
The token request to the EDC controlplane looks like the following:
```
curl -0 -v -X GET 'http://<your-k8s-host>/data/adapter/asset/sync/<asset-id>?providerUrl=<provider-url>' \
-H 'Content-Type: application/json; charset=utf-8' \
-H 'X-Api-Key: <password>'
```

The response constist of a JSON payload which contains the token we need for the next step. Copy the token from the `authCode` field and save it for the next request.
To access the registry send the next request to the EDC consumer dataplane:
```
curl -0 -v -X GET 'http://<your-k8s-host>api/public/registry/shell-descriptors>' \
-H 'Content-Type: application/json; charset=utf-8' \
-H 'Authorization: <authCode-from-previous-step>'
```


Execute the following request to fetch the registry:
```
curl -u someuser:somepassword -0 -v -X GET 'http://<your-k8s-host>/api/service/registry-id/registry/shell-descriptors?provider-connector-url=http://plato-edc-controlplane:8282' \
-H 'Content-Type: application/json; charset=utf-8'
```

