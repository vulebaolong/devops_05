FROM node:24-alpine AS build

WORKDIR /app

COPY package-lock.json .
COPY package.json .
RUN npm install

COPY . .
RUN npm run build

FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html

# docker build -t js_react_vite_img .
# docker run --name js_react_vite -d -p 3001:80 js_react_vite_img