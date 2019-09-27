FROM golang:latest as builder
LABEL maintainer="Manu Raj"
WORKDIR /app
COPY main.go .
RUN go build -o goecho

FROM ubuntu:latest
RUN apt-get update -y
RUN mkdir /app
COPY --from=builder /app .
EXPOSE 8080
CMD ["./goecho", "-port", "8080"]

