FROM golang:buster as builder

WORKDIR /go/src/app

COPY . /go/src/app

RUN go mod download

RUN go mod tidy

RUN go build -o /go/bin/app 
ENV MONGO_URI=${MONGO_URI}
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/app /


CMD ["/app"]