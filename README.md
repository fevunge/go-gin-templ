
# Let GO

The web API used in the GoLink web application, developed entirely in Go lang

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

[![fevunge](https://img.shields.io/badge/author-fevunge-blue)](https://github.com/fevunge)

[![backend](https://img.shields.io/badge/stack-backend-red)](https://github.com/fevunge)

[![golang](https://img.shields.io/badge/lang-go-rose)](https://go.dev/)

## Authors

- [@fevunge](https://www.github.com/fevunge)

## API Reference

#### Get all items

API [documentation](https://4geojvlclh.apidog.io/), export to OpenAPI, Swagger, Postman... etc!

Contact by [email](mailto:fevunge@outlook.com) to have access

## Test

To text this project, **after clone it**, run

```bash
    cd let-go
    go mod tidy
```

```bash
    go get github.com/labstack/echo/v4
    go get github.com/steebchen/prisma-client-go
    go get -u github.com/golang-jwt/jwt/v5
```

```sh
    # sync the database with your schema
    go run github.com/steebchen/prisma-client-go db push
```

If you just want to re-generate the client, run

```sh
    go run github.com/steebchen/prisma-client-go generate.
```

Finally, run the program at the root of the project

```sh
    go run main.go
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`DATABASE_URL`

## Acknowledgements

- [Go Web Dev with Echo Project](https://echo.labstack.com/docs/quick-start)
- [Go and Prisma](https://goprisma.org/)
- [Jwt in Go Lang](https://golang-jwt.github.io/)
