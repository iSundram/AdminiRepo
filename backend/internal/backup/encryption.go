
package backup

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
)

type EncryptionManager struct {
    Key []byte
}

func NewEncryptionManager(key []byte) *EncryptionManager {
    return &EncryptionManager{
        Key: key,
    }
}

func (em *EncryptionManager) Encrypt(data []byte) ([]byte, error) {
    block, err := aes.NewCipher(em.Key)
    if err != nil {
        return nil, err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }
    
    ciphertext := gcm.Seal(nonce, nonce, data, nil)
    return ciphertext, nil
}

func (em *EncryptionManager) Decrypt(data []byte) ([]byte, error) {
    block, err := aes.NewCipher(em.Key)
    if err != nil {
        return nil, err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    nonceSize := gcm.NonceSize()
    nonce, ciphertext := data[:nonceSize], data[nonceSize:]
    
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    return plaintext, err
}
