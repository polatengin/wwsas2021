FROM node:15.10.0-alpine AS builder

WORKDIR /src/api-store

COPY ./package*.json ./

RUN npm install

COPY ./api-store/ ./

ENTRYPOINT [ "npm", "run", "start" ]
