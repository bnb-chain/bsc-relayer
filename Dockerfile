FROM golang:1.19-alpine as build

# Set up apk dependencies
ENV PACKAGES make git libc-dev bash gcc linux-headers eudev-dev curl ca-certificates

# Set working directory for the build
WORKDIR /opt/app

# Add source files
COPY . .

# Install minimum necessary dependencies, remove packages
RUN apk add --no-cache $PACKAGES && \
    make build

FROM alpine:3.17
WORKDIR /opt/app
RUN mkdir /opt/app/config

COPY config/config.json /opt/app/config/config.json
COPY --from=build /opt/app/build/bsc-relayer /opt/app/bsc-relayer

ENV BSC_RELAYER_HOME /opt/app

ENV BBC_NETWORK 1
ENV CONFIG_FILE_PATH $BSC_RELAYER_HOME/config/config.json
ENV CONFIG_TYPE "local"
# You need to specify aws s3 config if you want to load config from s3
ENV AWS_REGION ""
ENV AWS_SECRET_KEY ""

# Run as non-root user for security
USER 1000

VOLUME [ $BSC_RELAYER_HOME ]

# Run the app
CMD ./bsc-relayer --bbc-network-type $BBC_NETWORK --config-type $CONFIG_TYPE --config-path $CONFIG_FILE_PATH --aws-region $AWS_REGION --aws-secret-key $AWS_SECRET_KEY