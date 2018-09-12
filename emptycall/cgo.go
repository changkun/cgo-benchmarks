package emptycall

/*
void empty() {
}
*/
import "C"

// Cempty is an empty cgo call
func Cempty() {
	C.empty()
}
