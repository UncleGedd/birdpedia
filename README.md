# Birdpedia
A sample project for building a Go web app and deploying it in K8s (Big  Bang)

Inspired by:
- https://www.sohamkamani.com/golang/how-to-build-a-web-application/
- https://www.sohamkamani.com/golang/sql-database/

## Database
`docker run -p 5432:5432 --name bird-postgres -e POSTGRES_HOST_AUTH_METHOD=trust -d postgres`

## todo:
- Install Flux
- Install Istio
- logging?