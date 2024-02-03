from io import StringIO
from pdfminer.high_level import extract_text_to_fp


class ResumeParser:
    def __init__(self, resume_filepath):
        self._resume_filepath = resume_filepath
        self._content = self._get_pdf_content()

    def _extract_content_into_keywords(self) -> dict[str, str]:
        d = {}

    def _get_pdf_content(self) -> str:
        content = StringIO()
        with open(self._resume_filepath, "rb") as f_in:
            extract_text_to_fp(content, f_in)

        return content.getvalue().strip()
