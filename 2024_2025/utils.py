import argparse
import os
import shutil

import requests
from dotenv import load_dotenv

load_dotenv()

URL = "https://adventofcode.com/{year}/day/{day}/input"


def get_input(day: int, year: int = 2025) -> str:
    url = URL.format(year=year, day=day)
    response = requests.get(url, cookies={"session": os.environ["AOC_SESSION_ID"]})
    return response.text


# sets up a folder for the day

BASE_FOLDER = "days"


def setup_day(day: int, year: int = 2025):
    folder_name = f"{year}/{BASE_FOLDER}/{day}"
    os.makedirs(folder_name, exist_ok=False)

    # get input, remove last newline
    inp = get_input(day, year)
    inp = inp[:-1]
    with open(f"{folder_name}/input.txt", "w") as f:
        f.write(inp)

    # copy template
    shutil.copy("template.py", f"{folder_name}/main.py")

    # make and empty example.txt
    with open(f"{folder_name}/example.txt", "w") as f:
        f.write("")


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("day", type=int)
    parser.add_argument("year", type=int, nargs="?", default=2025)
    args = parser.parse_args()

    setup_day(args.day, args.year)
