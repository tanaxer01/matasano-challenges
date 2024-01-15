from itertools import starmap

def xor(a: bytes, b: bytes) -> bytes:
    return bytes( starmap(lambda i,j : i^j, zip(a, b)) )

if __name__ == "__main__":
    a = "1c0111001f010100061a024b53535009181c"
    b = "686974207468652062756c6c277320657965"
    res = "746865206b696420646f6e277420706c6179"

    a_bytes = bytes.fromhex(a)
    b_bytes = bytes.fromhex(b)

    xor_bytes = xor(a_bytes, b_bytes)

    assert xor_bytes.hex() == res

    print("challenge 2:\n\t", xor_bytes.hex())

