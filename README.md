# BE Festival - Movie API

## install depedences

```sh
go install
```

## update environment

```sh
cp .env-exm
```

## running application

```sh
go run .
```

## Open swagger documentation

```read
http://localhost:8080/api/swagger/index.html
```

### Example register and login

- Url: http://localhost:8080/api/login
- Method: POST
- Payload json

```json
{
  "password": "edp",
  "username": "edp"
}
```
