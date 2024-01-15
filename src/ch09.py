
def pkcs7(input_bytes: bytes, target_size: int = 16) -> bytes:
    padding_size = target_size - len(input_bytes)
    return input_bytes + bytes([padding_size]) * padding_size

if __name__ == "__main__":
    input_bytes  = b"YELLOW SUBMARINE"
    padded_bytes = pkcs7(input_bytes, 20)

    print("challenge 9:\n\t", padded_bytes)
