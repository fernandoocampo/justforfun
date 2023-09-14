package interfunc_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/interfunc"
)

func TestHealWithStruct(t *testing.T) {
	ahealer := aHealer{}
	err := interfunc.HelpSomeone(ahealer)

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestHealWithFunc(t *testing.T) {
	ahealer := makeHealerFunc()
	err := interfunc.HelpSomeone(ahealer)

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

type aHealer struct {
}

func (h aHealer) Heal() error {
	return nil
}

func makeHealerFunc() interfunc.HealerFunc {
	return func() error {
		return nil
	}
}
