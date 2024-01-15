
from Crypto.Random import get_random_bytes
from typing import Tuple
from ch07 import ECB_encrypt
from ch11 import is_ecb_mode
import base64


def oracle_12(input_bytes: bytes) -> bytes:
    target_bytes = base64.b64decode( b"Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")
    return ECB_encrypt(input_bytes + target_bytes, KEY)

def calc_padding_length() -> Tuple[int, int]:
    pad = b""
    empty_length = len(oracle_12(pad))
    while empty_length == len(oracle_12(pad)):
        pad += b"A"

    return len(pad), empty_length

def match_single_byte(cipher_bytes: bytes, plain_bytes: bytes, offset: int) -> int:
    posibles = { oracle_12(plain_bytes + bytes([i]))[16*offset:16*(offset+1)]:i for i in range(256) }
    assert cipher_bytes in posibles.keys()

    return posibles[cipher_bytes]

def match_block(padding: bytes, decoded: bytes, payload_length:int, offset: int):
    for _ in range(16):
        if len(decoded) == payload_length:
            break

        cipher_bytes = oracle_12(padding)[16*offset:16*(offset+1)]
        decoded_byte = match_single_byte(cipher_bytes, padding+decoded, offset)

        decoded += bytes([ decoded_byte ])
        padding = padding[:-1]

    return decoded


if __name__ == "__main__":
    KEY = get_random_bytes(16)

    pad, length = calc_padding_length()
    pad = b"A" * pad
    assert len(oracle_12(pad)) - len(oracle_12(b"")) == 16
    assert is_ecb_mode(oracle_12(b"A"*32)) == True

    decoded_bytes = b""
    padding_bytes = b"A" * 15

    payload_length = len(oracle_12(b"")) - len(pad)
    for i in range(length-2):
        decoded_bytes = match_block(padding_bytes, decoded_bytes, payload_length, i)

    print("challenge 12:")
    print(*[ "\t"+i for i in decoded_bytes.decode().split("\n") ], sep="\n")




