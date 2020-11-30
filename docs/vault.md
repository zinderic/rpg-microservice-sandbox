# Install Vault

To install Vault in your Kubernetes cluster you can follow the steps from this guide:

https://learn.hashicorp.com/tutorials/vault/kubernetes-raft-deployment-guide

Here's an example install:

```
kubectl create namespace vault
helm repo add hashicorp https://helm.releases.hashicorp.com
helm search repo hashicorp/vault -l
helm install vault hashicorp/vault --namespace vault --version 0.8.0
```

You can access the UI using port-forward:

```
kubectl -n vault port-forward vault-0 8200:8200
```

and opening in browser http://localhost:8200/.

# Initialize Vault

Once intalled you can open its service (port 8200) and initialize it. Once done you should get a root token. Keep it safe. Unseal it with the key and login with the root token:

Example root token and unseal key:

- Initial root token
s.TwSPs8xspmXzylGYhZr5CIAY
- Key 1
laSxueGmZN/hxO0D38tnWxVnlNaDHKEn3xPcpHWjG+0=

# Put secret into Vault

- login first using the root token:

```
VAULT_ADDR=http://localhost:8200/ vault login
```

- put the db-creds secret (example):

```
VAULT_ADDR=http://localhost:8200/ vault write cubbyhole/db-creds POSTGRES_HOST=85c2b41f-e1ac-4530-902b-28f16c9395c7.k8s.civo.com POSTGRES_DB=postgresdb POSTGRES_USER=MAO6MBZP0f POSTGRES_PASSWORD=KwgvXqssTav2xYyptCCWCHvXD9OWkG
```

- read the db-creds:

VAULT_ADDR=http://localhost:8200/ vault read cubbyhole/db-creds

