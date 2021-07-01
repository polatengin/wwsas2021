FROM node:15.10.0-alpine AS builder

WORKDIR /src/api-campaign

COPY ./package*.json ./

RUN npm install

COPY . .

ENTRYPOINT [ "npm", "run", "start" ]
