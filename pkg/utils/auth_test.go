package utils

import "testing"

func TestAuth(t *testing.T) {
	samplePass := "test"

	passConf := &PasswordConfig{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}
	pass, err := GeneratePassword(passConf, samplePass)
	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	passN, errN := GeneratePassword(passConf, samplePass)
	if errN != nil {
		t.Errorf("unexpected error %s", errN.Error())
	}
	if passN == pass {
		t.Errorf("passwords must be different %s vs %s", pass, passN)
	}

	equal, errC := ComparePassword(samplePass, pass)
	if errC != nil {
		t.Errorf("unexpected error %s", errC.Error())
	}

	if !equal {
		t.Errorf("mismatch hash and pass. pass %s hash %s", samplePass, pass)
	}

	equal, errC = ComparePassword(samplePass, passN)
	if errC != nil {
		t.Errorf("unexpected error %s", errC.Error())
	}

	if !equal {
		t.Errorf("mismatch hash and pass. pass %s hash %s", samplePass, passN)
	}

	equal, errC = ComparePassword(samplePass+"1", passN)
	if errC != nil {
		t.Errorf("unexpected error %s", errC.Error())
	}

	if equal {
		t.Errorf("pass invalid, but match hash %s hash %s", samplePass+"1", passN)
	}

	equal, errC = ComparePassword(samplePass, "passN")
	if errC == nil {
		t.Error("error expected, got nill")
	}

	if equal {
		t.Errorf("hash invalid, but match hash %s hash %s", samplePass+"1", passN)
	}
}
