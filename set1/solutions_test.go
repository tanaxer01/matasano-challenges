package set1

import (
	"Matasano/utils"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var input    string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	var expected string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	var base64   string = utils.To_base64(  utils.From_hex(input) )
	var hex      []byte = utils.To_hex( utils.From_base64(base64) )

	if base64 != expected || string(hex) != input {
		t.Errorf("Got: %v \nAnd: %v", base64, string(hex))
	} else {
		t.Logf("[1] Base64 = %v", base64)
		t.Logf("[2] Hex    = %v", string(hex))
	}
}

func TestChallenge2(t *testing.T) {
	var expected string = "746865206b696420646f6e277420706c6179"

	intsA    := utils.From_hex("1c0111001f010100061a024b53535009181c")
	intsB    := utils.From_hex("686974207468652062756c6c277320657965")
	received := utils.To_hex(utils.XOR(intsA,intsB))

	if expected != string(received) {
		t.Errorf("Expected: %v, got: %v",expected, string(received))
	} else {
		t.Logf("[2] XOR => %v", string(received))
	}
}

func TestChallenge3(t *testing.T) {
	var input    string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	var expected string = "Cooking MC's like a pound of bacon"
	received, char ,maximum := Xor_cipher(input)
	
	if received != expected {
		t.Errorf("Expected: %v, got: %v",received,received)
	} else {
		t.Logf("[3] XOR_Cipher => %v %v %c", received, maximum, char)
	}
}

func TestChallenge4(t *testing.T) {
	var expected string = "Now that the party is jumping\n"
	var received string = DetectXor("test_data/4.txt")

	if received != expected {
		t.Errorf("Expected: %v, got: %v",received,received)
	} else {
		t.Logf("[4] DetectXOR => %v", received)
	}
}

func TestChallenge5(t *testing.T) {
	var expected string = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	var input    string = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	var key      string = "ICE"
	var received string = RepeatingKeyXor(key,input)

	if received != expected {
		t.Errorf("Expected: %v, got: %v",received,expected)
	} else {
		t.Logf("[5] RepeatingKeyXOR => %v", received)
	}

}

func TestChallange6(t *testing.T) {
	var path     string = "test_data/6.txt"
	var expected string = "Terminator X: Bring the noise"
	key, _ := BreakRepetingKeyXor(path)

	if key != expected {
		t.Errorf("Expected: %v, got: %v", key, expected)
	} else {
		t.Logf("[6] BreakRepetingkeyXOR => %v", key)
	}
}

func TestChallenge7(t *testing.T) {
	var key   string = "YELLOW SUBMARINE"
	var input string = utils.ReadFile("test_data/7.txt")
	var input_decoded []byte = utils.From_base64(input)

	decoded  := AES_decode([]byte(key), input_decoded)
	received := AES_encode([]byte(key), decoded)

	t.Logf("AES ==> %v",string(received) == string(input_decoded))

}

func TestChallange8(t *testing.T) {
	var path = "test_data/8.txt"
	text, _ := DetectECB(path)

	t.Logf("[8] DetectECB => %v", text)
}
