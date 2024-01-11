from Crypto.Cipher import AES
import base64

key = b"YELLOW SUBMARINE"

def aes_decrypt(cipher_text: bytes, key: bytes) -> bytes:
    cipher = AES.new(key, AES.MODE_ECB)
    return cipher.decrypt(cipher_text)
    
with open("input/ch07.txt", "r") as file:
    data = base64.b64decode("".join([ i.rstrip() for i in file.readlines() ]))


cipher = AES.new(key, AES.MODE_ECB)


plain_text = aes_decrypt(data, key)

print(plain_text.decode())
