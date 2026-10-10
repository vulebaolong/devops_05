FROM go_gin_be_img AS golang

FROM js_react_vite_img AS reat_vite

FROM nginx:alpine

WORKDIR /app

COPY --from=golang /app .
COPY --from=reat_vite /usr/share/nginx/html /usr/share/nginx/html

CMD ["sh", "-c", "nginx && ./application"]

# docker build -t fe_be_img .
# docker run --name fe_be -d -p 3003:80 -p 8006:8080 fe_be_img