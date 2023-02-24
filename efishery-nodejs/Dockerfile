FROM node:16-alpine

WORKDIR /home/node/app
COPY package*.json ./
RUN npm install --silent
COPY . ./

EXPOSE 7314

CMD [ "node", "bin/www"]
