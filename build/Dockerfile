FROM golang:1.19

RUN mkdir -p /app/stock

WORKDIR /app

COPY . .

ENV GOPATH="/app"

RUN go build ws-cpt-go.go

EXPOSE 80

# Run
CMD ["/app/ws-cpt-go"]