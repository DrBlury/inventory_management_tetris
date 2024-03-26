FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

WORKDIR /build

# Copy and download dependency using go mod
ADD ./src/go.* /build/
RUN go mod download

# Copy sources to build container
ADD ./src /build/

# Build the app
WORKDIR /build
RUN ls -al 
RUN go build -a -tags musl -ldflags="-X 'main.version=v0.0.1' -X 'main.buildUser=$(id -u -n)' -X 'main.buildDate=$(date)'" -o /build/linuxcode/inventory_manager
######################################
FROM alpine:3
LABEL AUTHOR="Julian Bensch (Linuxcode)"

# install curl for healthcheck
RUN apk --no-cache add curl

#RUN apk --no-cache add curl
USER nobody
COPY --from=builder --chown=nobody /build/linuxcode/inventory_manager /custom/linuxcode/inventory_manager

ENTRYPOINT [ "/custom/linuxcode/inventory_manager" ]
