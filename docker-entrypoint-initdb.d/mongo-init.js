db = db.getSiblingDB('test');

db.createCollection('books');
db.createCollection('authors');
db.createCollection('comments');

db.authors.insertMany([
    {
        "_id": {"$oid": "654e618a60034d917aa0ae63"},
        "name": "Dostoyevski"
    },
    {
        "_id": {"$oid": "654e619760034d917aa0ae64"},
        "name": "Marcus Aurelius"
    }
])

db.books.insertMany([
    {
        "author": {
            "id": "654e618a60034d917aa0ae63",
            "name": "Dostoyevski"
        },
        "likes": 12,
        "title": "Crime and Punishment"
    },
    {
        "author": {
            "id": "654e618a60034d917aa0ae63",
            "name": "Dostoyevski"
        },
        "likes": 100,
        "title": "Notes From The Underground"
    },
    {
        "author": {
            "id": "654e619760034d917aa0ae64",
            "name": "Marcus Aurelius"
        },
        "likes": 200,
        "title": "Meditations"
    }
]);

db.comments.insertMany([
    {
        "comment": "great read",
        "likes": 3,
        "postTitle": "Crime and Punishment"
    },
    {
        "comment": "good info",
        "likes": 0,
        "postTitle": "Notes From The Underground"
    },
    {
        "comment": "I liked this post",
        "likes": 12,
        "postTitle": "Notes From The Underground"
    },
    {
        "comment": "very nice book",
        "likes": 8,
        "postTitle": "Meditations"
    }
])