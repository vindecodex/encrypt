package event

// Write File
func Write(value string) (bool, error) {
	_, err := f.WriteString(value)
	if err != nil {
		return false, err
	}
	return true, nil
}
