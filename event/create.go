package event

import "os"

// Create File
func Create(name string) (*os.File, error) {
	f, err := os.Create("sample.t30")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f, nil
}
