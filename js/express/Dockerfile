FROM node:slim

# Get environment argument
ARG env

# Set work dir
WORKDIR /usr/src

# Copy source fil
COPY . .

# Tune environment
RUN ./scripts/tune_env.sh $env

# Run API
ENTRYPOINT [ "scripts/launch.sh" ]
CMD [ "run" ]