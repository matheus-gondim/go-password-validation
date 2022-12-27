FROM golang:1.19 AS build

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -v -o /usr/local/bin/ ./...

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /usr/local/bin/http /usr/local/bin/http
COPY --from=build /usr/local/bin/graphql /usr/local/bin/graphql

EXPOSE 3000
EXPOSE 3001

USER nonroot:nonroot

ENTRYPOINT [ "/usr/local/bin/graphql", "/usr/local/bin/http" ]