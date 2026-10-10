FROM node:24-alpine AS build

WORKDIR /app

COPY package-lock.json .
COPY package.json .

RUN npm install

COPY . .

RUN npm run build

FROM node:24-alpine

WORKDIR /app

COPY --from=build /app/dist ./dist
COPY --from=build /app/node_modules ./node_modules

EXPOSE 3000

CMD ["node", "dist/main.js"]

# docker build -t js_nestjs_img .
# docker run --name js_nestjs -d -p 8004:3000 --env-file=.env js_nestjs_img