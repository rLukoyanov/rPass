package commands

import (
	"fmt"
	"rp/internal/storage"

	"github.com/urfave/cli/v2"
)

func SetPassword(c *cli.Context) error {
	s, err := storage.ReadStorage()
	if err != nil {
		return err
	}

	key := c.String("key")
	value := c.String("value")

	for i, entry := range s.Entries {
		if entry.Key == key {
			s.Entries[i].Value = value
			return storage.WriteStorage(s)
		}
	}

	s.Entries = append(s.Entries, storage.PasswordEntry{
		Key:   key,
		Value: value,
	})

	return storage.WriteStorage(s)
}

func GetPassword(c *cli.Context) error {
	storage, err := storage.ReadStorage()
	if err != nil {
		return err
	}

	key := c.String("key")

	for _, entry := range storage.Entries {
		if entry.Key == key {
			fmt.Printf("Пароль для %s: %s\n", key, entry.Value)
			return nil
		}
	}

	return fmt.Errorf("пароль под ключ '%s' не найден", key)
}

func ListPasswords(c *cli.Context) error {
	storage, err := storage.ReadStorage()
	if err != nil {
		return err
	}

	if len(storage.Entries) == 0 {
		fmt.Println("Нет сохраненных паролей")
		return nil
	}

	fmt.Println("Сохранненные пароли:")
	for _, entry := range storage.Entries {
		fmt.Printf("- %s\n", entry.Key)
	}

	return nil
}

func DeletePassword(c *cli.Context) error {
	s, err := storage.ReadStorage()
	if err != nil {
		return err
	}

	key := c.String("key")
	found := false

	for i, entry := range s.Entries {
		if entry.Key == key {
			s.Entries = append(s.Entries[:i], s.Entries[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("пароль под ключ '%s' не найден", key)
	}

	return storage.WriteStorage(s)
}
