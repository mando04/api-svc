FROM golang:1.16.2 AS build

ARG BUILD_VERSION
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.BuildVersion=$BUILD_VERSION" -o app-svc

FROM scratch 
COPY --from=build /app/app-svc /app-svc
COPY --from=build /app/ssl /ssl

CMD ["/app-svc"]