package encrypt

import (
	"testing"
	"fmt"
)

func TestDESEncrypt(t *testing.T) {
	src := `cagent=81288128/\\\\/method=tc`
	key := "12341234"
	fmt.Println(len([]byte(key)))
	t.Logf("[%v]\n", DesEcbPkc5Encrypt(src, key))
}

func TestAesEcbEncrypt(t *testing.T) {
	src := "hell word"
	key := "1457C7D93CDBC19D"
	t.Log(AesEcbEncrypt(src, key))
}
