<!--
    Copyright (c) 2023 Robert Bosch Manufacturing Solutions GmbH
    Copyright (c) 2023 Contributors to the Eclipse Foundation

    See the NOTICE file(s) distributed with this work for additional 
    information regarding copyright ownership.
    
    This program and the accompanying materials are made available under the
    terms of the Apache License, Version 2.0 which is available at
    https://www.apache.org/licenses/LICENSE-2.0.
     
    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
    License for the specific language governing permissions and limitations
    under the License.
    
    SPDX-License-Identifier: Apache-2.0
-->

# Decentralized Digital Twin Registry Setup
This document contains instructions how to set up the `dec-dtr` [Helm](https://helm.sh/) chart.
The chart contains deployments for two EDC instances, a DAPS instance, a keycloak and one registry. It relies on the EDC [tractusx-connector](https://github.com/eclipse-tractusx/tractusx-edc) Helm chart.

Please be aware, that you need a Kubernetes cluster to deploy the setup using Helm. You can use [minikube](https://minikube.sigs.k8s.io/docs/start/), [K3s](https://k3s.io/), or [MicroK8s](https://microk8s.io/) to deploy a Kubernetes cluster on your local machine.

## Architecture

```
                  ┌────────────────┐     ┌────────────────┐
                  │                │     │                │
                  │EDC Controlplane│     │  EDC Dataplane │
     ┌────────────┤ (Consumer)     ├─────┤   (Consumer)   │
     │            │                │     │                │
┌────┴──────┐     │                │     │                │
│           │     └───────┬────────┘     └────────┬───────┘
│ DAPS      │             │                       │
│           │             │                       │
│           │             │                       │
└────┬──────┘     ┌───────┴────────┐     ┌────────┴───────┐
     │            │                │     │                │
     │            │EDC Controlplane│     │  EDC Datplane  │
     └────────────┤ (Provider)     ├─────┤   (Provider    ├──────┐
                  │                │     │                │      │
                  │                │     │                │      │
                  └────────────────┘     └─────────┬──────┘      │
                                                   │             │
                                                   │             │
                          ┌──────────────┐         │    ┌────────┴────┐
                          │              │         │    │             │
                          │              ├─────────┘    │             │
                          │     DDTR     │              │  Keycloak   │
                          │              ├──────────────┤             │
                          │              │              │             │
                          └──────────────┘              └─────────────┘                                 
```

## Setup
1. To start the deployment run:
```
helm install decentral-registry-setup -n dec-reg . --create-namespace
```
This will install all necessary modules (as described above) on you cluster in the namespace `decentral-registry-setup`.

2. Add the public key which is located in `/config/defaultKeys/daps.crt` to DAPS. Therefore following the steps:
* Open `config/registry_test.sh` and update the DAPS parameter (DAPS_URL, DAPS_ADMIN, DAPS_ADMIN_SECRET)
* to upload public key for consumer run:
```
./config/register_test.sh consumer http://consumer-edc-controlplane/consumer "idsc:BASE_SECURITY_PROFILE" "config/defaultKeys/daps.crt"
```
* to upload public key for provider run:
```
./config/register_test.sh provider http://provider-edc-controlplane/provider "idsc:BASE_SECURITY_PROFILE" "config/defaultKeys/daps.crt"
```

3. After all components running, download the postman collection from [EDC Postman Collection](https://github.com/eclipse-tractusx/tractusx-edc/tree/main/docs/development/postman) and start to create edc-assets, policies and contracts.
