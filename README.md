# Postgres Go
A basic Golang project using postgres and pgadmin
with docker compose.

The goal of the project and have a stable connection to the database
and have migrations automatically.

## Getting started
To build all containers, run:
```bash
make build
```

This command will raise the containers and the migrations will happen.

To access the pgadmin, open it in your browser:
http://localhost:16543

The default credentials are described in the docker compose as env vars,
are the following:

```
PGADMIN_DEFAULT_EMAIL: "user@admin.com"
PGADMIN_DEFAULT_PASSWORD: "secret"
```
