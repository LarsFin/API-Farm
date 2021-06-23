# Determine build environment stage (used for linting & testing)
FROM golang:1.16 AS build-env

# Get environment argument
ARG env

# Set directory to run following steps from
WORKDIR /go/src/apifarm

# Copy source files
COPY . .

# Tune environment
RUN ./scripts/tune_env.sh $env

# Build executable
RUN CGO_ENABLED=0 GOOS=linux go build

# Define build-env entrypoint
ENTRYPOINT [ "scripts/launch.sh" ]
CMD [ "run" ]

#-----------------------------------------------------------------------

# Determine production environment stage
FROM alpine:latest as prod-env

# Set directory to run following steps from
WORKDIR /root/

# Copy binary from build environment stage
COPY --from=build-env /go/src/apifarm/apifarm .

# Run executable
ENTRYPOINT [ "./apifarm" ]
