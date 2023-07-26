# dating-app with JWT token

## setup dependencies

- [golang-migrate](https://github.com/golang-migrate/migrate) to install migrate CLI

- Postgresql

  ```
  docker compose up -d postgres_db
  ```

- Migration
  ```
  make migrate-up
  ```

## Run API service

- Installl dependencies
  ```
  go mod tidy && go mod vendor
  ```
- Run server
  ```
  make run
  ```
  Run Unit test by package
  ```
  make test package={package_name}
  ```

## Docs

- Postman collection [here](https://elements.getpostman.com/redirect?entityId=21574323-cc1a5a9e-b6ae-4692-94dd-c708ccaa02b6&entityType=collection)
