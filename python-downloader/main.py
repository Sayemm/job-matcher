from infrastructure.kaggle.kaggle_client import KaggleClient
from infrastructure.storage.file_storage import FileStorage
from config.settings import AppConfig
from domain.services.download_service import DownloadService

def main():
        config = AppConfig.load()

        # Infra Layer
        kaggleClient = KaggleClient(config.kaggle)
        fileStorage = FileStorage()

        # Application Service
        download_service = DownloadService(
                kaggle_client= kaggleClient,
                file_storage= fileStorage
        )

        # download dataset
        dataset = download_service.download_dataset(
                dataset_id=config.kaggle.dataset_id
        )

        # copy to destination from the downloaded path to the /data/..
        download_service.copy_to_destination(
                dataset=dataset,
                destination=config.storage.output_path
        )


if __name__ == "__main__":
        main()