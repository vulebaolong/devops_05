# ---------- builder ----------
FROM node:24-alpine AS builder

WORKDIR /fe

ENV NEXT_TELEMETRY_DISABLED=1

# Copy file dependency trước để tận dụng Docker cache
COPY package.json package-lock.json ./

# Cài đúng dependency theo package-lock.json
RUN npm ci

# Copy toàn bộ source code
COPY . .

# Các biến NEXT_PUBLIC_* được nhúng vào source tại thời điểm build
ARG NEXT_PUBLIC_API_URL

ENV NEXT_PUBLIC_API_URL=${NEXT_PUBLIC_API_URL}

# Kiểm tra các biến build
RUN printenv | grep '^NEXT_PUBLIC_' || true

# Yêu cầu next.config.ts hoặc next.config.js có:
# output: "standalone"
RUN npm run build


# ---------- runner ----------
FROM node:24-alpine AS runner

WORKDIR /fe

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1

# Copy ứng dụng Next.js standalone
COPY --from=builder --chown=nextjs:nodejs /fe/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /fe/public ./public
COPY --from=builder --chown=nextjs:nodejs /fe/.next/static ./.next/static

EXPOSE 3000

ENV PORT=3000
ENV HOSTNAME="0.0.0.0"

CMD ["node", "server.js"]

# docker build -t nodepad-fe .
