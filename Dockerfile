FROM golang:1.20.0

WORKDIR /go/delivery

COPY . .
RUN go build cmd/main.go
ENTRYPOINT [ "/bin/sh", "-c" ]
CMD [ "./main" ]

