package event

import "os"

// Create File
func Create(name string) error {
	f, err := os.Create("sample.ta")
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
