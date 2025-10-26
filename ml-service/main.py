from config import Config
from infrastructure.database.postgres_job_repository import PostgresJobRepository
from domain.services.clustering_service import ClusteringService
from application.cluster_jobs_use_case import ClusterJobsUseCase


def main():
    # Create instances (Dependency Injection)
    repository = PostgresJobRepository(Config.db_url())
    clustering_service = ClusteringService(Config.NUM_CLUSTERS)
    use_case = ClusterJobsUseCase(repository, clustering_service)

    use_case.execute()

    repository.close()

if __name__ == "__main__":
    main()