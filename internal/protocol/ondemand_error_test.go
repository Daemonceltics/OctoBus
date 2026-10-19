package protocol

import (
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestParseOctobusErrorFindsStructuredLine(t *testing.T) {
	stderr := "some diagnostic\nOCTOBUS_ERROR:{\"code\":\"INVALID_ARGUMENT\",\"message\":\"bad input\"}\ntrailing output\n"
	code, message, ok := parseOctobusError(stderr)
	if !ok || code != codes.InvalidArgument || message != "bad input" {
		t.Fatalf("code=%v message=%q ok=%v", code, message, ok)
	}
}

func TestOnDemandProcessErrorDoesNotLeakStderr(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell fixture is unix-only")
	}
	cmd := exec.Command("sh", "-c", "exit 3")
	err := cmd.Run()
	if err == nil {
		t.Fatal("expected a failing helper process")
	}
	got := onDemandProcessError(err, "TOKEN=super-secret\nstack trace\n")
	if status.Code(got) != codes.Internal {
		t.Fatalf("code=%v want Internal", status.Code(got))
	}
	if strings.Contains(got.Error(), "super-secret") {
		t.Fatalf("raw stderr leaked to the caller: %v", got)
	}
}
