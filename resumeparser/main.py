import functions_framework
import json
import os
from custom_logging.custom_logging import setup_logging
from google.cloud import storage
from typing import List
from parsing_utils import init_company, is_job_title, is_location, is_start_and_end_date
from pdfminer.high_level import extract_pages
from pdfminer.layout import LAParams, LTTextContainer


GCP_PROJECT_ID = os.environ["GCP_PROJECT_ID"]
CLOUDSTORAGE_BUCKET_NAME = os.environ["CLOUDSTORAGE_BUCKET_NAME"]
CLOUDSTORAGE_FILENAME_TO_FETCH = os.environ["CLOUDSTORAGE_FILENAME_TO_FETCH"]
CLOUDSTORAGE_FILENAME_TO_DOWNLOAD_TO = os.environ[
    "CLOUDSTORAGE_FILENAME_TO_DOWNLOAD_TO"
]


def is_valid_cloud_event(cloud_event: functions_framework.cloud_event) -> bool:
    if cloud_event["type"] != "google.storage.object.finalize":
        return False

    if cloud_event.data["contentType"] != "application/pdf":
        return False

    if (
        cloud_event.data["id"]
        != f"{CLOUDSTORAGE_BUCKET_NAME}/{CLOUDSTORAGE_FILENAME_TO_FETCH}"
    ):
        return False

    return True


def download_resume(filename_to_fetch: str, filename_to_download_to: str):
    client = storage.Client(project=GCP_PROJECT_ID)
    bucket = client.bucket(CLOUDSTORAGE_BUCKET_NAME)

    blob = bucket.blob(filename_to_fetch)
    blob.download_to_filename(filename_to_download_to)


def extract_text_from_resume(params: LAParams, filename: str) -> List[LTTextContainer]:
    pages = extract_pages(filename, laparams=params)
    elements: List[LTTextContainer] = []

    for page in pages:
        for element in page:
            if isinstance(element, LTTextContainer):
                elements.append(element)

    return elements


def parse_experience(elements: List[LTTextContainer]) -> dict:
    is_in_experience_section = False
    is_in_specific_job = False
    is_first_line_in_specific_job = True
    company_name = ""
    experience = {}

    for element in elements:
        line = element.get_text().strip()

        if line == "Experience":
            is_in_experience_section = True
            is_first_line_in_specific_job = True
            continue
        elif "Projects" in line:
            is_in_experience_section = False
            break

        if is_in_experience_section:
            if line.split(" ", 1)[0] == "Tools:":
                is_in_specific_job = False
                is_first_line_in_specific_job = True
            else:
                is_in_specific_job = True

            if is_in_specific_job:
                if is_first_line_in_specific_job:
                    company_name = line
                    init_company(experience, company_name)
                    is_first_line_in_specific_job = False
                    continue
                else:
                    if is_job_title(line):
                        experience[company_name]["job_title"] = line
                    elif is_start_and_end_date(line):
                        experience[company_name]["start_and_end_date"] = line
                    elif is_location(line):
                        experience[company_name]["location"] = line
                    else:
                        if line[0] == "\u2022":
                            experience[company_name]["bullet_points"].append(
                                line.replace("\u2022", "").strip()
                            )
                        else:
                            experience[company_name]["bullet_points"][
                                len(experience[company_name]["bullet_points"]) - 1
                            ] += f" {line.strip()}"
    return experience


@functions_framework.cloud_event
def parse_resume(cloud_event):
    logger = setup_logging()
    logger.info("received cloud event")

    if not is_valid_cloud_event(cloud_event):
        logger.error("invalid cloud event received")
        return

    download_resume(
        CLOUDSTORAGE_FILENAME_TO_FETCH, CLOUDSTORAGE_FILENAME_TO_DOWNLOAD_TO
    )

    logger.info("successfully downloaded resume")

    params = LAParams(line_margin=0.0)

    elements = extract_text_from_resume(params, CLOUDSTORAGE_FILENAME_TO_DOWNLOAD_TO)
    logger.info("successfully extracted elements from resume")

    experience = parse_experience(elements)

    logger.info("successfully parsed experience")
    print(json.dumps(experience, indent=4))
