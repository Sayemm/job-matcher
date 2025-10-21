FROM python:3.11-slim

WORKDIR /app

RUN pip install --no-cache-dir kagglehub

COPY download.py .

CMD ["python", "download.py"]