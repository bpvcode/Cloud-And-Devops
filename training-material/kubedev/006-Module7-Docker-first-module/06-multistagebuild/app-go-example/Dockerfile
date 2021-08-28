FROM golang:1.16

WORKDIR /app/go

COPY . .

#Comando de build da aplicação go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

CMD ["go","run","main.go"]