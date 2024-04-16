FROM golang:1.22.0 

WORKDIR /todo-app

COPY . .

RUN go mod download 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./todo-app ./cmd/main.go

FROM ubuntu:latest

COPY --from=0 ./todo-app ./

ARG TODO_PORT=7540
ARG TODO_DBFILE
ARG TODO_PASSWORD
ARG TODO_SECRET

ENV TODO_PORT=${TODO_PORT}
ENV TODO_DBFILE=${TODO_DBFILE}
ENV TODO_PASSWORD=${TODO_PASSWORD}
ENV TODO_SECRET=${TODO_SECRET}

EXPOSE ${TODO_PORT} 

CMD [ "./todo-app" ]
