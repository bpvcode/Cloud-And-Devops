FROM golang:1.16 as build

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.14

WORKDIR /app

COPY --from=build /build/main .

CMD ["./main"]