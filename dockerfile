FROM golang:1.26

WORKDIR /app-go-grupal

COPY . .

RUN go build -o webapp

EXPOSE 8080

CMD ["./webapp"]