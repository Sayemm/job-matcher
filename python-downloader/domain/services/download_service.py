from typing import Optional
from domain.entities.dataset import Dataset

class DownloadService:
    def __init__(
        self,
        kaggle_client,
        file_storage
    ):
        self._kaggle = kaggle_client
        self._storage = file_storage

    def download_dataset(self, dataset_id: str) -> Dataset:
        print(f"Starting to download from dataset: {dataset_id}")

        download_path = self._kaggle.download_dataset(dataset_id)

        dataset = Dataset(
            dataset_id=dataset_id,
            download_path=download_path
        )
        print(f"Dataset downloaded to: {dataset.download_path}")

        dataset.size_bytes = self._storage.get_file_size(dataset.csv_path)
        dataset.record_count = self._storage.count_csv_records(dataset.csv_path)
        
        return dataset
    
    def copy_to_destination(
            self,
            dataset: Dataset,
            destination: str
    ) -> None:
        print(f"Copying dataset to: {destination}")

        self._storage.copy_file(
            source=dataset.csv_path,
            destination=destination
        )
        copied_size = self._storage.get_file_size(destination)

        print(f"Copied {copied_size} MB!")