# Badger DB API

A minimalist HTTP API for [Badger DB](https://github.com/dgraph-io/badger/).

## Running

Copy the `.env.example` file to `.env` and provide the required data.

```sh
go build
./badger-api
```

## Usage

| Verb      | Path        | Description |
| --------- | ----------- | ----------- |
| `GET`     | `/`         | It provides information about the API and the Database. |
| `GET`     | `/key` | It returns the value for the given _key_. Status `200` if it exists, `404` if not. |
| `PUT`     | `/key` | It sets the value for the given _key_. The value for the key should be sent in the request _body_ as _Plain Text_. Status `201` if created, status `204` if updated. |
| `DELETE`  | `/key` | It deletes the value for the given _key_. Status `200` if deleted, status `404` if it already doesn't exists. |

## Development

```sh
go run main.go
go fmt
```
