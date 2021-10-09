FROM golang:buster as builder

WORKDIR /go/src/app

COPY . /go/src/app

RUN go mod download

COPY . .

RUN go mod tidy

RUN go build -o /go/bin/app 

ENV MONGO_URI=${MONGO_URI}

FROM golang:buster

COPY --from=builder /go/bin/app /

EXPOSE 5000

CMD ["/app"]
