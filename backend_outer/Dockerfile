FROM golang:1.15-alpine as builder

RUN apk --no-cache add git

WORKDIR /go/src/github.com/ChrisLi03/WatchDOG/backend_outer

COPY . .

RUN go get 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /go/src/github.com/ChrisLi03/WatchDOG/backend_outer .

CMD ["./app"]