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

RUN npm prune --omit=dev

FROM node:24-alpine

WORKDIR /app

# COPY --from=build /app/.next ./.next
# COPY --from=build /app/node_modules ./node_modules
# COPY --from=build /app/package.json ./package.json
# COPY --from=build /app/public ./public

COPY --from=build /app/.next/standalone ./
COPY --from=build /app/public ./public
COPY --from=build /app/.next/static ./.next/static

CMD ["node", "server.js"]

# docker build -t js_nextjs_img .
# docker run --name js_nextjs -d -p 12345:3000 js_nextjs_img
# 918.91 MB
# 727.78 MB
# 615.56 MB
# 221.51 MB