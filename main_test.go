package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunTamanho(t *testing.T) {
	var out, errb bytes.Buffer
	code := run([]string{"-n", "12", "-c", "2"}, &out, &errb)
	if code != 0 {
		t.Fatalf("código %d, esperado 0", code)
	}
	linhas := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(linhas) != 2 || len(linhas[0]) != 12 {
		t.Fatalf("saída inesperada %q", out.String())
	}
}
