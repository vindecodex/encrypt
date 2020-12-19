package service

import "os"

// Create File
func Create(name string) (*os.File, error) {
	f, err := os.Create(name + ".enc")
	if err != nil {
		return nil, err
	}
	return f, nil
}
