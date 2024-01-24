FROM golang:1.21.3

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /reezanvisramportfolio-server ./cmd

EXPOSE 3000

CMD ["/reezanvisramportfolio-server"]