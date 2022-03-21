# Birdpedia
A sample project for building a Go web app and deploying it in K8s (Big  Bang)

Inspired by:
- https://www.sohamkamani.com/golang/how-to-build-a-web-application/
- https://www.sohamkamani.com/golang/sql-database/

## Database
`docker run -p 5432:5432 --name bird-postgres -e POSTGRES_HOST_AUTH_METHOD=trust -d postgres`

## Flux Notes
- Install Flux with `flux install`
- Install Flux secret:
```
flux create secret git github-auth \
   --url=https://github.com/UncleGedd/birdpedia \
   --username=<username> \ 
   --password=<pat_token> \ 
   --namespace=birdpedia
```

## Istio Notes
Installing Istio: https://istio.io/latest/docs/setup/install/helm/

- Installs istio-system namespace
- Installs the istio-base which creates cluster-wide resources used by the Istio control plane
- Installs istiod which is the Istio control plane
- Installs an ingress gateway
- Note that this failed the first time so I deleted the istio-ingress namespace and tried again with helm install istio-ingress istio/gateway -n istio-ingress --wait --create-namespace