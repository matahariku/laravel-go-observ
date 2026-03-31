FROM golang:1.20-alpine

WORKDIR /app
COPY . .
RUN go build -o observ-svc

EXPOSE 8080

CMD ["/app/observ-svc"]


