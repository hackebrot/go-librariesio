package librariesio

// Bool returns a *bool for the given value
func Bool(b bool) *bool {
	return &b
}

// Int returns a *int for the given value
func Int(i int) *int {
	return &i
}

// String returns a *string for the given value
func String(s string) *string {
	return &s
}
