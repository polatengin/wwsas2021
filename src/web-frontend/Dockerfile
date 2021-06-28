FROM node:15.11.0-buster as builder

RUN apt-get update

WORKDIR /app

COPY . /app

RUN npm install

RUN npm run build

FROM nginx:1.19.7-alpine as production

USER root

COPY --from=builder /app/dist /usr/share/nginx/html

RUN echo "server { listen 80; server_name localhost; location / { root /usr/share/nginx/html; add_header 'Referrer-Policy' 'origin'; add_header X-Frame-Options 'DENY'; add_header X-Content-Type-Options nosniff; add_header Referrer-Policy 'strict-origin'; try_files \$uri \$uri/ /index.html; index index.html index.htm; } }" > /etc/nginx/conf.d/default.conf

RUN nginx

EXPOSE 80
