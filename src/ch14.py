from Crypto.Random import get_random_bytes, random
from ch07 import ECB_encrypt
from itertools import count
from typing import Tuple
import base64

def oracle_14(input_bytes: bytes) -> bytes:
    target_bytes = base64.b64decode( b"Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")
    return ECB_encrypt(PREFIX + input_bytes + target_bytes, KEY)

def calc_prefix_length() -> Tuple[int, int]:
    pad = 0
    prev_ct, curr_ct = oracle_14(b""), oracle_14(b"A")
    eq_blocks = sum( prev_ct[i:i+16] == curr_ct[i:i+16] for i in range(0, len(prev_ct), 16))

    for pad in count():
        prev_ct, curr_ct = curr_ct, oracle_14(b"A"*pad)
        new_eq_blocks = sum( prev_ct[i:i+16] == curr_ct[i:i+16] for i in range(0, len(prev_ct), 16))

        if new_eq_blocks - eq_blocks > 1:
            eq_blocks = new_eq_blocks
            break

    return  (eq_blocks * 16) - pad + 1, pad - 1

def match_single_byte(cipher_bytes: bytes, plain_bytes: bytes, offset: int) -> int:
    posibles = { oracle_14(plain_bytes + bytes([i]))[16*offset:16*(offset+1)]:i for i in range(256) }
    #assert cipher_bytes in posibles.keys()

    if cipher_bytes in posibles.keys():
        return posibles[cipher_bytes]
    else:
        return -1

def match_block(padding: bytes, decoded: bytes, payload_length:int, offset: int):
    for _ in range(16):
        if len(decoded) == payload_length:
            break

        cipher_bytes = oracle_14(padding)[16*offset:16*(offset+1)]
        decoded_byte = match_single_byte(cipher_bytes, padding+decoded, offset)
        if decoded_byte == -1:
            break

        decoded += bytes([ decoded_byte ])
        padding = padding[:-1]

    return decoded

if __name__ == "__main__":
    KEY = get_random_bytes(16)
    PREFIX = get_random_bytes(random.randint(1,16))

    prefix_len, padding_len = calc_prefix_length()
    payload_len = len(oracle_14(b"")) - prefix_len
    decoded_bytes = b""

    start = (prefix_len+padding_len)//16
    for i in range(start, start + payload_len//16):
        decoded_bytes = match_block(b"B"*padding_len + b"A"*15, decoded_bytes, payload_len, i)

    print("challenge 14:")
    print(*[ "\t"+i for i in decoded_bytes.decode().split("\n") ], sep="\n")

