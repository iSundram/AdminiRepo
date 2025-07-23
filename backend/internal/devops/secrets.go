
package devops

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
)

type SecretsManager struct {
    gcm cipher.AEAD
}

func NewSecretsManager(key []byte) (*SecretsManager, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    return &SecretsManager{gcm: gcm}, nil
}

func (sm *SecretsManager) Encrypt(data []byte) ([]byte, error) {
    nonce := make([]byte, sm.gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }
    
    ciphertext := sm.gcm.Seal(nonce, nonce, data, nil)
    return ciphertext, nil
}

func (sm *SecretsManager) Decrypt(data []byte) ([]byte, error) {
    nonceSize := sm.gcm.NonceSize()
    nonce, ciphertext := data[:nonceSize], data[nonceSize:]
    
    return sm.gcm.Open(nil, nonce, ciphertext, nil)
}
