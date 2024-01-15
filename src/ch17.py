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

def filtrate_single_byte(pad_num: int) -> int:
    ct  = oracle_17()
    c2  = list()
    res = None

    for j in range(1, 256):
        mod = b"\x00"*15 + bytes([j])
        pad = b"\x00"*(len(ct)//16 - 32) + mod + b"\x00"*16
        print(pad.hex())

        xored_ct = xor(ct, pad)

        if valid_padding(xored_ct):
            c2.append(j ^ ct[-16])
            res = j ^ ct[-16]
            break


    assert res != None
    return res


if __name__ == "__main__":
    KEY = get_random_bytes(16)
    IV  = get_random_bytes(16)

    ct  = oracle_17()
    #a = aaa(len(ct)//16)

    print( filtrate_single_byte(1) )
    print( chr(filtrate_single_byte(1)) )




