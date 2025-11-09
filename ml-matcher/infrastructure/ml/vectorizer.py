import pickle
import numpy as np
import re

class ResumeVectorizer:
    """
    Infrastructure component - uses sklearn for vectorization
    """
    
    def __init__(self, model_path: str = None):
        self.vectorizer = None
        if model_path:
            self.load(model_path)
    
    def load(self, model_path: str):
        """Load trained vectorizer from disk"""
        with open(model_path, 'rb') as f:
            self.vectorizer = pickle.load(f)
        print(f"✅ Vectorizer loaded from {model_path}")
    
    def transform(self, text: str) -> np.ndarray:
        """
        Convert text to vector
        
        CRITICAL: Must use same preprocessing as ml-service!
        """
        if self.vectorizer is None:
            raise ValueError("Vectorizer not loaded!")
        
        # Clean text (same as ml-service)
        cleaned_text = self._clean_text(text)
        
        # Transform
        vector_sparse = self.vectorizer.transform([cleaned_text])
        vector = vector_sparse.toarray()[0]
        
        return vector
    
    def _clean_text(self, text: str) -> str:
        """
        Clean text - MUST MATCH ml-service cleaning!
        """
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