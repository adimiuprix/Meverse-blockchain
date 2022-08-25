package key

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/meverselabs/meverse/common"
	"github.com/meverselabs/meverse/common/hash"
)

// MemoryKey is the in-memory crypto key
type MemoryKey struct {
	PrivKey *ecdsa.PrivateKey
	pubkey  common.PublicKey
	ChainID *big.Int
}

// NewMemoryKey returns a MemoryKey
func NewMemoryKey(chainID *big.Int) (*MemoryKey, error) {
	PrivKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ac := &MemoryKey{
		PrivKey: PrivKey,
		ChainID: chainID,
	}
	if err := ac.calcPubkey(); err != nil {
		return nil, errors.WithStack(err)
	}
	return ac, nil
}

// NewMemoryKeyFromString parse memory key by the hex string
func NewMemoryKeyFromString(chainID *big.Int, sk string) (*MemoryKey, error) {
	ac := &MemoryKey{
		PrivKey: &ecdsa.PrivateKey{
			PublicKey: ecdsa.PublicKey{
				Curve: crypto.S256(),
			},
			D: new(big.Int),
		},
		ChainID: chainID,
	}
	ac.PrivKey.D.SetString(sk, 16)
	ac.PrivKey.PublicKey.X, ac.PrivKey.PublicKey.Y = ac.PrivKey.Curve.ScalarBaseMult(ac.PrivKey.D.Bytes())
	if err := ac.calcPubkey(); err != nil {
		return nil, errors.WithStack(err)
	}
	return ac, nil
}

// NewMemoryKeyFromBytes parse memory key by the byte array
func NewMemoryKeyFromBytes(chainID *big.Int, pk []byte) (*MemoryKey, error) {
	ac := &MemoryKey{
		PrivKey: &ecdsa.PrivateKey{
			PublicKey: ecdsa.PublicKey{
				Curve: crypto.S256(),
			},
			D: new(big.Int),
		},
		ChainID: chainID,
	}
	ac.PrivKey.D.SetBytes(pk)
	ac.PrivKey.PublicKey.X, ac.PrivKey.PublicKey.Y = ac.PrivKey.Curve.ScalarBaseMult(ac.PrivKey.D.Bytes())
	if err := ac.calcPubkey(); err != nil {
		return nil, errors.WithStack(err)
	}
	return ac, nil
}

// Clear removes private key bytes data
func (ac *MemoryKey) Clear() {
	ac.PrivKey.D.SetBytes([]byte{0})
	ac.PrivKey.X.SetBytes([]byte{0})
	ac.PrivKey.Y.SetBytes([]byte{0})
}

func (ac *MemoryKey) calcPubkey() error {
	bs := elliptic.Marshal(ac.PrivKey.PublicKey.Curve, ac.PrivKey.PublicKey.X, ac.PrivKey.PublicKey.Y)
	copy(ac.pubkey[:], bs[:])
	return nil
}

// PublicKey returns the public key of the private key
func (ac *MemoryKey) PublicKey() common.PublicKey {
	return ac.pubkey
}

// Sign generates the signature of the target hash
func (ac *MemoryKey) Sign(h hash.Hash256) (common.Signature, error) {
	bs, err := crypto.Sign(h[:], ac.PrivKey)
	if err != nil {
		return common.Signature{}, errors.WithStack(err)
	}
	v := big.NewInt(0).SetInt64(int64(bs[64]))

	ChainCap := common.GetChainCap(ac.ChainID)

	c := v.Add(ChainCap, v)

	var sig common.Signature
	sig = append(sig, bs[:64]...)
	sig = append(sig, c.Bytes()...)
	return sig, nil
}

// SignWithPassphrase doesn't implemented yet
func (ac *MemoryKey) SignWithPassphrase(h hash.Hash256, passphrase []byte) (common.Signature, error) {
	panic("not implemented")
}

// Verify checks that the signatures is generated by the hash and the key or not
func (ac *MemoryKey) Verify(h hash.Hash256, sig common.Signature) bool {
	return crypto.VerifySignature(ac.pubkey[:], h[:], sig[:])
}

// Bytes returns the byte array of the key
func (ac *MemoryKey) Bytes() []byte {
	return ac.PrivKey.D.Bytes()
}
