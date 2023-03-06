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
    COPY ./libs ./libs
    COPY ./handlers ./handlers
    COPY ./repository ./repository
    
    # RUN go install github.com/swaggo/swag/cmd/swag@latest
    # RUN swag init

    RUN go build -o main .
    # Removing source code after build
    RUN rm -Rdf ./libs ./handlers ./repository # Remove folders
    RUN rm -Rdf ./main.go  ./go.mod ./go.sum # Remove source files

    EXPOSE 8080

    CMD [ "/app/main" ]

    SAVE IMAGE  chocode:latest