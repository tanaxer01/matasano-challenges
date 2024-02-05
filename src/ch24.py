from ch02 import xor
from ch03 import english_score
from ch21 import MT19937
from time import time
from math import ceil

def PRNG_transform(text: bytes, seed: int) -> bytes:
    new_len = len(text)
    ks_len = ceil(new_len/4)

    assert seed <= 0xFFFF
    prng = MT19937(seed)
    ks_nums  = [ prng.extract_number() for _ in range(ks_len) ]
    ks_bytes = b"".join( n.to_bytes(4, "big") for n in ks_nums )[:new_len]

    return xor(text, ks_bytes)

def guess_16_bit_seed(cipher_text: bytes):
    max_score = seed = 0
    for i in range(0xFFFF):
        posible_pt = PRNG_transform(cipher_text, i)
        try:
            score = english_score( posible_pt[2:].decode() )
            if score > max_score:
                max_score = score
                seed = i

                print(posible_pt.hex())
        except:
            pass

    return seed

if __name__ == "__main__":
    now = int(time())
    prng = MT19937(now)

    num = prng.extract_number() & 0xFFFF
    known_text = num.to_bytes(2, "big") + (b"A" * 14)

    prng = MT19937(now)
    seed = prng.extract_number() & 0xFFFF
    ct   = PRNG_transform(known_text, seed)

    guess  = guess_16_bit_seed(ct)
    assert guess == seed

    print("challenge 24:\n\tseed ->",guess)

