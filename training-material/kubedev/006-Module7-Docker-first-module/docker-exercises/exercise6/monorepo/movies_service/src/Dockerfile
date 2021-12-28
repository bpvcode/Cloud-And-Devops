FROM node:17

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 8181

# Start script in package.json
CMD [ "node", "server.js" ]