# ---- Build Phase ----
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /root/space_trouble cmd/main.go

# ---- Deploy Phase ----
FROM ubuntu:22.04
# FROM alpine:3.20

# RUN apk --no-cache add ca-certificates
RUN apt-get update
RUN apt-get install -y
RUN apt install certbot -y
# RUN snap install core
# RUN snap install --classic certbot
# RUN ln -s /snap/bin/certbot /usr/bin/certbot

WORKDIR /root/

COPY --from=builder /root/space_trouble /root/

RUN chmod a+x /root/space_trouble

EXPOSE 8080

CMD ["/root/space_trouble"]
