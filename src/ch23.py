from ch21 import MT19937, as_int32
from time import time


def inverse_right_shift(y: int, shift: int) -> int:
    mask = (1 << shift) - 1 << (32 - shift)

    print(mask)
    print((32 + shift - 1) // shift )
    


    return y

def untemper(y: int) -> int:
    y ^= (y >> MT19937.l)

    y ^= ((y >> MT19937.t) & MT19937.c)
    y ^= ((y >> MT19937.s) & MT19937.b)
    y ^= ((y << MT19937.u) & MT19937.d)

    return as_int32(y)



if __name__ == "__main__":
    prng = MT19937(int(time()))

    rand_bytes     = [ prng.extract_number() for _ in range(MT19937.n) ]
    untemped_bytes = map(untemper, rand_bytes)

    #print(list(untemped_bytes)[:5])
    a = rand_bytes[0]
    inverse_right_shift(a, MT19937.u)




