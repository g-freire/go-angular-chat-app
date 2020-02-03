# build stage
FROM golang:latest
WORKDIR /go/src/app
RUN go get github.com/gorilla/websocket
RUN go get github.com/satori/go.uuid
COPY controllers.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# run stage
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/app .
EXPOSE 5000
CMD ["./main"]
