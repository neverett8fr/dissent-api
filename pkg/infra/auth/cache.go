package auth

import "fmt"

type TokenCache map[string]struct{}

func (t TokenCache) CacheToken(tok string) {
	t[tok] = struct{}{}
}

func (t TokenCache) Exists(tok string) error {
	if _, ok := t[tok]; ok {
		return nil
	}

	return fmt.Errorf("error token not found in cache, please re-authenticate")
}

func NewTokenCache() TokenCache {
	return TokenCache{}
}
