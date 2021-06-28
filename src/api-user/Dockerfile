FROM mcr.microsoft.com/dotnet/sdk:5.0.102-1-focal AS builder

WORKDIR /src

COPY . .

WORKDIR /src/api-user

RUN dotnet restore

RUN dotnet build -c Release -o /dist
