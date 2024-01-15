import base64

def main():
    hex = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    b64 = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

    hex_bytes = bytes.fromhex(hex)
    b64_bytes = base64.b64encode(hex_bytes)

    assert b64_bytes.decode() == b64

    print("challenge1:\n\t", hex_bytes)


if __name__ == "__main__":
    main()
