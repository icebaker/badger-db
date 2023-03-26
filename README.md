# Badger DB

A minimalist HTTP API for [Badger DB](https://github.com/dgraph-io/badger/).

## Running

Copy the `.env.example` file to `.env` and provide the required data.

```sh
go build
./badger-db
```

## Usage

| Verb      | Path        | Description |
| --------- | ----------- | ----------- |
| `GET`     | `/`         | It provides information about the API and the Database. |
| `GET`     | `/items/:key` | It returns the value for the given `:key`. Status `200` if it exists, `404` if not. |
| `HEAD`    | `/items/:key` | Checks if the `:key` exists without returning its value. Status `204` if it exists, `404` if not. |
| `PUT`     | `/items/:key` | It sets the value for the given `:key`. The value for the key should be sent in the request _body_ as _Plain Text_. Status `201` if created, status `204` if updated. |
| `DELETE`  | `/items/:key` | It deletes the value for the given `:key`. Status `204` if deleted, status `404` if it already doesn't exists. |

## Docker

```yaml
version: '3.7'

services:
  badger-db:
    image: icebaker/badger-db:0.0.1
    environment:
      BADGER_DB_DATA_PATH: /badger-db/data
      BADGER_DB_CONTEXT: my-service-name
      BADGER_DB_HOST: 0.0.0.0
      BADGER_DB_PORT: 9701
    volumes:
      - ./my-project/data/badger-db:/badger-db/data
    ports:
      - 9701:9701
```

## Development

```sh
go run main.go
go fmt
```
