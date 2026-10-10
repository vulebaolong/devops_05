FROM python:3.13-slim AS builder

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir --prefix=/install -r requirements.txt

FROM python:3.13-slim AS runtime

WORKDIR /app

# Copy các thư viện đã cài từ stage builder sang
COPY --from=builder /install /usr/local

# Copy mã nguồn cần thiết từ máy vào
COPY app.py gunicorn.conf.py ./

EXPOSE 8080

# Gunicorn tự đọc gunicorn.conf.py trong thư mục /app.
CMD ["gunicorn", "app:app"]

# docker build -t python_flask_be_img .
# docker run --name python_flask_be -d -p 8005:8080 --env-file=.env python_flask_be_img
