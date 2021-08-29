FROM node:16-alpine3.11 as build

WORKDIR /usr/local/app

##########################################################################
#### Only do npm install if, package.json or package-lock.json change ####
#### Optimizing the dockerfile, using the cache                       ####
#### Some examples don't use package-lock.json to minimize updates    ####
##########################################################################
COPY package*.json /usr/local/app/
RUN npm install

COPY . .

RUN npm run build

# NEXT STAGE (Only have nginx alpine image and the builded code from the stage build )
FROM nginx:1.20.1-alpine

# COPY the builded code from the stage build and past into
COPY --from=build /usr/local/app/dist/angular-template-application-image-multistage /usr/share/nginx/html

# EXPOSE CONTAINER PORT 80
EXPOSE 80
