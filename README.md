# ReadME

hey! thanks for checking this out. Sorry that it is still wip, I am trying to
learn how to build Domain Driven Design style backend. I figured the best way to
do that is to build it!

I chose go for this project becuase ive been wanting to learn it for a while
now, and it just seems like a good fit.


## Current Working Routes:
- /register
- /login
- /createPost
- /posts


mainly trying to get very simple MVP done to turn in for school, then will implement a follower sytems and custom feeds


## Setup (not finished)

- make sure to have go installed
- run `go get` or `go mod tidy` to install dependencies
- make sure to have docker installed
- at the root of the project run `docker-compose up -d` to start the postgres database
- cd into the `cmd/api` folder and run `go run main.go` to start the server
- for now, you can only use curl to test the routes. Starting on client app now


## Janky migration setup
im literally just running these sql statements in postico for now:

```
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255)
);

CREATE TABLE posts (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id UUID NOT NULL,
    content TEXT NOT NULL,
    timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


DROP TABLE users;
DROP TABLE posts;
```


