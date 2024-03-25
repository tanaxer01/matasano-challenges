package set2

import (
	"Matasano/set1"
	"fmt"
	"log"
	"strings"
)

func ParseExpr(expr string) map[string]string {
	keys := make(map[string]string)
	for _, p := range strings.Split(expr, "&") {
		parts := strings.Split(p, "=")
		keys[parts[0]] = parts[1]
	}

	return keys
}

func ProfileFor(user string) string {
	user = strings.ReplaceAll(user, "&", "%26")
	user = strings.ReplaceAll(user, "=", "%3D")

	return fmt.Sprintf("email=%s&uid=10&role=user", user)
}

func oracle13(user string) []byte {
	profile := ProfileFor(user)
	return set1.ECBEncrypt([]byte(profile), KEY)
}

func Challenge13() {
	RandomBytes(KEY)

	var user string = "fooo@baar.com"
	var fake []byte = set1.Pkcs7([]byte("admin"), 16)

	normalCt := oracle13(user)
	moddedCt := oracle13("AAAAAAAAAA" + string(fake))

	for i := 0; i < 16; i++ {
		normalCt[32+i] = moddedCt[16+i]
	}

	pt := set1.ECBDecrypt(normalCt, KEY)
	cleanText, err := set1.RemovePkcs7(pt)
	if err != nil {
		log.Fatalln("Error while removing text padding:", err)
	}

	parsedExpr := ParseExpr(string(cleanText))
	if _, ok := parsedExpr["role"]; !ok {
		log.Fatalf("Role is not in profile")
	} else if parsedExpr["role"] != "admin" {
		log.Fatalf("User is not a Admin: %s", parsedExpr["role"])
	}

	log.Printf("\t[ch 13] %s", cleanText)
}
