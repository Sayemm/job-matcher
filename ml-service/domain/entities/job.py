from dataclasses import dataclass

@dataclass
class Job:
    id: int
    job_id: str
    title: str
    description: str
    cluster_id: int = None