from Crypto.Random import random
from ch21 import MT19937
import time

def guess_timestamp_seed(target_num: int, now: int, guess_range: range):
    res = -1
    for i in guess_range:
        temp_gen = MT19937(now - i)
        if temp_gen.extract_number() == target_num:
            res = now - i
            break

    return res


if __name__ == "__main__":
    start = int(time.time())
    wait1 = random.randint(40, 1000)
    wait2 = random.randint(40, 1000)

    seed  = start + wait1
    now   = start + wait1 + wait2
    gen   = MT19937(seed)
    rand_bytes = gen.extract_number()

    res = guess_timestamp_seed(rand_bytes, now, range(40, 1000))
    assert res == seed

    print("challenge 22:\n\t seed -> ", res)

