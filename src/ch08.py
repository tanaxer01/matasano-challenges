from ch07 import ECB_decrypt

if __name__ == "__main__":
    with open("input/ch08.txt", "r") as file:
        lines = map(lambda x: bytes.fromhex(x.rstrip()), file.readlines())

        decrypted_lines  = map(lambda x: ECB_decrypt(x, b"A"*16), lines)
        decrypted_chunks = map(lambda x: [ x[i:i+16] for i in range(0,len(x),16) ], decrypted_lines) 
        ecb_candidates   = list( filter(lambda x: len(set(x)) != len(x), decrypted_chunks) )

        assert len(ecb_candidates) == 1

        print("challenge 8:", end="\n\t")
        print(*ecb_candidates[0], sep="\n\t")

