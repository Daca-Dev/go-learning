# Gin Framework - Learning

## Setup environment

Install the dependencies with the following command:

```bash
go mod downlaod
```

You need to execute the docker-compose file to start the database.

```bash
docker-compose up -d
```

and copy the `.env.example` file to `.env` and set the environment variables.

```bash
cp .env.example .env
```

finally, populate the data base with the following command:

```bash
go run scripts/populateDB.go
```
