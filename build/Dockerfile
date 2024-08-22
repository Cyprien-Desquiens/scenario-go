FROM --platform=linux/amd64 amd64/golang:1.22

RUN mkdir -p /app/stock

WORKDIR /app

COPY . .

ENV GOPATH="/app"

RUN go mod init prom_example
RUN go mod tidy

RUN go build ws-cpt-go.go

EXPOSE 80

# Run
CMD ["/app/ws-cpt-go"]