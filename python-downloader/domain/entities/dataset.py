from dataclasses import dataclass
from typing import Optional
import os

@dataclass
class Dataset:
    dataset_id: str
    download_path: str
    csv_filename: str = "postings.csv"
    size_bytes: Optional[int] = None
    record_count: Optional[int] = None

    @property
    def csv_path(self) -> str:
        return os.path.join(self.download_path, self.csv_filename)

