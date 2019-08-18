# マルチステージビルドでイメージサイズを節約する
FROM golang:1.12.7-alpine3.10 AS build

WORKDIR /
COPY . /go/src/github.com/YoshijiFujiwara/u22/backend
RUN apk update \
  && apk add --no-cache git \
  && go get github.com/gorilla/mux \
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql \
  && go get github.com/dgrijalva/jwt-go \
  && go get github.com/pkg/errors \
  && go get github.com/joho/godotenv \
  && go get github.com/gorilla/handlers \
  && go get github.com/justinas/alice \
  && go get github.com/rs/cors \
  && go get golang.org/x/crypto/bcrypt

#RUN cd /go/src/github.com/YoshijiFujiwara/u22/backend && go build -o bin/apiserver ./cmd/api/main.go
RUN cd /go/src/github.com/YoshijiFujiwara/u22/backend/database
#RUN cd /go/src/github.com/YoshijiFujiwara/u22/backend/database && make install && make migrate/init && make migrate/up
#&& cd /go/src/github.com/YoshijiFujiwara/u22/backend && go run ./cmd/api/main.go
#FROM alpine:3.10

#COPY --from=build /go/src/github.com/YoshijiFujiwara/u22/backend/bin/apiserver /usr/local/bin/
#CMD ["apiserver"]
