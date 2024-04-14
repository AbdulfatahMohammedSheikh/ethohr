package rolemodaltest

import (
	"testing"

	testrunner "github.com/AbdulfatahMohammedSheikh/backend/core/testRunner"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
)

func TestAllRoles(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}
	roles := rolemodal.All(r)

	if len(roles) == 0 {
		t.Fatal("no roles in db")
	}

}

func TestHasRoleWithId(t *testing.T) {

	r, err := testrunner.GetConfig()
	defer r.Close()

	if nil != err {
		t.Fatal(err.Error())
	}
	roles := rolemodal.All(r)

	if len(roles) == 0 {
		t.Fatal("no roles in db")
	}

	for _, role := range roles {
		_, err = rolemodal.HasRoleWithId(r, role.Id)

		if nil != err {
			t.Fatal(err.Error())
		}
	}
}
