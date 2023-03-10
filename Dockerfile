FROM golang:1.19-alpine

WORKDIR /go/src/app
COPY go.mod .
RUN go mod download
RUN go mod tidy

COPY . .

EXPOSE 8090
CMD [ "go", "run", "./cmd/server/main.go" ]