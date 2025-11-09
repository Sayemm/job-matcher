from domain.entities.resume import Resume
from domain.services.matching_service import MatchingService
from infrastructure.ml.vectorizer import ResumeVectorizer

class MatchResumeUseCase:
    def __init__(self, matching_service: MatchingService, vectorizer: ResumeVectorizer):
        self.matching_service = matching_service
        self.vectorizer = vectorizer
    
    def execute(self, resume_text: str) -> dict:
        # Step 1: Create domain entity
        resume = Resume(text=resume_text)
        
        # Step 2: Validate (business rule)
        if not resume.is_valid():
            raise ValueError("Resume must be at least 50 characters")
        
        # Step 3: Get clean text
        clean_text = resume.get_text_for_matching()
        
        # Step 4: Vectorize (infrastructure)
        resume_vector = self.vectorizer.transform(clean_text)
        
        # Step 5: Find matching cluster (domain service)
        result = self.matching_service.find_closest_cluster(resume_vector)
        
        return result