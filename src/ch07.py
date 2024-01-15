from Crypto.Cipher import AES
import base64

from ch09 import pkcs7

key = b"YELLOW SUBMARINE"

def ECB_decrypt(cipher_bytes: bytes, key: bytes) -> bytes:
    cipher = AES.new(key, AES.MODE_ECB)
    return cipher.decrypt(cipher_bytes)

def ECB_encrypt(plain_bytes: bytes, key: bytes) -> bytes:
    cipher = AES.new(key, AES.MODE_ECB)
    if len(plain_bytes)%16 != 0:
        plain_bytes = pkcs7(plain_bytes, ((len(plain_bytes)//16)+1)*16 )

    return cipher.encrypt(plain_bytes)
    
if __name__ == "__main__":
    with open("input/ch07.txt", "r") as file:
        data = base64.b64decode("".join([ i.rstrip() for i in file.readlines() ]))
        plain_text = ECB_decrypt(data, key).decode()

        print("challenge 7:")
        print(*[ "\t"+i for i in plain_text.split("\n")[:3] ], sep="\n")

