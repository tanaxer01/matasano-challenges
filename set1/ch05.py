from ch02 import xor
from itertools import cycle

def repeating_key_xor(text: bytes, key: bytes) -> bytes:
    output = bytearray()

    for i, j in zip(text, cycle(key)):
        output.append(i ^ j)

    return output

if __name__ == "__main__":
    text = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
    key = "ICE"

    res = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

    xor_bytes = repeating_key_xor(text.encode(), key.encode())

    assert xor_bytes.hex() == res

    print("challenge 5:", text)

