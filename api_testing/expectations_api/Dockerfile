FROM node:slim

# Set work dir
WORKDIR /home/node

# Copy files to work dir
COPY . .

# Get dependencies
RUN npm install

# Execute app
ENTRYPOINT [ "npm", "start" ]