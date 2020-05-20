FROM golang:1.13-alpine

# Set up apk dependencies
ENV PACKAGES make git libc-dev bash gcc linux-headers eudev-dev curl ca-certificates

ENV BSC_RELAYER_HOME /opt/app

ENV BBC_NETWORK 1
ENV CONFIG_FILE_PATH $BSC_RELAYER_HOME/config/config.json
ENV CONFIG_TYPE "local"
# You need to specify aws s3 config if you want to load config from s3
ENV AWS_REGION ""
ENV AWS_SECRET_KEY ""

# Set working directory for the build
WORKDIR /opt/app

# Add source files
COPY . .

# Install minimum necessary dependencies, remove packages
RUN apk add --no-cache $PACKAGES && \
    make build

# Run as non-root user for security
USER 1000

VOLUME [ $BSC_RELAYER_HOME ]

# Run the app
CMD ./build/bsc-relayer --bbc-network-type $BBC_NETWORK --config-type $CONFIG_TYPE --config-path $CONFIG_FILE_PATH --aws-region $AWS_REGION --aws-secret-key $AWS_SECRET_KEY