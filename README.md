[![Build Status](https://travis-ci.org/marccampbell/secretstore.svg?branch=master)](https://travis-ci.org/marccampbell/secretstore) [![Go Report Card](https://goreportcard.com/badge/github.com/marccampbell/secretstore)](https://goreportcard.com/report/github.com/marccampbell/secretstore)

# Quick Start
```bash
# Run a vault dev server
$ docker run -p 8200:8200 -e 'VAULT_DEV_ROOT_TOKEN_ID=my-root-token' -e 'VAULT_DEV_LISTEN_ADDRESS=0.0.0.0:8200' -e 'VAULT_ADDR=http://127.0.0.1:8200'  --name vault vault

# Create a couple of secrets
$ docker exec vault vault write secret/staging/postgres user=admin password=secret host=10.1.1.1 port=5432
$ docker exec vault vault write secret/staging/mandrill key=123456

# Initialize secretstore
$ secretstore init --vault-address http://127.0.0.1:8200 --vault-token=my-root-token
$ secretstore env add --name staging --path secret/staging

# Deploy the secrets to kubernetes
$ secretstore yamls --env staging | kubectl apply -f -
secret "mandrill" created
secret "postgres" created
```

# Why
We needed this to have a single source-of-truth for secrets per-environment for a project we are using Kubernetes to manage. It's easy for various developers to run `kubectl create secret` but that can create problems when different environments go out of sync, or you want to spin up a new environment from the source repos. Using vault to store the secrets and then manually running `secretstore` to deploy the secrets to the Kubernetes clusters has created a workflow that's more managable for us.

# All commands
```bash
secretstore help
secretstore init
secretstore env
secretstore k8s
```

