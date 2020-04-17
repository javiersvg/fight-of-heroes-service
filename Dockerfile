FROM golang

ADD . /go/src/github.com/javiersvg/fight-of-heroes-service


RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/handlers
RUN go install github.com/javiersvg/fight-of-heroes-service

ENTRYPOINT /go/bin/fight-of-heroes-service

EXPOSE 8088