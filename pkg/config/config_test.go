package config

import (
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestBuildHome(t *testing.T) {
	buildHome()
	t.Log(MinionId)

	t.Log(uuid.NewV4().String())
}
