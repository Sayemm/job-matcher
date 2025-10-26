import os
from dotenv import load_dotenv

load_dotenv('../.env')

class Config:
    DB_HOST = os.getenv('DB_HOST', 'localhost')
    DB_PORT = os.getenv('DB_PORT', '5432')
    DB_NAME = os.getenv('DB_NAME', 'jobmatcher')
    DB_USER = os.getenv('DB_USER', 'postgres')
    DB_PASSWORD = os.getenv('DB_PASSWORD', '')
    NUM_CLUSTERS = int(os.getenv('NUM_CLUSTERS', '50'))
    
    @classmethod
    def db_url(cls):
        return f"postgresql://{cls.DB_USER}:{cls.DB_PASSWORD}@{cls.DB_HOST}:{cls.DB_PORT}/{cls.DB_NAME}"