FROM node:18 as builder

ARG VITE_DOMOTICHAT_API_BASE_URL
ENV VITE_DOMOTICHAT_API_BASE_URL=$VITE_DOMOTICHAT_API_BASE_URL

ADD ./ ./

RUN npm install && npm run build

FROM caddy:alpine

ARG PORT
ENV PORT=$PORT

COPY --from=builder ./dist /var/www
COPY ./Caddyfile /etc/caddy/Caddyfile
