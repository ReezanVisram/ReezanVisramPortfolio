from google.cloud import storage


class CloudStorageClient:
    def __init__(self, project_id: str, storage_bucket: str):
        client = storage.Client(project=project_id)
        self.bucket = client.bucket(storage_bucket)
        #

    def download_file(self, file_to_download: str, filepath_to_save_to: str) -> None:
        blob = self.bucket.blob(file_to_download)
        blob.download_to_filename(filepath_to_save_to)
