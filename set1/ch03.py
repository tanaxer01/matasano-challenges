from collections import Counter
from typing import Tuple

def english_score(text: str) -> int:
    score = 0
    freqs = Counter(text)

    for letter in "ETAOIN SHRDLU":
        if letter in freqs:
            score += freqs[letter]

        if letter.lower() in freqs:
            score += freqs[letter.lower()]

    return score

def break_xor_cipher(ciphered: bytes) -> Tuple[str, int]:
    max, max_str = 0, ""
    for i in range(256):
        try:
            pos = bytes([ i^k for k in ciphered ]).decode()
            score = english_score(pos)

            if score > max:
                max, max_str = score, pos
        except:
            pass

    return max_str, max

if __name__ == "__main__": 
    hex = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
    hex_bytes = bytes.fromhex(hex)

    text, _ = break_xor_cipher(hex_bytes)
    print("challenge 03:\t", text)

