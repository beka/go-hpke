package hpke

import (
	"bytes"
	"crypto/rand"
	"testing"
)

func TestBase(t *testing.T) {
	original := []byte("Beauty is truth, truth beauty")
	aad := []byte("that is all // Ye know on earth, and all ye need to know")
	info := []byte("Ode on a Grecian Urn")

	for id, suite := range ciphersuites {
		skR, pkR, err := suite.KEM.Generate(rand.Reader)
		if err != nil {
			t.Fatalf("[%d] Error generating DH key pair: %s", id, err)
		}
		enc, encrypted, err := Seal(suite, rand.Reader, pkR, info, aad, original)
		if err != nil {
			t.Fatalf("[%d] Error in Seal: %s", id, err)
		}

		decrypted, err := Open(suite, skR, enc, info, aad, encrypted)
		if err != nil {
			t.Fatalf("[%d] Error in Open: %s", id, err)
		}

		if !bytes.Equal(decrypted, original) {
			t.Fatalf("[%d] Incorrect decryption: [%x] != [%x]", id, decrypted, original)
		}
	}
}

func TestPSK(t *testing.T) {
	psk := []byte("mellon")
	pskID := []byte("Ennyn Durin aran Moria")
	original := []byte("Beauty is truth, truth beauty")
	aad := []byte("that is all // Ye know on earth, and all ye need to know")
	info := []byte("Ode on a Grecian Urn")

	for id, suite := range ciphersuites {
		skR, pkR, err := suite.KEM.Generate(rand.Reader)
		if err != nil {
			t.Fatalf("[%d] Error generating DH key pair: %s", id, err)
		}

		enc, encrypted, err := SealPSK(suite, rand.Reader, pkR, psk, pskID, info, aad, original)
		if err != nil {
			t.Fatalf("[%d] Error in Seal: %s", id, err)
		}

		decrypted, err := OpenPSK(suite, skR, enc, psk, pskID, info, aad, encrypted)
		if err != nil {
			t.Fatalf("[%d] Error in Open: %s", id, err)
		}

		if !bytes.Equal(decrypted, original) {
			t.Fatalf("[%d] Incorrect decryption: [%x] != [%x]", id, decrypted, original)
		}
	}
}

func TestAuth(t *testing.T) {
	original := []byte("Beauty is truth, truth beauty")
	aad := []byte("that is all // Ye know on earth, and all ye need to know")
	info := []byte("Ode on a Grecian Urn")

	for id, suite := range ciphersuites {
		skI, pkI, err := suite.KEM.Generate(rand.Reader)
		if err != nil {
			t.Fatalf("[%d] Error generating initiator DH key pair: %s", id, err)
		}

		skR, pkR, err := suite.KEM.Generate(rand.Reader)
		if err != nil {
			t.Fatalf("[%d] Error generating responder DH key pair: %s", id, err)
		}

		enc, encrypted, err := SealAuth(suite, rand.Reader, pkR, skI, info, aad, original)
		if err != nil {
			t.Fatalf("[%d] Error in Seal: %s", id, err)
		}

		decrypted, err := OpenAuth(suite, skR, pkI, enc, info, aad, encrypted)
		if err != nil {
			t.Fatalf("[%d] Error in Open: %s", id, err)
		}

		if !bytes.Equal(decrypted, original) {
			t.Fatalf("[%d] Incorrect decryption: [%x] != [%x]", id, decrypted, original)
		}
	}
}