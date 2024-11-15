package main

import (
	"github.com/ProtonMail/gopenpgp/v3/constants"
	"fmt"
	"github.com/ProtonMail/gopenpgp/v3/crypto"
	)
	
const (
  name = "Max Mustermann"
  email = "max.mustermann@example.com"
  passphrase = []byte("LongSecret")
)

pgpDefault := crypto.PGPWithProfile(profile.Default())
pgp4880 := crypto.PGPWithProfile(profile.RFC4880())
pgpCryptoRefresh := crypto.PGPWithProfile(profile.RFC9580())

// Note that RSA keys should not be generated anymore according to
// RFC9580.

keyGenHandle := pgp4880.KeyGeneration().AddUserId(name, email).New()
// Generates rsa keys with 3072 bits
rsaKey, err := keyGenHandle.GenerateKey()
// Generates rsa keys with 4092 bits
rsaKeyHigh, err := keyGenHandle.GenerateKeyWithSecurity(constants.HighSecurity)

keyGenHandle = pgpDefault.KeyGeneration().AddUserId(name, email).New()
// Generates curve25519 v4 keys.
ecKey, err := keyGenHandle.GenerateKey()

keyGenHandle = pgpCryptoRefresh.KeyGeneration().AddUserId(name, email).New()
// Generates curve25519 v6 keys with RFC9580.
ecKey, err = keyGenHandle.GenerateKey()
// Generates curve448 v6 keys with RFC9580.
ecKeyHigh, err = keyGenHandle.GenerateKeyWithSecurity(constants.HighSecurity)