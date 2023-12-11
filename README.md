# ReadME

hey! thanks for checking this out. Sorry that it is still wip, I am trying to
learn how to build Domain Driven Design style backend. I figured the best way to
do that is to build it!

I chose go for this project becuase ive been wanting to learn it for a while
now, and it just seems like a good fit.

## setup
- install go
```
sudo apt update
sudo apt install golang
```

### install dependencies
- at root of project, run `go mod tidy`

### run the backend service
- `go run cmd/main.go`

- after a moment, you should see a message that it has connected to the sqlite database


## Current Working Routes:
- /register
- /login
- /createPost
- /posts


mainly trying to get very simple MVP done to turn in for school, then will implement a follower system and custom feeds





