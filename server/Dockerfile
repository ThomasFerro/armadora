FROM golang:1.13.5 as back-builder

WORKDIR /src

ADD go.mod .
ADD go.sum .

ADD main.go .
ADD game ./game
ADD infra ./infra

RUN go test ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o /dist/armadora

FROM scratch

WORKDIR /root

COPY --from=back-builder /dist .

EXPOSE 80

CMD [ "./armadora" ]
