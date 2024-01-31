from Crypto.Random import random
from ch21 import MT19937
import time

if __name__ == "__main__":
    start = int(time.time())
    wait1 = random.randint(40, 1000)
    wait2 = random.randint(40, 1000)

    seed  = start + wait1
    now   = start + wait1 + wait2
    gen   = MT19937(seed)
    rand_bytes = gen.extract_number()

    res = -1
    for i in range(40, 1000):
        temp_gen = MT19937(now - i)
        if temp_gen.extract_number() == rand_bytes:
            res = now - i
            break

    assert res == seed

    print("challenge 22:\n\t seed -> ", res)

