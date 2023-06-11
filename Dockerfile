FROM golang:1.20 as build

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd cmd
COPY internal internal
COPY pkg pkg
COPY main.go .

RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /
CMD ["/app"]
