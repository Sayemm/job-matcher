FROM python:3.11-slim

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY config/ ./config/
COPY infrastructure/ ./infrastructure/
COPY domain/ ./domain/
COPY main.py .

ENV PYTHONPATH=/app

CMD ["python", "main.py"]