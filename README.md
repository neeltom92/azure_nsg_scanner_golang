# Introduction

Container Registry API completes the docker command line to allow you to fully manage your namespaces, images and tags.

# Concepts

## Registries

A Registry is a service which stores docker images. Container Registry proposes one registry per region (currently nl-ams and fr-par)

# Getting Started

**Please note** that examples below are using ams1 region (nl-ams), but you may also use par1 (fr-par) if you want by replacing `{ "region": "ams1" }` to `{ "region": "par1" }` in namespace POST request.
Then, change registry host from `rg.nl-ams.scw.cloud` to `rg.fr-par.scw.cloud`.

Create a namespace:

```
curl -X POST --header "x-auth-token: ${SCALEWAY_TOKEN}" https://api.scaleway.com/registry/v1/namespaces -d '{"name": "myfirstnamespace", "organization_id": "${SCALEWAY_ORG}", "description": "mydevnamespace", "region": "ams1"}'
```

Push an image:

```powershell
docker login rg.nl-ams.scw.cloud -u anyuser -p ${SCALEWAY_TOKEN}
docker tag nginx:latest rg.nl-ams.scw.cloud/myfirstnamespace/nginx:latest
docker push rg.nl-ams.scw.cloud/myfirstnamespace/nginx:latest
```
