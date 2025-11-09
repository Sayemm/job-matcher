import numpy as np
from sklearn.feature_extraction.text import TfidfVectorizer

class TextVectorizer:
    def __init__(self, max_features: int = 3000):  # ← Increased from 1000
        self.vectorizer = TfidfVectorizer(
            max_features=max_features,
            max_df=0.7,              # Ignore words in >70% of documents
            min_df=3,                # Word must appear in at least 3 documents
            stop_words='english',    # Remove common words (the, is, at...)
            ngram_range=(1, 2),      # Use single words AND pairs (e.g., "machine learning")
            sublinear_tf=True        # Scale term frequency logarithmically
        )
    
    def fit_transform(self, texts: list) -> np.ndarray:
        return self.vectorizer.fit_transform(texts).toarray()