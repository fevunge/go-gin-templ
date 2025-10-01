
CC=go
CFLAGS=-ldflags "-s -w"
SRC=main.go
NAME=app

all: $(NAME)


push_prisma:
	go run github.com/steebchen/prisma-client-go db push

gen_prisma:
	go run github.com/steebchen/prisma-client-go generate

install_deps:
	go get github.com/labstack/echo/v4
	go get github.com/steebchen/prisma-client-go
	go get -u github.com/golang-jwt/jwt/v5

setup: install_deps
	$(CC) mod tidy

$(NAME):
	go build $(CFLAGS) -o $(NAME) $(SRC)

run: 

