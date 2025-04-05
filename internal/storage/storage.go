package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"rp/internal/crypt"

	"github.com/urfave/cli/v2"
)

const storageFileName = ".passman_store"

type PasswordEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PasswordStorage struct {
	Entries []PasswordEntry `json:"entries"`
}

func getStoragePath() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("не удалось определить текущую директорию: %v", err)
	}

	storageDir := filepath.Join(currentDir, "data")

	if err := os.MkdirAll(storageDir, 0700); err != nil {
		return "", fmt.Errorf("не удалось создать папку для данных: %v", err)
	}

	return filepath.Join(storageDir, storageFileName), nil
}
func InitStorage(c *cli.Context) error {
	storagePath, err := getStoragePath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(storagePath); err == nil {
		return errors.New("хранилище уже существует")
	}

	storage := PasswordStorage{Entries: []PasswordEntry{}}
	data, err := json.Marshal(storage)
	if err != nil {
		return err
	}

	key, err := crypt.GetMasterKey()
	if err != nil {
		return err
	}

	encryptedData, err := crypt.Encrypt(data, key)
	if err != nil {
		return err
	}

	return os.WriteFile(storagePath, encryptedData, 0600)
}

func ReadStorage() (PasswordStorage, error) {
	storagePath, err := getStoragePath()
	if err != nil {
		return PasswordStorage{}, err
	}

	encryptedData, err := os.ReadFile(storagePath)
	if err != nil {
		return PasswordStorage{}, err
	}

	key, err := crypt.GetMasterKey()
	if err != nil {
		return PasswordStorage{}, err
	}

	data, err := crypt.Decrypt(encryptedData, key)
	if err != nil {
		return PasswordStorage{}, err
	}

	var storage PasswordStorage
	err = json.Unmarshal(data, &storage)
	return storage, err
}

func WriteStorage(storage PasswordStorage) error {
	storagePath, err := getStoragePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(storage)
	if err != nil {
		return err
	}

	key, err := crypt.GetMasterKey()
	if err != nil {
		return err
	}

	encryptedData, err := crypt.Encrypt(data, key)
	if err != nil {
		return err
	}

	return os.WriteFile(storagePath, encryptedData, 0600)
}
