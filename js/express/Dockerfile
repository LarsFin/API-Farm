FROM node:slim

# Set work dir
WORKDIR /usr/src

# Copy source fil
COPY . .

# Install app dependencies
RUN npm install

# Run API
CMD [ "npm", "run", "start" ]