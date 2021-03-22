FROM golang:1.16.2 AS build

ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app-svc

from scratch 
COPY --from=build /app/app-svc /app-svc

CMD ["/app-svc"]