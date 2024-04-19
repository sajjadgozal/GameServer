FROM golang:alpine

LABEL maintainer="sg"

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .
COPY .env .

RUN go get -d -v ./...
RUN go install -v ./...

## Build the Go app
# RUN go build -o bin/gameserver
# EXPOSE 3000
# CMD [ "./bin/gameserver" ]

## build with hot reload
RUN go install -mod=mod github.com/githubnemo/CompileDaemon
# RUN go get -v golang.org/x/tools/gopls
ENTRYPOINT CompileDaemon --build="go build -a -installsuffix cgo -o bin/gameserver ." --command=./bin/gameserver