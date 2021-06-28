FROM node:15.11.0-buster as builder

RUN apt-get update

WORKDIR /app

COPY . /app

RUN npm install
