package security
import "testing"
func TestInventoryOnly(t *testing.T) { left, right := Digest([]byte("seed")); if len(left) == 0 || len(right) == 0 { t.Fatal("expected digest output") } }
