package rolemodal

import (
	"errors"

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
)

type Role struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name"`
	Isband bool   `json:"isband"`
}

// new cole
func New(name string) *Role {
	return &Role{
		Name:   name,
		Isband: false,
	}
}

// create
func (r *Role) Create(a *surreal.AppRepository) error {

	// TODO: add db constaints to

	_, err := a.Db.Create("roles", r)
	return err

}

func HasRoleWithId(a *surreal.AppRepository, id string) (Role, error) {
	q := "select * from roles where id = $id limit 1;"
	m := map[string]interface{}{
		"id": id,
	}

	role, err := surreal.Find[Role](a, q, m, []Role{})
	if nil != err {
		return Role{}, err
	}

	if 0 == len(role) {
		return Role{}, errors.New("there is no role with given id")
	}

	if "" == role[0].Id {
		return Role{}, errors.New("there is no role with given id")
	}

	return role[0], nil
}

func Update(a *surreal.AppRepository, role Role) error {

	r, err := HasRoleWithId(a, role.Id)

	if nil != err {
		return err
	}

	err = surreal.Update(a, r.Id, r)

	if nil != err {
		return err
	}
	return nil
}

func Band(a *surreal.AppRepository, id string) error {
	role, err := HasRoleWithId(a, id)

	if nil != err {
		return err
	}

	role.Isband = true

	err = Update(a, role)

	if nil != err {
		return err
	}

	return nil
}

func All(a *surreal.AppRepository) []Role {

	roles, _ := surreal.All[Role](a, "roles")
	return roles
}
