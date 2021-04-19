FROM node:lts-alpine3.12 as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY ./ .
ARG VUE_APP_API_URL="http://localhost:8080"
ENV VUE_APP_API_URL=$VUE_APP_API_URL
RUN npm run build

FROM nginx as production-stage
RUN mkdir /app
COPY --from=build-stage /app/dist /app
COPY nginx.conf /etc/nginx/nginx.conf
