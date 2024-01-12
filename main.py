import os
from pathlib import Path

def run_set(num: int):
    folder = Path(f"./set{num}")
    files  = folder.glob("*.py")

    for challenge in files:
        print(challenge)


run_set(1)
