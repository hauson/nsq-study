package examples

import (
	"fmt"
	"os"
	"testing"
)

func TestAA(t *testing.T) {
	hostname, err := os.Hostname()
	fmt.Println(hostname, err)
}
