import numpy as np
from sklearn.feature_extraction.text import TfidfVectorizer

class TextVectorizer:
    def __init__(self, max_features: int = 5000):
        self.vectorizer = TfidfVectorizer(max_features=max_features)

    def fit_transform(self, texts: list) -> np.ndarray:
        return self.vectorizer.fit_transform(texts).toarray()