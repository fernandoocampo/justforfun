package interfunc

import "errors"

type Healer interface {
	Heal() error
}

type HealerFunc func() error

func (h HealerFunc) Heal() error {
	return h()
}

func HelpSomeone(healer Healer) error {
	err := healer.Heal()
	if err != nil {
		return errors.New("sorry we can't")
	}

	return nil
}
