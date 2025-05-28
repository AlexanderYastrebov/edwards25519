package edwards25519

import (
	"testing"

	"filippo.io/edwards25519/field"
)

func TestCountBaseMult(t *testing.T) {
	var p Point

	// run once to init tables
	p.ScalarBaseMult(dalekScalar)

	before := field.Multiplications

	p.ScalarBaseMult(dalekScalar)
	_ = p.BytesMontgomery()

	after := field.Multiplications

	t.Logf("Multiplications: %d", after-before)
}

func TestCountIncrement(t *testing.T) {
	var p Point

	before := field.Multiplications

	p.Add(B, B)
	_ = p.BytesMontgomery()

	after := field.Multiplications

	t.Logf("Multiplications: %d", after-before)
}
