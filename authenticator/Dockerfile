FROM golang:alpine as build

WORKDIR /app

COPY . .
RUN go build -o athn main.go

FROM scratch
COPY keys/ keys/

COPY --from=build /app/athn /bin/

ENV PORT=80

CMD ["/bin/athn"]