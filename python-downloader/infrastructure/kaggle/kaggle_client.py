import kagglehub, os, time
from config.settings import KaggleConfig

class KaggleClient:
    def __init__(self, config: KaggleConfig):
        self._config = config

        if not self._config.username or not self._config.api_key:
            raise ValueError(
                "Kaggle credentials missing. "
                "Check KAGGLE_USERNAME and KAGGLE_KEY environment variables."
            )
    
    def download_dataset(self, dataset_id: str) -> str:
        print("Downloading data....")
        max_retries = 3
        for attempt in range(max_retries):
            try:
                os.environ['KAGGLE_TIMEOUT'] = '300'
                path = kagglehub.dataset_download(dataset_id)
                break
            except Exception as download_error:
                if attempt < max_retries - 1:
                    wait_time = (attempt + 1) * 30
                    time.sleep(wait_time)
                    print(f"Download Failed: {download_error}")
                    print(f"Waiting {wait_time} before retry")
                else:
                    raise
        
        print(f"Downloaded to:{path}")
        return path