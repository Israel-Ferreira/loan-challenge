FROM golang:1.21-alpine AS builder

WORKDIR /usr/app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go


FROM scratch

COPY --from=builder /usr/app/api .

EXPOSE 8000

CMD ["./api"]