VERSION 0.6
FROM golang:latest

WORKDIR /app

requirements:
    COPY go.mod ./
    COPY go.sum ./
    RUN go mod download
    SAVE IMAGE chocode:requirements

build:
    FROM +requirements
    COPY *.go ./
    COPY ./pkg ./pkg
    COPY ./core ./core
    
    RUN go install github.com/swaggo/swag/cmd/swag@latest
    RUN swag init

    RUN go build -o main .
    RUN mv ./main ./run_regos
    # Removing source code after build
    RUN rm -Rdf ./pkg ./core ./main.go ./main_test.go ./go.mod ./go.sum

    EXPOSE 8080

    CMD [ "/app/run_regos" ]

    SAVE IMAGE  chocode:latest