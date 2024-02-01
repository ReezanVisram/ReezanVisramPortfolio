import json
import functions_framework
import logging.config
import logging.handlers
import pathlib

logger = logging.getLogger("resume_parser")


def setup_logging():
    config_file = pathlib.Path("logging/config.json")
    with open(config_file) as f_in:
        config = json.load(f_in)
    logging.config.dictConfig(config)


@functions_framework.cloud_event
def parse_resume(cloud_event):
    setup_logging()
    print(cloud_event.data)

    logger.info("Received cloud event")
