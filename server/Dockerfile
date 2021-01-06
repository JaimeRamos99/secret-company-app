FROM golang:alpine
#gocache qué es?
RUN GOCACHE=OFF

RUN go env -w GOPRIVATE=github.com/JaimeRamos99/prueba-truora-2

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN apk add git

RUN git config --global url."https://JaimeRamos99:1c9f5eb32b83aaff36b473cdb02f2f14d2744008@github.com".insteadOf "https://github.com"

RUN go build -o main .

ENV PORT 3000
EXPOSE 3000

CMD ["./main"]
