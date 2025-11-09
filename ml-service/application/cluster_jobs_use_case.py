from domain.repositories.job_repository import JobRepository
from domain.services.clustering_service import ClusteringService
from infrastructure.ml.vectorizer import TextVectorizer

class ClusterJobsUseCase:
    def __init__(self, repository: JobRepository, clustering_service: ClusteringService):
        self.repository = repository
        self.clustering_service = clustering_service
        self.vectorizer = TextVectorizer(max_features=3000)  # ← Changed from 1000

    def execute(self):
        print("📖 Loading jobs from database...")
        all_jobs = self.repository.get_all_jobs()
        print(f"✅ Loaded {len(all_jobs)} jobs")

        # TRAINING PHASE
        print("\n🎓 TRAINING PHASE:")
        training_size = min(10000, len(all_jobs))
        training_jobs = all_jobs[:training_size]
        print(f"   Using {training_size} jobs for training K-means")

        # Use more text (1000 chars instead of 500)
        print("   Vectorizing training data...")
        training_texts = [
            self._clean_text(f"{job.title} {job.title} {job.description[:1000]}")  # ← Changed!
            for job in training_jobs
        ]
        training_vectors = self.vectorizer.fit_transform(training_texts)
        print(f"   ✅ Created vectors with shape {training_vectors.shape}")

        # Training
        print("   Training K-means model...")
        self.clustering_service.cluster(training_vectors)
        print("   ✅ K-means model trained")

        # PREDICTION PHASE
        print(f"\n🔮 PREDICTION PHASE:")
        print(f"   Assigning clusters to all {len(all_jobs)} jobs...")
        
        batch_size = 5000
        all_assignments = []
        
        for i in range(0, len(all_jobs), batch_size):
            batch_jobs = all_jobs[i:i+batch_size]
            batch_num = (i // batch_size) + 1
            
            print(f"   Processing batch {batch_num}: jobs {i} to {i+len(batch_jobs)}")
            
            # Use same text processing as training
            batch_texts = [
                self._clean_text(f"{job.title} {job.title} {job.description[:1000]}")  # ← Changed!
                for job in batch_jobs
            ]
            batch_vectors = self.vectorizer.vectorizer.transform(batch_texts).toarray()
            
            batch_labels = self.clustering_service.predict(batch_vectors)
            
            for j, job in enumerate(batch_jobs):
                all_assignments.append((job.id, int(batch_labels[j])))
        
        print(f"   ✅ Predicted clusters for {len(all_assignments)} jobs")

        # Update database
        print(f"\n💾 Updating database...")
        self.repository.update_clusters(all_assignments)
        print(f"✅ Updated {len(all_assignments)} jobs")

        # Save models
        print("\n💾 Saving models for ml-matcher...")
        self._save_models()
        print("✅ Models saved!")
        
        # Statistics
        self._print_statistics(all_assignments)
    
    def _clean_text(self, text: str) -> str:
        """
        Clean text for better matching
        """
        import re
        
        # Convert to lowercase
        text = text.lower()
        
        # Remove URLs
        text = re.sub(r'http\S+|www\S+', '', text)
        
        # Remove email addresses
        text = re.sub(r'\S+@\S+', '', text)
        
        # Remove special characters but keep spaces
        text = re.sub(r'[^a-z0-9\s]', ' ', text)
        
        # Remove extra whitespace
        text = ' '.join(text.split())
        
        return text

    def _save_models(self):
        """Save trained models for ml-matcher"""
        import pickle
        import os
        
        model_dir = '/data/model'
        os.makedirs(model_dir, exist_ok=True)
        
        # Save vectorizer
        vectorizer_path = f'{model_dir}/vectorizer.pkl'
        with open(vectorizer_path, 'wb') as f:
            pickle.dump(self.vectorizer.vectorizer, f)
        print(f"   Saved vectorizer to {vectorizer_path}")
        
        # Save cluster centers
        cluster_centers_path = f'{model_dir}/cluster_centers.pkl'
        with open(cluster_centers_path, 'wb') as f:
            pickle.dump(self.clustering_service.kmeans.cluster_centers_, f)
        print(f"   Saved cluster centers to {cluster_centers_path}")
    
    def _print_statistics(self, assignments):
        """Print cluster statistics"""
        print("\n📊 Cluster Distribution:")
        
        clusters = {}
        for _, cluster_id in assignments:
            clusters[cluster_id] = clusters.get(cluster_id, 0) + 1
        
        for cluster_id in sorted(clusters.keys())[:10]:
            print(f"   Cluster {cluster_id}: {clusters[cluster_id]} jobs")
        
        if len(clusters) > 10:
            print(f"   ... and {len(clusters) - 10} more clusters")
        
        print(f"\n   Total clusters: {len(clusters)}")
        print(f"   Avg jobs per cluster: {len(assignments) // len(clusters)}")
        print(f"   Min cluster size: {min(clusters.values())}")
        print(f"   Max cluster size: {max(clusters.values())}")