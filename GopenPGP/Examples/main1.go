package main

import (
"github.com/ProtonMail/gopenpgp/v3/crypto"
)

func main() {
	pgp := crypto.PGP()

	password := []byte("hunter2")

	pgp := crypto.PGP()
	// Encrypt data with a password
	encHandle, err := pgp.Encryption().Password(password).New()
	pgpMessage, err := encHandle.Encrypt([]byte("my message"))
	armored, err := pgpMessage.ArmorBytes()

	// Decrypt data with a password
	decHandle, err := pgp.Decryption().Password(password).New()
	decrypted, err := decHandle.Decrypt(armored, crypto.Armor)
	myMessage := decrypted.Bytes()
}