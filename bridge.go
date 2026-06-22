package rustlib

/*
#cgo LDFLAGS: -L./target/release -lrustlib
#include <stdlib.h>
*/
import "C"

func Process(input string) (string, error) {
    return "", nil
}
