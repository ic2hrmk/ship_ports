# Build:            docker build -t ship_ports .
# Run gateway:      docker run --rm -p=8081:8081 ship_ports
# Run service:      docker run --rm -p=10001:10001 ship_ports

FROM golang:1.10

WORKDIR /go/src/github.com/ic2hrmk/ship_ports
COPY . .

RUN go build -o port-service entry/entry.go && mv port-service /go/bin/

CMD ["port-service","--env=.env","--kind=port-gtw","--address=:8081"]
# CMD ["port-service","--env=.env","--kind=port-srv","--address=:10001"]
