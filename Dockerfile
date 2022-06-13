# Step 1: Modules caching
FROM golang:rc-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:rc-alpine as builder
ARG MODULE_NAME
RUN apk add --no-cache ca-certificates apache2-utils
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app/cmd/$MODULE_NAME
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app -buildvcs=false

# Step 3: Create final image from scratch
FROM scratch
COPY --from=builder /bin/app /bin/app
WORKDIR /bin
#CMD ["/bin/sh"]
ENTRYPOINT [ "/bin/app" ]