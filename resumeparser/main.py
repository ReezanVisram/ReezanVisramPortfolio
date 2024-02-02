import functions_framework
import os
from adapters.cloudstorage.cloudstorage_client import CloudStorageClient
from custom_logging import setup_logging
from io import StringIO
from pdfminer.high_level import extract_text_to_fp


GCP_PROJECT_ID = os.environ["GCP_PROJECT_ID"]
CLOUDSTORAGE_BUCKET_NAME = os.environ["CLOUDSTORAGE_BUCKET_NAME"]
CLOUDSTORAGE_EXPECTED_FILENAME = os.environ["CLOUDSTORAGE_EXPECTED_FILENAME"]
DOWNLOADED_FILEPATH = os.environ["DOWNLOADED_FILEPATH"]


@functions_framework.cloud_event
def parse_resume(cloud_event):
    logger = setup_logging()
    logger.info("Received cloud event")
    print(cloud_event)

    if cloud_event["type"] != "google.storage.object.finalize":
        logger.info("An unexpected trigger was received")
        return

    if cloud_event.data["contentType"] != "application/pdf":
        logger.info("A non-pdf file was received, ignoring")
        return

    if (
        cloud_event.data["id"]
        != f"{CLOUDSTORAGE_BUCKET_NAME}/{CLOUDSTORAGE_EXPECTED_FILENAME}"
    ):
        logger.info("A file with an unexpected id was received, ignoring")
        return

    if cloud_event.data["name"] != CLOUDSTORAGE_EXPECTED_FILENAME:
        logger.info("A file with an unexpected name was received, ignoring")
        return

    storage_client = CloudStorageClient(GCP_PROJECT_ID, CLOUDSTORAGE_BUCKET_NAME)

    storage_client.download_file(cloud_event.data["name"], DOWNLOADED_FILEPATH)

    logger.info("Successfully download file to %s", DOWNLOADED_FILEPATH)
    output_string = StringIO()
    with open(DOWNLOADED_FILEPATH) as f:
        extract_text_to_fp(f, output_string)

    print(output_string.getvalue().strip())
