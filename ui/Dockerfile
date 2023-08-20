FROM node:18 as builder

ADD ./ ./

RUN npm install && npm run build

FROM nginx:alpine-slim

COPY --from=builder dist /usr/share/nginx/html
