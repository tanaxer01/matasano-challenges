from Crypto.Cipher import AES
from ch02 import xor
from ch09 import pkcs7
import base64

def CBC_decrypt(cipher_bytes: bytes, key: bytes, iv: bytes) -> bytes:
    cipher = AES.new(key, AES.MODE_CBC, iv)
    return cipher.decrypt(cipher_bytes)

def CBC_encrypt(plain_bytes: bytes, key: bytes, iv: bytes) -> bytes:
    if len(plain_bytes)%16 != 0:
        plain_bytes = pkcs7(plain_bytes, ((len(plain_bytes)//16)+1)*16)

    cipher = AES.new(key, AES.MODE_CBC, iv)
    return cipher.encrypt(plain_bytes)

if __name__ == "__main__":
    key = b"YELLOW SUBMARINE"
    iv  = bytes.fromhex("00" * 16)
    
    with open("input/ch10.txt", "r") as file:
        data = base64.b64decode("".join([ i.rstrip() for i in file.readlines() ]))
        plain_bytes = CBC_decrypt(data, key, iv)
        print("challenge 10:", end="\n\t")
        print(*[ i for i in plain_bytes.decode().split("\n")[:5] ], sep="\n\t")

