FROM ghcr.io/astral-sh/uv:python3.11-alpine
WORKDIR /app
COPY . .
RUN pip install flask
EXPOSE 8002
CMD ["uv","run","app.py"]
