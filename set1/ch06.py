import base64
from ch03 import break_xor_cipher
from ch05 import repeating_key_xor

def hamming_distance(a: bytes, b: bytes) -> int:
    total = 0
    for i, j in zip(a, b):
        total += bin(i ^ j).count("1")

    return total

def keysize_edit_distance(data: bytes, keysize: int) -> float:
    chunk_cant = (len(data) + 1) // keysize

    chunk_pairs = map(lambda c: (data[c:c+keysize],data[c+keysize:c+(keysize*2)]), range(chunk_cant))
    distances   = map(lambda pair: hamming_distance(*pair)/keysize, chunk_pairs)

    return sum(distances) / chunk_cant 

#def break_repeating_key_xor(text: bytes, key: bytes) -> bytes:
#    return b""

if __name__ == "__main__":
    # 1.
    POSIBLE_KEYSIZES = range(2,41)
    # 2.
    test_dist = hamming_distance(b"this is a test", b"wokka wokka!!!")
    assert test_dist == 37
    
    with open("input/ch06.txt", "r") as file:
        data = base64.b64decode("".join([ i.rstrip() for i in file.readlines() ]))

    key_size = min(POSIBLE_KEYSIZES, key=lambda x: keysize_edit_distance(data, x))

    data_blocks = map(lambda x: data[x:x+key_size], range(0, len(data), key_size))
    transposed_blocks = map(lambda x: data[x::key_size], range(key_size))

    solved_blocks = map(break_xor_cipher,transposed_blocks)
    key = "".join( map(lambda x: chr(x[1]), solved_blocks) )
    print(key)

    decrypted_data = repeating_key_xor(data, key.encode())
