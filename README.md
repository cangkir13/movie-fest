# BE Festival - Movie API

## install depedences

```sh
go install
```

## Database

1. create db
2. insert table schema in file -> database/db-schemas
3. insert dummy data user as admin/user -> database/db_movies-user
4. insert dummy data list movies -> database/db_movies-movies

## update environment

```sh
cp .env-exm .env
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
