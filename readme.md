# Golang, MongoDB Rest Api With Integration Tests

This is a simple project to demonstrate how to do integration tests
for your API's with golang.

Here is an example request and response from this API


```bash
curl --location 'http://localhost:3030/api/author/2/posts'
```

Response

```json
{
    "posts": [
        {
            "title": "Meditations",
            "author": {
                "name": "Marcus Aurelius",
                "id": 2
            },
            "likes": 200,
            "comments": [
                {
                    "postTitle": "Meditations",
                    "comment": "very nice book",
                    "likes": 8
                }
            ]
        }
    ]
}
```

To see how the end2end tests are done, please check [integrationtest](internal/controllers/integration_test) folder.

## Run the project

first create the mongo database.

```bash
docker-compose up -d mongo
```

then run

```bash
go run cmd/main.go
```

you will see

```

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.11.3
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:3030
```