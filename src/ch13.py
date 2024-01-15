from Crypto.Random import get_random_bytes
from ch07 import ECB_encrypt, ECB_decrypt
from ch09 import pkcs7

def parse_cookie(cookie: str) -> dict:
    key_values = map(lambda x: x.split("="), cookie.split("&"))
    return { i: j for i, j in key_values }

def profile_for(email: str) -> str:
    clean_email = email.replace("&", "").replace("=", "")
    return f"email={clean_email}&uid:10&role=user"

def oracle_13(email: str):
    return ECB_encrypt(profile_for(email).encode(), KEY)

if __name__ == "__main__":
    KEY = get_random_bytes(16)

    test_cookie = "foo=bar&baz=qux&zap=zazzle"
    assert parse_cookie(test_cookie) == { "foo": "bar", "baz": "qux", "zap": "zazzle" }

    email = "AAAA@BBBB.com"
    nrml_profile = profile_for(email)
    fake_profile = nrml_profile.replace("user", "admin")

    nrml_cipher = oracle_13(email)
    fake_cipher = oracle_13("A"*10 + pkcs7(b"admin", 16).decode())

    nrml_blocks = [ nrml_cipher[i:i+16] for i in range(0,len(nrml_profile),16) ]
    fake_blocks = [ fake_cipher[i:i+16] for i in range(0,len(fake_profile),16) ]
    nrml_blocks[2] = fake_blocks[1]

    decrypted_profile = ECB_decrypt(b"".join(nrml_blocks), KEY)

    assert b"role=admin" in decrypted_profile

    print("challenge 13:\n\t", decrypted_profile.decode())
