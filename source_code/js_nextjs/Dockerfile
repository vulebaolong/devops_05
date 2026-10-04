FROM node:24-alpine AS build

WORKDIR /app

# CÓ THAY ĐỔI THƯ VIỆN KHÔNG
# tải thư viên tạo node_module
COPY package-lock.json .
COPY package.json .
RUN npm install

# CÓ THAY ĐỔI CODE HAY KHÔNG
COPY . .

# build source code
RUN npm run build

FROM node:24-alpine

WORKDIR /app

COPY --from=build /app/.next ./.next
COPY --from=build /app/node_modules ./node_modules
COPY --from=build /app/package.json ./package.json
COPY --from=build /app/public ./public

CMD ["npm", "run", "start"]

# docker build -t js_nextjs_img .
# 918.91 MB