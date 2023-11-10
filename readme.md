# Golang, MongoDB Rest Api With Integration Tests

This is a simple project to demonstrate how to do integration tests
for your API's with golang.

Here is an example request and response from this API


```bash
curl --location 'http://localhost:3030/api/books'
```

Response

```json
{
  "books": [
    {
      "title": "Crime and Punishment",
      "author": {
        "id": "654e618a60034d917aa0ae63",
        "name": "Dostoyevski"
      },
      "likes": 12,
      "comments": [
        {
          "postTitle": "Crime and Punishment",
          "comment": "great read",
          "likes": 3
        }
      ]
    },
    {
      "title": "Notes From The Underground",
      "author": {
        "id": "654e618a60034d917aa0ae63",
        "name": "Dostoyevski"
      },
      "likes": 100,
      "comments": [
        {
          "postTitle": "Notes From The Underground",
          "comment": "good info",
          "likes": 0
        },
        {
          "postTitle": "Notes From The Underground",
          "comment": "I liked this post",
          "likes": 12
        }
      ]
    },
    {
      "title": "Meditations",
      "author": {
        "id": "654e619760034d917aa0ae64",
        "name": "Marcus Aurelius"
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

Another API Request

```bash
curl --location 'http://localhost:3030/api/author/654e619760034d917aa0ae64/books'
```

Response

```json
{
    "books": [
        {
            "title": "Meditations",
            "author": {
                "id": "654e619760034d917aa0ae64",
                "name": "Marcus Aurelius"
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

It will generate a new MongoDB container and populate the collections with some data.

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