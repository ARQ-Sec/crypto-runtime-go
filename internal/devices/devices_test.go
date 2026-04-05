package devices
import "testing"
func TestBuildSmoke(t *testing.T) { if len(Load()) == 0 { t.Fatal("expected seed") } }
