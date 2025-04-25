FROM golang:1.22.3-alpine AS build

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk add --no-cache make

RUN go build -o server main.go

FROM alpine AS runner

WORKDIR app

RUN apk add --no-cache curl

COPY --from=build /build/server ./server

COPY config/config.yml ./config.yml

CMD ["/app/server", "-config=/app/config.yml"]
