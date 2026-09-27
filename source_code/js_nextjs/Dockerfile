FROM node:24-alpine

WORKDIR /app

COPY . .

# tạo node_module
RUN npm install
# build source code
RUN npm run build

CMD ["npm", "run", "start"]