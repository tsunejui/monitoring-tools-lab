## Grafana Migration


1. pull image

```
docker pull ghcr.io/dimitri/pgloader:3.6.9
```

2. run command

```
docker run --rm -it ghcr.io/dimitri/pgloader:3.6.9 pgloader --version
```

### Reference
- https://www.digitalocean.com/community/tutorials/how-to-migrate-mysql-database-to-postgres-using-pgloader
- https://github.com/dimitri/pgloader
- https://github.com/wbh1/grafana-sqlite-to-postgres
- https://polyglot.jamie.ly/programming/2019/07/01/grafana-sqlite-to-postgres.html
- https://community.grafana.com/t/migration-from-sqlite3-to-postgres/71022