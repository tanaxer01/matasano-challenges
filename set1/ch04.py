from ch03 import break_xor_cipher

def find_xored_line(file: str) -> str:
    with open(file, "r") as challenge_file:
        lines   = map(lambda l: bytes.fromhex(l.rstrip()), challenge_file.readlines())
        options = map(break_xor_cipher, lines)

        xored_line = max(options, key=lambda x: x[1])

    return xored_line


if __name__ == "__main__":
    xored_line = find_xored_line("../input/ch04.txt")
    print("challenge 4: ", xored_line[0])
