
import base64
import math
from typing import List
from Crypto.Random import get_random_bytes
from ch18 import CRT_encrypt
from ch19 import break_fixed_nonce_crt


def encrypt_lines() -> List[bytes]:
    with open("input/ch20.txt", "r") as cipher_file:
        ct_lines = map(lambda x: CRT_encrypt(base64.b64decode(x), KEY, b"\x00"), cipher_file.readlines())
        return list(ct_lines)

if __name__ == "__main__":
    KEY = get_random_bytes(16)
    ct = encrypt_lines()
    pt = break_fixed_nonce_crt(ct)

    print("challenge 20:", *[ i.decode() for i in pt ][:5], sep="\n\t")


