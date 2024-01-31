
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
        self.MT = [ 0 for _ in range(self.n) ]
        self.index = self.n + 1

        self.lower_mask = 0xFFF & ((1 << self.r) - 1)
        self.upper_mask = 0xFFF & ~self.lower_mask

        self.seed_mt(seed)


    def seed_mt(self, seed: int) -> None:
        ''' Initialize the generator from a seed '''
        self.index = self.n
        self.MT[0] = seed

        for i in range(1, self.n):
            self.MT[i] = as_int32(self.f * (self.MT[i-1] ^ (self.MT[i-1] >> (self.w-2))) + 1)

    def extract_number(self) -> int:
        ''' 
            Extract a tempered value based on MT[index]
            calling twist() every n numbers 
        '''

        if self.index >= self.n:
            assert self.index <= self.n, "Generator was never seeded"
            self.twist()

        y  = self.MT[self.index]
        y ^= ((y >> self.u) & self.d)
        y ^= ((y << self.s) & self.b)
        y ^= ((y << self.t) & self.c)
        y ^= (y >> self.l)

        self.index = self.index+1
        return as_int32(y)

    def twist(self) -> None:
        ''' Generate the next n values from the series x_i '''

        for i in range(self.n):
            x  = (self.MT[i] & self.upper_mask) | (self.MT[(i+1) % self.n] & self.lower_mask)
            xA = x >> 1 if x % 2 == 0 else (x >> 1) ^ self.a
            self.MT[i] = self.MT[(i + self.m) % self.n] ^ xA

        self.index = 0

if __name__ == "__main__":
    genA, genB = MT19937(10), MT19937(10)
    A = [ genA.extract_number() for _ in range(10) ]
    B = [ genB.extract_number() for _ in range(10) ]

    print("challenge 21:", A, B, [ i == j for i, j in zip(A,B) ], sep="\n\t")


