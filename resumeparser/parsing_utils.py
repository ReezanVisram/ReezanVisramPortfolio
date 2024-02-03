import re


def init_company(experience: dict, company_name: str):
    experience[company_name] = {}
    experience[company_name]["bullet_points"] = []


def is_job_title(line: str) -> bool:
    return "engineer" in line.lower() or "developer" in line.lower()


def is_location(line: str) -> bool:
    return re.match(r"^\s*[\w\s.]+,\s[a-zA-Z]{2}$", line) is not None


def is_start_and_end_date(line: str) -> bool:
    return (
        re.match(
            r"\b(?:Jan\.|Feb\.|Mar\.|Apr\.|May|Jun\.|Jul\.|Aug\.|Sept\.|Oct\.|Nov\.|Dec\.) \d{4} [-–] (?:Jan\.|Feb\.|Mar\.|Apr\.|May|Jun\.|Jul\.|Aug\.|Sept\.|Oct\.|Nov\.|Dec\.) \d{4}\b",
            line,
        )
        is not None
    )
