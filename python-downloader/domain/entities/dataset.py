from dataclasses import dataclass, field
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

    @property
    def size_mb(self) -> float:
        if self.size_bytes is None:
            return 0.0
        
        return self.size_bytes / (1024 * 1024)
    
    # BUSINESS RULE: "A dataset is valid only if its CSV file exists"
    def validate(self) -> bool:
        return os.path.exists(self.csv_path)
    
    # This is a BUSINESS RULE: "Dataset is ready when it's validated AND has metadata"
    def is_ready_for_processing(self) -> bool:
        return (self.validate() and
                self.size_bytes is not None and
                self.record_count is not None and
                self.record_count > 0)
