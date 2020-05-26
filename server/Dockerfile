FROM gcr.io/cloud-builders/go as build

ENV GOPATH /go
ENV GO111MODULE on

WORKDIR ${GOPATH}/src

RUN wget -O /bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.3.1/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

COPY go.mod go.sum ./

RUN go mod download

COPY server.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -installsuffix "static" .

FROM gcr.io/distroless/static:nonroot

COPY --from=build /bin/grpc_health_probe /bin/grpc_health_probe
COPY --from=build /go/bin/server /bin/greeter-server

ENTRYPOINT ["/bin/greeter-server"]
