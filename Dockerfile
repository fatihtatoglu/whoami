# Build
FROM golang:1.18.3 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /code

COPY go.mod ./
RUN go mod download && go mod verify

COPY *.go ./

RUN go build -v


# Deploy
FROM alpine

WORKDIR /app

COPY --from=builder /code/whoami /app/whoami

EXPOSE 5000

CMD ["./whoami"]