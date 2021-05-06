FROM node:14

# Set work dir
WORKDIR /usr/src

# App dependencies
COPY package*.json ./

RUN npm install

# Bundle app source
COPY . .

EXPOSE 8080
CMD [ "node", "server.js"]