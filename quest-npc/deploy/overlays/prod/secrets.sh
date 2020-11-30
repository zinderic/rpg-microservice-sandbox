#!/bin/bash

token=$(VAULT_ADDR=http://localhost:8200/ vault print token)

if [[ ! $token ]]; then
    echo "You need to login using 'VAULT_ADDR=http://localhost:8200/ vault login'"
    exit 2
fi

# VAULT_ADDR=http://localhost:8200/ vault read cubbyhole/db-creds -format=json | jq .data > db-creds-file

echo "POSTGRES_DB=" `VAULT_ADDR=http://localhost:8200/ vault read cubbyhole/db-creds -format=json | jq -r .data.POSTGRES_DB` > db-creds-file
echo "POSTGRES_HOST=" `VAULT_ADDR=http://localhost:8200/ vault read cubbyhole/db-creds -format=json | jq -r .data.POSTGRES_HOST` >> db-creds-file
echo "POSTGRES_PASSWORD=" `VAULT_ADDR=http://localhost:8200/ vault read cubbyhole/db-creds -format=json | jq -r .data.POSTGRES_PASSWORD` >> db-creds-file
echo "POSTGRES_USER=" `VAULT_ADDR=http://localhost:8200/ vault read cubbyhole/db-creds -format=json | jq -r .data.POSTGRES_USER` >> db-creds-file
