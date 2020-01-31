FROM node:13.7.0-stretch as front-builder

WORKDIR /src

ADD front/package.json .
ADD front/package-lock.json .

RUN npm i

ADD front .

ENV BUILD_DIR /dist/build
ADD front/static /dist

RUN npm run build

RUN ls /dist
RUN ls /dist

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

COPY --from=front-builder /dist ./static/
COPY --from=back-builder /dist .

EXPOSE 80

CMD [ "./armadora" ]
