FROM golang:alpine

RUN GOCACHE=OFF

RUN go env -w GOPRIVATE=github.com/JaimeRamos99/prueba-truora-2

WORKDIR /app

COPY . .

RUN apk add git

RUN git config --global url."https://JaimeRamos99:1c9f5eb32b83aaff36b473cdb02f2f14d2744008@github.com".insteadOf "https://github.com"

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
