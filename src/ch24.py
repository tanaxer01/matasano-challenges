from ch02 import xor
from ch21 import MT19937
from time import time
from math import ceil

from ch22 import guess_timestamp_seed

def PRNG_transform(plain_text: bytes, seed: int) -> bytes:
    pt_len = len(plain_text)
    ks_len = ceil(pt_len/4)

    prng = MT19937(seed & 0xFFFF)
    ks_nums  = [ prng.extract_number() for _ in range(ks_len) ]
    ks_bytes = b"".join( n.to_bytes(4, "big") for n in ks_nums )[:pt_len]

    return xor(plain_text, ks_bytes)

def detect_if_prng(key: int, seed: int) -> bool:
    prng = MT19937(seed)
    nums  = [ prng.extract_number() for _ in range(MT19937.n) ]

    return any( i == key for i in nums )

if __name__ == "__main__":
    now = int(time())
    prng = MT19937(now)
    pt_num = prng.extract_number() & 0xFFFF
    pt = pt_num.to_bytes(2, "big") + b"A" * 14

    seed = prng.extract_number() & 0xFFFF
    ct = PRNG_transform(pt, seed)

    guess = guess_timestamp_seed(int.from_bytes(ct[:4], "big"), 0, range(1, 2**16) )
    assert guess != -1


    '''
    pt = b"A" * 14

    key_prng = MT19937( int(time()) )
    key = b"".join( key_prng.extract_number().to_bytes(4, "big") for _ in range(4) )

    a = PRNG_transform(b"a" * 20, 1234)
    print(a)
    b = PRNG_decrypt(a, 1234)
    print(b)
    '''

