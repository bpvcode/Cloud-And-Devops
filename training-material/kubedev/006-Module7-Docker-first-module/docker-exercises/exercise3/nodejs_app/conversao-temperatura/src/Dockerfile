FROM node:17

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
COPY package*.json ./

RUN npm install

# Bundle app source
COPY . .

# Just instructs that we will map this port
EXPOSE 8080

# Start script in package.json
CMD [ "node", "server.js" ]