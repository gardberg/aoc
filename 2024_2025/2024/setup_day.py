import argparse
import os
import shutil
from utils import get_input
# sets up a folder for the day

BASE_FOLDER = "days"

def setup_day(day: int, year: int = 2024):
    folder_name = f"{BASE_FOLDER}/{day}"
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
    parser.add_argument("year", type=int, default=2024)
    args = parser.parse_args()

    setup_day(args.day, args.year)
