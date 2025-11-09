import numpy as np

class MatchingService:    
    def __init__(self):
        self.cluster_centers = None
        self.num_clusters = None
    
    def load_cluster_centers(self, cluster_centers: np.ndarray):
        self.cluster_centers = cluster_centers
        self.num_clusters = len(cluster_centers)
    
    def find_closest_cluster(self, resume_vector: np.ndarray) -> dict:
        if self.cluster_centers is None:
            raise ValueError("Cluster centers not loaded!")
        
        # Calculate distance to each cluster center
        distances = []
        for center in self.cluster_centers:
            distance = self._calculate_similarity(resume_vector, center)
            distances.append(distance)
        
        # Find closest cluster
        closest_cluster_id = int(np.argmax(distances))
        best_score = float(distances[closest_cluster_id])
        
        return {
            'cluster_id': closest_cluster_id,
            'score': best_score
        }
    
    def _calculate_similarity(self, vec1: np.ndarray, vec2: np.ndarray) -> float:
        # Dot product
        dot_product = np.dot(vec1, vec2)
        
        # Magnitudes
        magnitude1 = np.linalg.norm(vec1)
        magnitude2 = np.linalg.norm(vec2)
        
        # Avoid division by zero
        if magnitude1 == 0 or magnitude2 == 0:
            return 0.0
        
        # Cosine similarity
        similarity = dot_product / (magnitude1 * magnitude2)
        
        return similarity