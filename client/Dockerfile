FROM gcr.io/cloud-builders/go as build

ENV GOPATH /go
ENV GO111MODULE on

WORKDIR ${GOPATH}/src

COPY go.mod go.sum ./

RUN go mod download

COPY client.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -installsuffix "static" .

FROM gcr.io/distroless/static:nonroot

COPY --from=build /go/bin/client /bin/greeter-client

ENTRYPOINT ["/bin/greeter-client"]
