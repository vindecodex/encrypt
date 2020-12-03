package event

import "os"

// Write File
func Write(f *os.File, value string) (string, error) {
	_, err := f.WriteString(value)
	if err != nil {
		return value, err
	}
	return value, nil
}
