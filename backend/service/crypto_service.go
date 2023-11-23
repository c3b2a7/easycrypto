package service

import (
	"context"
	extismsdk "github.com/c3b2a7/HelloGo/thirdparty/extism-sdk-go"
	"sync"
)

var cryptoService *CryptoService
var onceCrypto sync.Once

func GetCryptoService() *CryptoService {
	onceCrypto.Do(func() {
		if cryptoService == nil {
			cryptoService = &CryptoService{
				extismsdk.Must(extismsdk.NewClient(context.Background())),
			}
		}
	})

	return cryptoService
}

type CryptoService struct {
	*extismsdk.Client
}

func (s CryptoService) GetKeys() ([]string, error) {
	return s.KeysAPI.GetKeys(context.Background())
}
