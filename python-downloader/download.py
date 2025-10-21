import os
import kagglehub
import shutil

username = os.getenv("KAGGLE_USERNAME")
api_key = os.getenv("KAGGLE_KEY")
if not username or not api_key:
    print("Kaggle Credential Error")
    exit(1)
print(f"logged in as: {username}")

try:
    path = kagglehub.dataset_download("arshkon/linkedin-job-postings")
    print(f"Downloaded to:{path}")

    csv_file = os.path.join(path, "postings.csv")
    if not os.path.exists(csv_file):
        print("posting.csv not found!")
        exit(1)

    os.makedirs("/data", exist_ok=True)
    shutil.copy(csv_file, "/data/postings.csv")

    size_bytes = os.path.getsize("/data/postings.csv")
    size_mb = size_bytes (1024*1024)
    print(f"File Ready: {size_mb} MB")

except Exception as e:
    print(f"Error: {r}")
    exit(1)