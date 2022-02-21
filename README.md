# Birdpedia
A sample project for building a Go web app and deploying it in K8s (Big  Bang)

Inspired by:
- https://www.sohamkamani.com/golang/how-to-build-a-web-application/
- https://www.sohamkamani.com/golang/sql-database/

## Database
`docker run -p 5432:5432 --name bird-postgres -e POSTGRES_HOST_AUTH_METHOD=trust -d postgres`

## todo:
- helm init
- PG chart
- run database seeder as an [init container](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/) in the deployment spec,
  - Also see [this](https://www.magalix.com/blog/kubernetes-patterns-the-init-container-pattern)