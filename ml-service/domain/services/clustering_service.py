import numpy as np
from sklearn.cluster import KMeans

class ClusteringService:
    def __init__(self, num_clusters: int):
        self.nun_clusters = num_clusters
        self.kmeans = None

    def cluster(self, vectors: np.ndarray) -> np.ndarray:
        self.kmeans = KMeans(n_clusters=self.nun_clusters, random_state=42)
        labels = self.kmeans.fit_predict(vectors)
        return labels
    
    def predict(self, vertors: np.ndarray) -> np.ndarray:
        if self.kmeans is None:
            raise ValueError("Model not trained yet! Call cluster() first.")
        return self.kmeans.predict(vertors)
