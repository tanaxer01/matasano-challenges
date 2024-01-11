from collections import Counter
from typing import Tuple

def english_score(text: str) -> float:
    score = 0
    freqs = Counter(text)
    values = {
        'a': 0.08167, 'b': 0.01492, 'c': 0.02782, 'd': 0.04253, 'e': 0.12702,
        'f': 0.02228, 'g': 0.02015, 'h': 0.06094, 'i': 0.06094, 'j': 0.00153,
        'k': 0.00772, 'l': 0.04025, 'm': 0.02406, 'n': 0.06749, 'o': 0.07507,
        'p': 0.01929, 'q': 0.00095, 'r': 0.05987, 's': 0.06327, 't': 0.09056,
    	'u': 0.02758, 'v': 0.00978, 'w': 0.02360, 'x': 0.00150, 'y': 0.01974,
        'z': 0.00074, ' ': 0.13000
    }

    for letter, freq in Counter(text).items():
        if letter.lower() in values:
            score += values[letter.lower()] * freq

    return score

def break_xor_cipher(ciphered: bytes) -> Tuple[str, int, float]:
    max_str = ""
    max_score = max_chr = 0
    for i in range(256):
        try:
            pos = bytes([ i^k for k in ciphered ]).decode()
            score = english_score(pos)

            if score > max_score:
                max_score, max_str, max_chr = score, pos, i
        except:
            pass

    return max_str, max_chr, max_score

if __name__ == "__main__": 
    hex = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
    hex_bytes = bytes.fromhex(hex)

    text, _, _ = break_xor_cipher(hex_bytes)
    print("challenge 03:\t", text)

