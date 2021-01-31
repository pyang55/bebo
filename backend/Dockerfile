FROM golang:1.15-alpine

RUN apk add --no-cache tzdata

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -mod vendor -o bebo .

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

CMD ["/app/bebo"]
