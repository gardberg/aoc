import requests
import os

URL = "https://adventofcode.com/{year}/day/{day}/input"

def check_session_id():
    if not os.environ.get("AOC_SESSION"):
        if os.path.exists("session_id"):
            with open("session_id", "r") as f:
                session_id = f.read()
                os.environ["AOC_SESSION"] = session_id
        else:
            set_session_id()

check_session_id()

def set_session_id():
    session_id = input("Enter your session ID: ")
    with open("session_id", "w") as f:
        f.write(session_id)
    os.environ["AOC_SESSION"] = session_id

def get_session_id() -> str:
    with open("session_id", "r") as f:
        return f.read()

def get_input(day: int, year: int = 2024) -> str:
    url = URL.format(year=year, day=day)
    response = requests.get(url, cookies={"session": os.environ["AOC_SESSION"]})
    return response.text
