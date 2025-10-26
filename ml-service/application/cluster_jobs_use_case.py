import random
from domain.repositories.job_repository import JobRepository
from domain.services.clustering_service import ClusteringService
from infrastructure.ml.vectorizer import TextVectorizer

class ClusterJobsUseCase:
    def __init__(self, repository: JobRepository, clustering_service: ClusteringService):
        self.repository = repository
        self.clustering_service = clustering_service
        self.vectorizer = TextVectorizer(max_features=1000)

    def execute(self):
        print("Loading jobs from database...")
        all_jobs = self.repository.get_all_jobs()
        print(f"Loaded {len(all_jobs)} jobs")

        # TRAINING PHASE
        training_size = min(10000, len(all_jobs))
        training_jobs = all_jobs[:training_size]
        print(f"Using {training_size} jobs for training K-means")

        # Vectorizing job descriptions
        training_texts = [f"{job.title} {job.description}" for job in training_jobs]
        training_vectors = self.vectorizer.fit_transform(training_texts)
        print(f"Created vectors with shape {training_vectors.shape}")

        # Training
        self.clustering_service.cluster(training_vectors)
        
        # Predict for ALL jobs in batches
        print(f"Assigning clusters to all {len(all_jobs)} jobs...")
        
        batch_size = 5000
        all_assignments = []
        
        for i in range(0, len(all_jobs), batch_size):
            batch_jobs = all_jobs[i:i+batch_size]
            batch_num = (i // batch_size) + 1
            
            print(f"Processing batch {batch_num}: jobs {i} to {i+len(batch_jobs)}")
            
            # Vectorize this batch
            batch_texts = [f"{job.title} {job.description[:500]}" for job in batch_jobs]
            batch_vectors = self.vectorizer.vectorizer.transform(batch_texts).toarray()
            
            # Predict clusters using trained model
            batch_labels = self.clustering_service.predict(batch_vectors)
            
            # Create assignments
            for j, job in enumerate(batch_jobs):
                all_assignments.append((job.id, int(batch_labels[j])))
        
        print(f"Predicted clusters for {len(all_assignments)} jobs")

        # Updating database
        self.repository.update_clusters(all_assignments)
        print(f"Updated {len(all_assignments)} jobs")