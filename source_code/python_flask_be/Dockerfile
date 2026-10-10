FROM python:3.13-slim AS build

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir --prefix=/install -r requirements.txt

FROM python:3.13-slim

WORKDIR /app

COPY --from=build /install /usr/local
COPY app.py .
COPY gunicorn.conf.py .

EXPOSE 8080

CMD ["gunicorn", "app:app"]

# docker build -t python_flask_be_img .
# docker run --name python_flask_be -d -p 8005:8080 --env-file=.env python_flask_be_img