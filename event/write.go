package event

import "os"

// Write File
func Write(f *os.File, value string) ([]byte, error) {
	_, err := f.WriteString(value)
	if err != nil {
		return []byte(value), err
	}
	defer f.Close()
	return []byte(value), nil
}
