package database

import (
	"encoding/json"
	"errors"
	"os"
)

type Client struct {
	dbPath string
}

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

func NewClient(path string) Client {
	return Client{
		path,
	}
}

func (c Client) EnsureDB() error {
	_, err := os.ReadFile(c.dbPath)
	if errors.Is(err, os.ErrNotExist) {
		return c.createDB()
	}
	return err
}

func (c Client) createDB() error {
	dat, err := json.Marshal(databaseSchema{
		Users: make(map[string]User),
		Posts: make(map[string]Post),
	})
	if err != nil {
		return err
	}
	err = os.WriteFile(c.dbPath, dat, 0600)
	if err != nil {
		return err
	}
	return nil
}
