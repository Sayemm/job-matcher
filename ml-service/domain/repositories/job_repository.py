from abc import ABC, abstractmethod
from typing import List
from domain.entities.job import Job

class JobRepository(ABC):
    @abstractmethod
    def get_all_jobs(self) -> List[Job]:
        pass

    @abstractmethod
    def update_clusters(self, assignments: List[tuple]) -> None:
        pass