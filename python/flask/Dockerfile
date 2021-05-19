FROM python:3.8-slim-buster

# Get environment argument
ARG env

# Set work dir
WORKDIR /usr/src

# Copy files to work dir
COPY . .

# Tune environment
RUN ./scripts/tune_env.sh $env

# Run specified command
ENTRYPOINT [ "scripts/launch.sh" ]
CMD [ "run" ]