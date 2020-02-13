FROM node:13.7.0-stretch as front-builder

WORKDIR /src

ADD package.json .
ADD package-lock.json .

RUN npm i

ADD . .

ENV BUILD_DIR /dist/build
ADD static /dist

RUN npm run build

FROM nginx:alpine
COPY --from=front-builder /dist /usr/share/nginx/html
