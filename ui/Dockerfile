FROM node:18-alpine as builder

RUN npm install && npm run build

FROM nginx:alpine-slim

ADD --from=builder dist /usr/share/nginx/html
