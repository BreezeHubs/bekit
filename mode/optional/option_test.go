package optional

import (
	"errors"
	"testing"
)

type User struct {
	Name string
}

func TestBindOpt(t *testing.T) {
	u := &User{}
	BindOpt[User](u, SetName("breeze"))
	t.Logf("%+v", u)
}

func TestBindOptWithExcept(t *testing.T) {
	u := &User{}
	if err := BindOptWithExcept[User](u, SetNameWithExcept("")); err != nil {
		t.Errorf("%s", err.Error())
	}
	t.Logf("%+v", u)
}

func SetName(name string) Opt[User] {
	return func(t *User) {
		t.Name = name
	}
}

func SetNameWithExcept(name string) OptWithExcept[User] {
	return func(t *User) error {
		if name == "" {
			return errors.New("name不能为空")
		}
		t.Name = name
		return nil
	}
}
