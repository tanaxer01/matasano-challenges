from typing import List

as_int32 = lambda x: x & 0xFFFFFFFF

class MT19937:
    w, n, m, r = (32, 624, 397, 31)
    a    = 0x9908B0DF
    u, d = 11, 0xFFFFFFFF
    s, b =  7, 0x9D2C5680
    t, c = 15, 0xEFC60000
    l    = 18
    f    = 1812433253

    def __init__(self, seed: int) -> None:
        ''' Create a length n array to store the state of the generator. '''
        self.MT = self.seed_mt(seed)
        self.index = self.n + 1

        self.lower_mask = (1 << self.r) - 1
        self.upper_mask = as_int32(~self.lower_mask)

    def seed_mt(self, seed: int) -> List[int]:
        ''' Initialize the generator from a seed '''

        MT = [ as_int32(seed) ] 
        for i in range(1, self.n):
            #MT.append( as_int32(self.f * (MT[i-1] ^ (MT[i-1] >> (self.w-2))) + i) )
            MT.append( as_int32((self.f * (MT[i - 1] ^ (MT[i - 1] >> (self.w - 2))) + i)) )

        return MT

    def extract_number(self) -> int:
        ''' 
            Extract a tempered value based on MT[index]
            calling twist() every n numbers 
        '''

        if self.index >= self.n:
            self.twist()
            self.index = 0

        y  = self.MT[ self.index ]
        y ^= (y >> self.u) & self.d
        y ^= (y << self.s) & self.b
        y ^= (y << self.t) & self.c
        y ^=  y >> self.l

        self.index += 1
        return as_int32(y)

    def twist(self) -> None:
        ''' Generate the next n values from the series x_i '''

        for i in range(self.n):
            x  = (self.MT[i] & self.upper_mask) + (self.MT[(i+1) % self.n] & self.lower_mask)
            xA = x >> 1
            if x % 2 != 0:
                xA ^= self.a

            self.MT[i] = self.MT[(i + self.m) % self.n] ^ xA


if __name__ == "__main__":
    genA, genB = MT19937(10), MT19937(10)
    A = [ genA.extract_number() for _ in range(10) ]
    B = [ genB.extract_number() for _ in range(10) ]

    print("challenge 21:", A, B, [ i == j for i, j in zip(A,B) ], sep="\n\t")

