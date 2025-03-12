FROM golang:1.22 AS build

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app

FROM scratch 

COPY --from=build /bin/app /sbin/app

EXPOSE 8000

USER 1000:1000

ENTRYPOINT ["/sbin/app"]