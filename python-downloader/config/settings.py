import os
from dataclasses import dataclass
from typing import Optional

@dataclass
class KaggleConfig:
    username: str
    api_key: str
    dataset_id: str = "arshkon/linkedin-job-postings"

    @classmethod
    def from_env(cls) -> 'KaggleConfig':
        username = os.getenv('KAGGLE_USERNAME')
        api_key = os.getenv('KAGGLE_KEY')

        if not username or not api_key:
            raise ValueError('Kaggle Credential Error')
        
        return cls(username=username, api_key=api_key)
    
@dataclass
class StorageConfig:
    output_dir: str = "/data"
    filename: str = "postings.csv"

    @property
    def output_path(self) -> str:
        return os.path.join(self.output_dir, self.filename)
    
@dataclass
class AppConfig:
    kaggle: KaggleConfig
    storage: StorageConfig

    @classmethod
    def load(cls) -> 'AppConfig':
        return cls(
            kaggle = KaggleConfig.from_env(),
            storage = StorageConfig()
        )