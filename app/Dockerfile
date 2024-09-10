FROM golang:alpine as build

WORKDIR /app

COPY . .
RUN  CGO_ENABLED=0 go build -o app main.go

FROM scratch
COPY --from=build /app/app /bin/

ENTRYPOINT ["/bin/app"]