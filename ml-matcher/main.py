import pickle
import numpy as np

from config import Config
from domain.services.matching_service import MatchingService
from infrastructure.ml.vectorizer import ResumeVectorizer
from application.match_resume_use_case import MatchResumeUseCase
from infrastructure.http.app import MatcherAPI

def main():
    print("🚀 ML Matcher Service Starting...")
    
    # Display configuration
    Config.display()
    
    # Step 1: Load vectorizer (Infrastructure)
    print("\n📖 Loading vectorizer...")
    vectorizer = ResumeVectorizer(Config.VECTORIZER_PATH)
    
    # Step 2: Load cluster centers
    print("📖 Loading cluster centers...")
    with open(Config.CLUSTER_CENTERS_PATH, 'rb') as f:
        cluster_centers = pickle.load(f)
    print(f"✅ Loaded {len(cluster_centers)} cluster centers")
    
    # Step 3: Create domain service
    print("🔧 Initializing matching service...")
    matching_service = MatchingService()
    matching_service.load_cluster_centers(cluster_centers)
    print("✅ Matching service ready")
    
    # Step 4: Create use case (Application)
    print("🔧 Creating use case...")
    use_case = MatchResumeUseCase(matching_service, vectorizer)
    print("✅ Use case ready")
    
    # Step 5: Create HTTP API (Infrastructure)
    print("🌐 Starting HTTP server...")
    api = MatcherAPI(use_case)
    
    print("\n" + "=" * 50)
    print("✅ ML Matcher Service Ready!")
    print("=" * 50)
    print(f"📍 Endpoints:")
    print(f"   GET  http://localhost:{Config.PORT}/health")
    print(f"   POST http://localhost:{Config.PORT}/match")
    print("=" * 50 + "\n")
    
    # Step 6: Start server
    api.run(host=Config.HOST, port=Config.PORT)

if __name__ == "__main__":
    main()