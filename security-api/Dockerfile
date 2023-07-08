FROM golang:1.20 as Builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY src src
RUN mkdir dist
RUN go build -v -o dist/main.sh ./src

FROM golang:1.20 as Runner
ENV GO_ENV=production
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY --from=Builder /app/dist .
RUN ls .

ENTRYPOINT [ "./main.sh" ]