FROM golang:1.17-alpine AS builder


WORKDIR /app
RUN apk add --no-cache git ca-certificates

COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o search_user main.go

FROM alpine as final
MAINTAINER "ugo_bh@yahoo.com"
LABEL service="search_user"
LABEL owner="hugobh"
RUN apk --no-cache add ca-certificates tzdata
RUN mkdir /app
RUN chmod 777 /app
RUN mkdir /app/search_user
RUN chmod 777 /app/search_user
WORKDIR /app/search_user
COPY --from=builder /app/configDB.json /app/search_user
COPY --from=builder /app/configHttp.json /app/search_user
COPY --from=builder /app/search_user /app/search_user
RUN chmod 777 /app/search_user/search_user
ENTRYPOINT ["/app/search_user/search_user"]

