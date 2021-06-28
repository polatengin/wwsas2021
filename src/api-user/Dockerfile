FROM mcr.microsoft.com/dotnet/sdk:5.0.102-1-focal AS builder

WORKDIR /src

COPY . .

WORKDIR /src/api-user

RUN dotnet restore

RUN dotnet build -c Release -o /dist

FROM mcr.microsoft.com/dotnet/aspnet:5.0.2-alpine3.12

RUN apk add icu-libs

ENV DOTNET_SYSTEM_GLOBALIZATION_INVARIANT=false

WORKDIR /app

COPY --from=builder /dist .

ENV ASPNETCORE_URLS=http://+:5000

EXPOSE 5000

ENTRYPOINT [ "dotnet", "api-user.dll" ]
