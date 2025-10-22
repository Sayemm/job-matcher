import os
import shutil
from typing import Optional

class FileStorage:
    def get_file_size(self, path: str) -> Optional[int]:
        return os.path.getsize(path) / (1024*1024)
    
    def count_csv_records(self, path: str) -> int:
        with open(path, 'r', encoding='utf-8') as f:
            total_lines = sum(1 for _ in f)
            record_count = max(0, total_lines-1)
            return record_count
        
    def copy_file(self, source: str, destination: str) -> None:
        shutil.copy(source, destination)