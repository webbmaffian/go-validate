package utils

type Flag string
type Flags []Flag

func (f Flags) Has(flag Flag) bool {
	for _, fl := range f {
		if fl == flag {
			return true
		}
	}
	return false
}
