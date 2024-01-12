from Crypto.Cipher import AES
import base64

key = b"YELLOW SUBMARINE"

def aes_decrypt(cipher_text: bytes, key: bytes) -> bytes:
    cipher = AES.new(key, AES.MODE_ECB)
    return cipher.decrypt(cipher_text)
    
def main():
    with open("input/ch07.txt", "r") as file:
        data = base64.b64decode("".join([ i.rstrip() for i in file.readlines() ]))
        plain_text = aes_decrypt(data, key).decode()

        print("challenge 7:\t", *plain_text.split("\n")[:5], sep="\n")

if __name__ == "__main__":
    main()
