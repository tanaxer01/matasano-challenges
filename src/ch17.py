from os import walk
from Crypto.Random import get_random_bytes, random
import base64
from ch02 import xor
from ch10 import CBC_decrypt, CBC_encrypt
from ch15 import remove_padding

options = ([
    b"MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=",
    b"MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=",
    b"MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==",
    b"MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==",
    b"MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl",
    b"MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==",
    b"MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==",
    b"MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=",
    b"MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=",
    b"MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93"
])

def oracle_17() -> bytes:
    option = base64.b64decode( random.choice(options) )
    return CBC_encrypt(option, KEY, IV)

def valid_padding(ct: bytes) -> bool:
    pt = CBC_decrypt(ct, KEY, IV)
    try:
        remove_padding(pt)
        return True
    except AssertionError:
        return False

def filtrate_block(ct: bytes) -> bytes:
    zero = [ 0 for _ in range(16) ]

    for idx in range(1, 17):
        pad  = [ i ^ idx for i in zero ]

        for num in range(256):
            pad[-idx] = num
            mask = bytes(pad).ljust(32, b"\x00")

            if valid_padding( xor(ct, mask) ):
                # Check for false positives with \x01 pad
                if idx == 1:
                    pad[14] ^= 1
                    mask = bytes(pad).ljust(32, b"\x00")

                    if not valid_padding( xor(ct, mask) ):
                        continue

                zero[-idx] = num ^ idx
                break

    return bytes(zero)

if __name__ == "__main__":
    KEY = get_random_bytes(16)
    IV  = get_random_bytes(16)

    print("===")
    ct  = IV + oracle_17()
    pt  = b"".join(filtrate_block(ct[i:i+32]) for i in range(0, len(ct), 16))

    print("challenge 17:\n\t", pt)
