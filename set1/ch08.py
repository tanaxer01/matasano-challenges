
from ch07 import aes_decrypt

def main():
    with open("input/ch08.txt", "r") as file:
        lines = map(lambda x: bytes.fromhex(x.rstrip()), file.readlines())

        decrypted_lines  = map(lambda x: aes_decrypt(x, b"A"*16), lines)
        decrypted_chunks = map(lambda x: [ x[i:i+16] for i in range(0,len(x),16) ], decrypted_lines) 
        ecb_candidates   = list( filter(lambda x: len(set(x)) != len(x), decrypted_chunks) )

        assert len(ecb_candidates) == 1

        print("challenge 8:\t", b"".join(ecb_candidates[0]))

if __name__ == "__main__":
    main()
