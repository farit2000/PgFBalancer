FROM golang:1.15.5-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/go_load_balancer

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o ./out/go_load_balancer ./src

FROM alpine:3.7
RUN apk add ca-certificates

COPY --from=build_base /tmp/go_load_balancer/src/PgFBalancer.cfg /PgFBalancer.cfg
COPY --from=build_base /tmp/go_load_balancer/src/web /web
COPY --from=build_base /tmp/go_load_balancer/out/go_load_balancer /app/go_load_balancer

EXPOSE 2345
EXPOSE 5678

CMD ["/app/go_load_balancer"]