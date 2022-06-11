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
WORKDIR /app/cmd/hello
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app -buildvcs=false
RUN apk add --no-cache ca-certificates apache2-utils

#CMD ["/bin/sh"]
ENTRYPOINT [ "/bin/app" ]