package service

import "os"

// Write File
func Write(f *os.File, value []byte) ([]byte, error) {
	_, err := f.WriteString(string(value))
	if err != nil {
		return value, err
	}
	defer f.Close()
	return value, nil
}
