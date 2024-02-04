from ch21 import MT19937
from time import time

def inverse_right_shift(y: int, shift: int, num: int = 0xFFFFFFFF) -> int:
    mask   = (1 << shift) - 1 << (32 - shift)
    chunks = (32 + shift - 1) // shift

    for _ in range(0, chunks):
        y ^= (y >> shift) & mask & num 
        mask >>= shift

    return y

def inverse_left_shift(y: int, shift: int, num: int = 0xFFFFFFFF) -> int:
    mask   = (1 << shift) - 1
    chunks = (32 + shift - 1) // shift

    for _ in range(0, chunks):
        y ^= (y << shift) & mask & num
        mask <<= shift

    return y

def untemper(y: int) -> int:
    y = inverse_right_shift(y, MT19937.l)
    y = inverse_left_shift(y, MT19937.t, MT19937.c)
    y = inverse_left_shift(y, MT19937.s, MT19937.b)
    y = inverse_right_shift(y, MT19937.u, MT19937.d)

    return y 


if __name__ == "__main__":
    prng = MT19937(int(time()))

    rand_bytes = [ prng.extract_number() for _ in range(MT19937.n) ]
    untemped   = map(untemper, rand_bytes)

    cloned_prng = MT19937(0)
    cloned_prng.MT = list( untemped )

    a = prng.extract_number()
    b = cloned_prng.extract_number()

    assert a == b

    print("challenge 23:\n\tnext num -> ", a, "==", b)

