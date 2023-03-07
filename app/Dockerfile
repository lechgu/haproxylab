FROM golang:alpine as build

WORKDIR /app

COPY . .
RUN go build -o app main.go

FROM scratch
COPY --from=build /app/app /bin/

ENTRYPOINT ["/bin/app"]