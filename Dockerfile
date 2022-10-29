FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/ToffaKrtek/goCrud/
WORKDIR /go/src/github.com/ToffaKrtek/goCrud
RUN go mod download
COPY . /go/src/github.com/ToffaKrtek/goCrud
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsufix cgo -o build/goCrud github/ToffaKrtek/goCrud


FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/ToffaKrtek/goCrud /usr/bin/gocrud
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/gocrud"]
