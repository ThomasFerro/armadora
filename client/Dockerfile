FROM node:13.7.0-stretch as front-builder

WORKDIR /src

ARG API_URL

ADD public /dist
ADD public/static /dist

ADD package.json .
ADD package-lock.json .

RUN npm i

ADD . .

RUN API_URL=$API_URL npm run build

RUN cp -r public/build/ /dist/

FROM nginx:alpine
COPY --from=front-builder /dist /usr/share/nginx/html
