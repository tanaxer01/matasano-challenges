import base64

def break_repeating_key_xor(text: bytes, key: bytes) -> bytes:
    return b""

def hamming_distance(a: bytes, b: bytes) -> int:
    total = 0
    for i, j in zip(a, b):
        total += bin(i ^ j).count("1")

    return total

if __name__ == "__main__":
    test_dist = hamming_distance(b"this is a test", b"wokka wokka!!!")
    print(test_dist)
    
    with open("../input/ch06.txt", "r") as file:
        data = base64.b64decode("".join([ i.rstrip() for i in file.readlines() ]))
        
    min_dist, min_key = None, None
    for KEYSIZE in range(2, 41):
        dist = hamming_distance(data[:KEYSIZE], data[KEYSIZE:KEYSIZE*2])/KEYSIZE

        if min_dist == None or min_dist > dist:
            min_dist = dist
            min_key = KEYSIZE

    print(min_dist, min_key)

    tansposed_data = [ ]

    #xored_line = find_xored_line("../input/ch06.txt")
    #print("challenge 4: ", xored_line[0])
