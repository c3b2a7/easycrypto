package service

import (
	"errors"
	"github.com/c3b2a7/decrypotr/backend/jstype"
	"github.com/c3b2a7/decrypotr/backend/util"
	"sync"
)

var cryptoService *CryptoService
var onceCrypto sync.Once

var (
	ErrInvalidEnv = errors.New("invalid env")
	cryptoMap     = map[string]util.Crypto{
		"browser-dev": util.NewAesCrypto("CAMG8wU1SwV9PoAJUsPsYV8TdiN+5LDxK7v2qwkATY8="),
		"browser-prd": util.NewAesCrypto("f2h3tGKAUzzbOEBZFt5n2bihFfOs4KwfX41As7tawTE="),
		"browser-db":  util.NewPrefixCrypto("E_0_0_", util.NewAesCrypto("V+MZQ0PyISJyL0Webmow5WB58yR6L6SmhbCl3KcGVuY=")),
		"smy":         util.NewPrefixCrypto("E_0000_0_", util.NewAesCrypto("UnpKaEkgieup86ixZkJStQ==")),
	}
)

func GetCryptoService() *CryptoService {
	onceCrypto.Do(func() {
		if cryptoService == nil {
			cryptoService = &CryptoService{
				crypto: cryptoMap,
			}
		}
	})
	return cryptoService
}

type CryptoService struct {
	crypto map[string]util.Crypto
	mu     sync.Mutex
}

func (c *CryptoService) Decrypt(req jstype.CryptoRequest) (resp jstype.JSResp) {
	crypto := c.getCrypto(req.Env)
	if crypto == nil {
		return jstype.Error(ErrInvalidEnv)
	}
	plaintext, err := crypto.Decrypt(req.Data)
	if err != nil {
		return jstype.Error(err)
	}
	return jstype.Success(plaintext)
}

func (c *CryptoService) Encrypt(req jstype.CryptoRequest) (resp jstype.JSResp) {
	crypto := c.getCrypto(req.Env)
	if crypto == nil {
		return jstype.Error(ErrInvalidEnv)
	}
	ciphertext, err := crypto.Encrypt(req.Data)
	if err != nil {
		return jstype.Error(err)
	}
	return jstype.Success(ciphertext)
}

func (c *CryptoService) getCrypto(env string) util.Crypto {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.crypto[env]
}
