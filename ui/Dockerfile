FROM node:18 as builder

ARG VITE_DOMOTICHAT_API_URL
ENV VITE_DOMOTICHAT_API_URL=$VITE_DOMOTICHAT_API_URL
ADD ./ ./

RUN npm install && npm run build

FROM caddy:alpine

COPY --from=builder ./dist /var/www
COPY ./Caddyfile /etc/caddy/Caddyfile
