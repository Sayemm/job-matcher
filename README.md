# Job Matcher - AI-Powered Resume to Job Matching

An intelligent job matching system that analyzes your resume and recommends relevant jobs from job postings using machine learning clustering.

## What It Does

Upload your resume (PDF or TXT) and instantly get personalized job recommendations. The system uses machine learning to understand your skills and experience, then matches you with jobs in similar categories.

## How It Works

- **Data Pipeline**: Downloads and processes 123K+ job postings from LinkedIn dataset
- **ML Clustering**: Groups similar jobs together using K-means clustering (50 clusters based on job descriptions)
- **Smart Matching**: When you upload a resume, it finds which cluster best matches your skills and shows jobs from that cluster
- **Real-Time Results**: Get match scores using cosine similarity and recommend jobs

## Architecture Highlights

- **Modular Monolith**: Clean separation into independent modules (data loading, ML training, matching, API, frontend) that work together
- **Domain-Driven Design**: Business logic separated from technical implementation using DDD principles (domain, application, infrastructure layers)
- **Batch + Real-Time**: ML training runs once on all jobs (batch), then real-time matching service responds to resume uploads instantly
- **Polyglot Stack**: Right tool for the job - Go for fast API, Python for ML, React for modern UI

## Prerequisites

- Docker
- Docker Compose
- Kaggle Account

## Quick Start

### 1. Clone Repository
```bash
git clone https://github.com/Sayemm/job-matcher.git
cd job-matcher
```

### 2. Configure

Create `.env` file:
```env
DB_HOST=postgres
DB_PORT=5432
DB_NAME=jobmatcher
DB_USER=postgres
DB_PASSWORD=yourpassword123
DB_ENABLE_SSL_MODE=disable

KAGGLE_USERNAME=your_kaggle_username
KAGGLE_KEY=your_kaggle_key

NUM_CLUSTERS=50
CSV_PATH=/data/postings.csv
```

### 3. Build All Services
```bash
docker-compose build
```

### 4. Run Setup (First Time Only)
```bash
# Start database
docker-compose up -d postgres

# Download data
docker-compose up downloader

# Load data
docker-compose up loader

# Train ML model
docker-compose up ml-service
```

### 5. Start Application
```bash
# Start services
docker-compose up -d ml-matcher api frontend
```

### 6. Access

Open browser: http://localhost:3000

## Services

- **Frontend**: http://localhost:3000
- **API**: http://localhost:8080
- **ML Matcher**: http://localhost:5000

## Stop Application
```bash
docker-compose down
```

## Tech Stack

- **Frontend**: React + Tailwind CSS
- **Backend**: Go
- **ML**: Python + scikit-learn
- **Database**: PostgreSQL

