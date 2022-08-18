package integration_tests

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestDummy(t *testing.T) {
	fmt.Println("hello it", uuid.UUID{})
}
