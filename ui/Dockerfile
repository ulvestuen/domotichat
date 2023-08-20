FROM node:18 as builder

ARG VITE_DOMOTICHAT_API_URL
ENV VITE_DOMOTICHAT_API_URL=$VITE_DOMOTICHAT_API_URL
ADD ./ ./

RUN npm install && npm run build

FROM nginx:alpine-slim

COPY --from=builder dist /usr/share/nginx/html
