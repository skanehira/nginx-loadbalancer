FROM golang:1.14.4 as builder
COPY ./main.go /go/src/server/main.go
WORKDIR /go/src/server
ENV CGO_ENABLED=0
RUN go build -o api main.go

FROM gcr.io/distroless/nodejs
COPY --from=builder /go/src/server/api /api
ENTRYPOINT ["/api"]
