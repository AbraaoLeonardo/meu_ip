FROM node:current-alpine3.22
RUN npm install -g http-server
WORKDIR /APP
COPY package*.json .
RUN npm install
COPY . .
RUN npm run build
EXPOSE 8081
ENTRYPOINT ["http-server", "dist"]
