FROM golang:alpine as builder
WORKDIR /app
COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/hello .


FROM scratch
COPY --from=builder /go/bin/hello /
CMD ["/hello"]
