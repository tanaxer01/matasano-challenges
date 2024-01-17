from Crypto.Cipher import AES
from Crypto.Util   import Counter
from ch02 import xor
import base64

def CRT_decrypt(cipher_text: bytes, key: bytes, nonce: bytes) -> bytes:
    ctr = Counter.new(64, prefix=nonce.rjust(8,b"\x00"), initial_value=0, little_endian=True)
    cipher = AES.new(key, AES.MODE_CTR, counter=ctr)

    return cipher.decrypt(cipher_text)

def CRT_encrypt(cipher_text: bytes, key: bytes, nonce: bytes) -> bytes:
    ctr = Counter.new(64, prefix=nonce.rjust(8,b"\x00"), initial_value=0, little_endian=True)
    cipher = AES.new(key, AES.MODE_CTR, counter=ctr)
    return cipher.encrypt(cipher_text)

def _crt_decrypt(ct: bytes, key: bytes, nonce: bytes) -> bytes:
    blocks = [ ct[i:i+16] for i in range(0,len(ct),16) ]
    countr = [ nonce.rjust(8, b"\x00") + i.to_bytes(8, "little") for i in range(len(blocks)) ]

    cipher = AES.new(key, AES.MODE_ECB)
    keystream = cipher.encrypt(b"".join(countr))

    return xor(ct, keystream[:len(ct)])

if __name__ == "__main__":
    ct    = base64.b64decode(b"L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ==")
    key   = b"YELLOW SUBMARINE"
    nonce = b"\x00"

    pt = CRT_decrypt(ct, key, nonce)
    print("challenge 18:\n\t", pt, "\n\t", _crt_decrypt(ct, key, nonce))
    
