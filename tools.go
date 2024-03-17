package toolkit

import "crypto/rand"

const randomStringSource  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+_"

// Tools is the type used to instantiate this module.
// Any varuable of this type will have access to akk the 
// methods with the receiver *Tools
type Tools struct {}

// Random string returns a string of random characters of len n
func (t *Tools) RandomString(n int) string {
	s,r := make([]rune, n), []rune(randomStringSource)
	for i:= range s {
		p, _ := rand.Prime(rand.Reader, len(s))
		x,y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}