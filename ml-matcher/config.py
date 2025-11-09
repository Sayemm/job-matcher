import os

class Config:
    """Configuration for ML Matcher service"""
    
    # Server settings
    HOST = os.getenv('HOST', '0.0.0.0')
    PORT = int(os.getenv('PORT', 5000))
    
    # Model paths
    MODEL_DIR = os.getenv('MODEL_DIR', './model')
    VECTORIZER_PATH = f"{MODEL_DIR}/vectorizer.pkl"
    CLUSTER_CENTERS_PATH = f"{MODEL_DIR}/cluster_centers.pkl"
    
    # ML settings
    NUM_CLUSTERS = int(os.getenv('NUM_CLUSTERS', 50))
    
    @classmethod
    def display(cls):
        """Print configuration for debugging"""
        print("=" * 50)
        print("ML Matcher Configuration")
        print("=" * 50)
        print(f"Host: {cls.HOST}")
        print(f"Port: {cls.PORT}")
        print(f"Model directory: {cls.MODEL_DIR}")
        print(f"Vectorizer: {cls.VECTORIZER_PATH}")
        print(f"Cluster centers: {cls.CLUSTER_CENTERS_PATH}")
        print(f"Number of clusters: {cls.NUM_CLUSTERS}")
        print("=" * 50)