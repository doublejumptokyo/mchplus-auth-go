package signer

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

type Signer struct {
	privKey *ecdsa.PrivateKey
}

func NewSignerFromHex(hex string) (*Signer, error) {
	b, err := decodeHex(hex)
	if err != nil {
		return nil, err
	}
	k, err := crypto.ToECDSA(b)
	if err != nil {
		return nil, err
	}
	return &Signer{
		privKey: k,
	}, nil
}

func (s *Signer) Public() ecdsa.PublicKey {
	return *s.privKey.Public().(*ecdsa.PublicKey)
}

func (s *Signer) Address() string {
	return crypto.PubkeyToAddress(s.Public()).String()
}

func (s *Signer) EthereumSign(msg []byte) ([]byte, error) {
	h := toEthSignedMessageHash(msg)

	sig, err := crypto.Sign(h, s.privKey)
	if err != nil {
		return nil, err
	}
	if sig[64] < 27 {
		sig[64] += 27
	}

	return sig, nil
}

func (s *Signer) PersonalSign(msg string) (string, error) {
	b, err := s.EthereumSign([]byte(msg))
	if err != nil {
		return "", err
	}

	return encodeToHex(b), nil
}

func recover(hash, sig []byte) (common.Address, error) {
	if sig[64] >= 27 {
		sig[64] -= 27
	}
	pub, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return common.HexToAddress("0x0"), err
	}
	return crypto.PubkeyToAddress(*pub), nil
}

func encodeToHex(b []byte) string {
	return "0x" + hex.EncodeToString(b)
}

func decodeHex(s string) ([]byte, error) {
	if s[0:2] != "0x" {
		return nil, errors.New("hex must start with 0x")
	}
	return hex.DecodeString(s[2:])
}

func toEthSignedMessageHash(message []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	return Keccak256([]byte(msg))
}

func Keccak256(data []byte) []byte {
	return crypto.Keccak256(data)
}
