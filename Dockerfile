FROM golang:1.21

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY cmd ./cmd
COPY internal ./internal 
RUN go build -v -o /usr/local/bin/app cmd/api/*.go

ENV GIN_MODE=release

EXPOSE 9000

CMD ["app"]