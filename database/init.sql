CREATE TABLE IF NOT EXISTS jobs (
    id SERIAL PRIMARY KEY,
    job_id VARCHAR(255) UNIQUE NOT NULL,
    company_name VARCHAR(500),
    title VARCHAR(500) NOT NULL,
    description text,
    location VARCHAR(500),
    remote_allowed BOOLEAN DEFAULT FALSE,
    experiment_level VARCHAR(100),
    min_salary NUMERIC(12, 2),
    max_salary NUMERIC(12, 2),
    cluster_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);