package pkg

import (
	"bytes"
	"testing"
)

func TestGenerateRandomBytes(t *testing.T) {
	mapB := make(map[int][]byte)
	for i := 0; i < 10; i++ {
		if randBytes, err := GenerateRandomBytes(100); err != nil {
			t.Error(err)
		} else {
			mapB[i] = randBytes
		}
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i != j {
				a, _ := mapB[i]
				b, _ := mapB[j]
				if bytes.Equal(a, b) {
					t.Errorf("Generate random bytes generate equals byte array.")
				}
			}
		}
	}
}

func TestGenerateRandomString(t *testing.T) {
	mapS := make(map[int]string)
	for i := 0; i < 10; i++ {
		if randS, err := GenerateRandomString(10); err != nil {
			t.Error(err)
		} else {
			mapS[i] = randS
		}
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i != j {
				a, _ := mapS[i]
				b, _ := mapS[j]
				if a == b {
					t.Errorf("Generate random stirng generate equals string.")
				}
			}
		}
	}
}

type TestEqualsStruct struct {
	A []byte
}

func TestEquals(t *testing.T) {
	structA := TestEqualsStruct{}
	structB := TestEqualsStruct{}

	bytesArr, _ := GenerateRandomBytes(100)
	structA.A = bytesArr
	structB.A = bytesArr

	if !Equals(structA, structB) {
		t.Error("Equals function no working")
	}
}
