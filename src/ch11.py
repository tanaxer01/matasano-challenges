from Crypto.Random import get_random_bytes, random
from ch07 import ECB_encrypt
from ch10 import CBC_encrypt

RES = None

def oracle_11(input_bytes: bytes) -> bytes:
    prefix = get_random_bytes(random.randrange(5,11))
    sufix  = get_random_bytes(random.randrange(5,11))
    plain_bytes = prefix + input_bytes + sufix

    if random.getrandbits(1):
        RES = True
        return CBC_encrypt(plain_bytes, KEY, IV)
    else:
        RES = False
        return ECB_encrypt(plain_bytes, KEY)


def is_ecb_mode(cipher_bytes: bytes):
    blocks = list(map(lambda x: cipher_bytes[x:x+16], range(0, len(cipher_bytes), 16)))
    return len(set(blocks)) != len(blocks)


if __name__ == "__main__":
    KEY = get_random_bytes(16)
    IV  = get_random_bytes(16)

    cipher_bytes = oracle_11(b"A"*32)

    assert is_ecb_mode(cipher_bytes) != RES
    print("challenge 11:\n\t", "ECB" if is_ecb_mode(cipher_bytes) else "CBC" )

