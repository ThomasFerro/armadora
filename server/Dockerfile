FROM golang:1.13.5 as back-builder

WORKDIR /src

ADD go.mod .
ADD go.sum .

ADD main.go .
ADD game ./game
ADD infra ./infra

# Not running tests on /infra/party since they are integration tests that need a mongodb
RUN go test `go list ./... | grep -v /infra/party`

RUN CGO_ENABLED=0 GOOS=linux go build -o /dist/armadora

FROM scratch

WORKDIR /root

COPY --from=back-builder /dist .

EXPOSE 80

CMD [ "./armadora" ]
