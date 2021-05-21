# Determine build environment stage (used for linting & testing)
FROM mcr.microsoft.com/dotnet/core/sdk:2.1 AS build-env

# Get environment argument
ARG env

# Set directory to run following steps from
WORKDIR /app

# Copy source files
COPY . .

# Tune environment
RUN ./scripts/tune_env.sh $env

# Define build-env entrypoint
ENTRYPOINT [ "scripts/launch.sh" ]
CMD [ "run" ]

#-----------------------------------------------------------------------

# Determine production envrionment stage (used only for running api)
FROM mcr.microsoft.com/dotnet/core/aspnet:2.1 AS prod-env

# Set directory to run following steps from
WORKDIR /root/

# Copy binary from build environment stage
COPY --from=build-env /app/ApiFarm/bin/Release/netcoreapp2.1/linux-x64/* .

# Run executable
ENTRYPOINT [ "./ApiFarm" ]
