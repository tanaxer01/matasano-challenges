

w, n , m, r = 32, 624, 397, 31
a = int("9908B0DF", 16)

f = 1812433253

MT  = [0] * n
idx = n + 1
lower_mask = (1 << r) - 1 # That is, the binary number of r 1's
upper_mask = ~lower_mask & 0xFFFF

def seed_mt(seed: int):
    index = n
    MT[0] = seed
    for i in range(1, n):
        MT[i] = (f * (MT[i-1] ^ (MT[i-1] >> (w-2))) + i) & ~20
