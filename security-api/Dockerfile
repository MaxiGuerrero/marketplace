FROM golang:alpine as Builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY src src
RUN mkdir dist
RUN go build -v -o dist/main.sh ./src

FROM alpine as Runner
ENV GO_ENV=production
WORKDIR /app

COPY --from=Builder /app/dist .
COPY --from=Builder /app/src/docs/swagger.yaml swagger.yaml

ENTRYPOINT [ "./main.sh" ]