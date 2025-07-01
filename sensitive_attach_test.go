package sensitive

import (
	"testing"
)

func TestAttachStructRestored(t *testing.T) {
	t.Parallel()

	type User struct {
		Password string `sensitive:"true"`
	}
	u := User{Password: "Password"}

	detached, sensitive, err := Detach(u)
	assertNoError(t, err)

	attached, err := Attach(detached, sensitive)
	assertNoError(t, err)
	assertDeepEquals(t, "Password", attached.Password, "field should be restored")
}

func TestAttachPointerRestored(t *testing.T) {
	t.Parallel()

	type User struct {
		Password string `sensitive:"true"`
	}
	u := User{Password: "Password"}

	detached, sensitive, err := Detach(&u)
	assertNoError(t, err)

	attached, err := Attach(detached, sensitive)
	assertNoError(t, err)
	assertDeepEquals(t, "Password", attached.Password, "field should be restored")
}

func TestAttachStructRestoredWithModifications(t *testing.T) {
	t.Parallel()

	type User struct {
		Password string `sensitive:"true"`
	}
	u := User{Password: "Password"}

	detached, sensitive, err := Detach(u)
	assertNoError(t, err)

	u.Password = "_changed"

	attached, err := Attach(detached, sensitive)
	assertNoError(t, err)
	assertTrue(t, attached.Password == "Password", "sensitive field should be restored")
}

func TestAttachPointerRestoredWithModifications(t *testing.T) {
	t.Parallel()

	type User struct {
		Password string `sensitive:"true"`
	}
	u := User{Password: "Password"}

	detached, sensitive, err := Detach(&u)
	assertNoError(t, err)

	u.Password = "_changed"

	attached, err := Attach(detached, sensitive)
	assertNoError(t, err)
	assertTrue(t, attached.Password == "Password", "sensitive field should be restored")
}
