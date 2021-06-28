FROM golang:1.16.5-buster AS build

WORKDIR /src

COPY ./ /src

RUN go build -o product .

FROM gcr.io/distroless/static AS runtime

WORKDIR /app

COPY --from=build /src/product /app/product

ENTRYPOINT [ "/app/product" ]
