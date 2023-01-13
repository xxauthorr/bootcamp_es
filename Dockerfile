FROM golang:1.19.3-alpine AS builder

# maintainer info
LABEL maintainer = "Harichandra Kishor <harisown6@gmail.com>"

WORKDIR /app

ADD ./go.mod ./

RUN go mod download

ADD . .

RUN go build -o main 

CMD [ "./main" ]