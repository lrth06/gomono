FROM node:buster as builder

WORKDIR /usr/src/app

COPY package.json .

RUN yarn install

COPY . .

RUN yarn build


FROM nginx:mainline-alpine

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY --from=builder /usr/src/app/build .

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]

# CMD ["/app"]
