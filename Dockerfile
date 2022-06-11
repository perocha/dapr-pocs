# Step 1: Modules caching
FROM golang:rc-alpine as modules
#COPY go.mod go.sum /modules/
COPY go.mod /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:rc-alpine as builder
#COPY --from=modules /go/pkg /go/pkg
COPY . /app

WORKDIR /

RUN apk add --no-cache ca-certificates apache2-utils

ADD bin /bin/

CMD ["/bin/sh"]
#ENTRYPOINT [ "bin/main" ]