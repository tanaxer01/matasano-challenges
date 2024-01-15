
def remove_padding(input_bytes: bytes) -> bytes:
    pad_length = int(input_bytes[-1])
    assert all([ i == pad_length for i in input_bytes[-1 * pad_length:] ]) == True

    return input_bytes[:-pad_length]

if __name__ == "__main__":
    valid_padding = b"ICE ICE BABY\x04\x04\x04\x04"
    invalid_padding = b"ICE ICE BABY\x05\x05\x05\x05"
    invalid_padding2 = b"ICE ICE BABY\x01\x02\x03\x04"


    print("challenge15:\n\t", remove_padding(valid_padding))
    try:
        print("\t", remove_padding(invalid_padding))
    except AssertionError:
        print("\t", invalid_padding, "has invalid padding")

    try:
        print("\t", remove_padding(invalid_padding2))
    except AssertionError:
        print("\t", invalid_padding2, "has invalid padding")

