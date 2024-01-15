from Crypto.Random import get_random_bytes
from ch02 import xor
from ch10 import CBC_decrypt, CBC_encrypt

def format_input(input_string: bytes) -> bytes:
    clean_input = input_string.replace(b";",b"").replace(b"=",b"")
    return b"comment1=cooking%20MCs;userdata=" + clean_input + b";comment2=%20like%20a%20pound%20of%20bacon"

def encrypt_input(input_string: bytes) -> bytes:
    input_bytes = format_input(input_string)
    return CBC_encrypt(input_bytes, KEY, IV)

def validate_creds(input_bytes: bytes) -> bool:
    plain_bytes = CBC_decrypt(input_bytes, KEY, IV)
    return b"admin=true" in plain_bytes


if __name__ == "__main__":
    KEY = get_random_bytes(16)
    IV  = get_random_bytes(16)

    input_bytes = b"A" * 16
    ct = encrypt_input(input_bytes*2)

    moded_bytes  = xor(input_bytes, b";admin=true".rjust(16, b"A"))
    padded_bytes = moded_bytes.rjust(48, b"\x00").ljust(len(ct), b"\x00")
    admin_input  = xor(padded_bytes, ct)


    print("challenge 16:\n\t 'admin = true' in ct? ", validate_creds(admin_input) )

