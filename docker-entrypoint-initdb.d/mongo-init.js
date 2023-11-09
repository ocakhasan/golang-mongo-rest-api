db = db.getSiblingDB('test');

db.createCollection('posts');
db.createCollection('comments');

db.posts.insertMany([
    {
        "author": {
            "id": 1,
            "name": "Dostoyevski"
        },
        "likes": 12,
        "title": "Crime and Punishment"
    },
    {
        "author": {
            "id": 1,
            "name": "Dostoyevski"
        },
        "likes": 100,
        "title": "Notes From The Underground"
    },
    {
        "author": {
            "id": 2,
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