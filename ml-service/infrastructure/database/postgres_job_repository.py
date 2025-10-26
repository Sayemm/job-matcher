import psycopg2
from typing import List
from domain.entities.job import Job
from domain.repositories.job_repository import JobRepository

class PostgresJobRepository(JobRepository):
    def __init__(self, db_url):
        self.conn = psycopg2.connect(db_url)

    def get_all_jobs(self) -> List[Job]:
        cursor = self.conn.cursor()
        cursor.execute("SELECT id, job_id, title, description, cluster_id FROM jobs")
        rows = cursor.fetchall()
        cursor.close()

        jobs = []
        for row in rows:
            if row[2] and row[3]:
                jobs.append(Job(
                    id=row[0],
                    job_id=row[1],
                    title=row[2],
                    description=row[3],
                    cluster_id=row[4]
                ))
        return jobs
    
    def update_clusters(self, assignments: List[tuple]) -> None:
        cursor = self.conn.cursor()
        for job_id, cluster_id in assignments:
            cursor.execute(
                "UPDATE jobs SET cluster_id = %s WHERE id = %s",
                (cluster_id, job_id)
            )
        self.conn.commit()
        cursor.close()

    def close(self):
        self.conn.close()