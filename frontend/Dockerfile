FROM node:25.1.0-alpine
RUN npm install -g http-server
WORKDIR /APP
COPY package*.json .
RUN npm install
COPY . .
RUN npm run build
EXPOSE 8081
ENTRYPOINT ["http-server", "dist","-p", "8081"]
