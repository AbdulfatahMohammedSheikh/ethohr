package employermodal

import (
	"errors"
	"fmt"
	"reflect"
	"slices"

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	tagmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/tagModal"
	log "github.com/sirupsen/logrus"
)

type Employer struct {
	Id     string `json:"id,omitempty"`
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Meto   string `json:"meto"`
	About  string `json:"about"`
	// TODO: change [Location] to be a struct
	Location string   `json:"location"`
	Phone    string   `json:"phone"`
	Tags     []string `json:"tags"`
}

func New(id, name, meto, about, location, phone string) Employer {
	return Employer{
		UserId:   id,
		Name:     name,
		Meto:     meto,
		About:    about,
		Location: location,
		Phone:    phone,
	}
}

// create

func (e *Employer) Create(a *surreal.AppRepository) error {

	// TODO: make sure that name is unique
	_, err := a.Db.Create("employers", e)

	if nil != err {
		log.Error(err.Error())
		return err
	}

	return nil
}

// index
func Index(a *surreal.AppRepository, id string) (Employer, error) {
	query := "select * from employers where id = $id limit 1"

	m := map[string]interface{}{
		"id": id,
	}
	employer, err := surreal.Find[Employer](a, query, m, []Employer{})

	if 0 == len(employer) {

		log.Errorf("no users with id: %v", id)
		return employer[0], errors.New("check the information you entered")
	}
	return employer[0], err
}

// all
func All(a *surreal.AppRepository) ([]Employer, error) {
	query := "select * from employers "

	m := map[string]interface{}{}
	employers, err := surreal.Find[Employer](a, query, m, []Employer{})

	if 0 == len(employers) {

		log.Errorf("no employers in database")
		return employers, errors.New("check the information you entered")
	}
	return employers, err
}

// put
// TODO: make put method
func Update(a *surreal.AppRepository, e Employer) error {

	if "" == e.Id || len(e.Id) == 0 {
		return errors.New("did not pass correct id")
	}

	// TODO: add the return object so later can be use to check for equality
	employer, err := Index(a, e.Id)
	if nil != err {
		return err
	}

	e.Tags = employer.Tags

	if reflect.DeepEqual(employer, e) {
		return nil
	}

	err = surreal.Update(a, e.Id, e)
	return nil
}

// delete

func Delete(a *surreal.AppRepository, id string) error {
	// TODO: add a surrealql event to manage deltetion

	_, err := Index(a, id)

	if nil != err {
		return err
	}

	err = surreal.Delete(a, id)

	if nil != err {
		return err
	}
	return nil
}

// TODO: check this using endpoints
func AddTag(a *surreal.AppRepository, id, tag string) error {
	employer, err := Index(a, id)

	if nil != err {
		return err
	}

	_, err = tagmodal.Index(a, tag)

	if nil != err {
		return err
	}
	// employers
	if slices.Contains(employer.Tags, tag) {
		return errors.New("employer already has given tag")
	}

	q := fmt.Sprintf("update employers set tags += %s where id = %s", tag, employer.Id)
	_, err = a.Db.Query(q, nil)

	if nil != err {
		return err
	}
	return nil

}

// TODO: add and remove tag
func RemoveTag(a *surreal.AppRepository, id, tag string) error {

	e, err := Index(a, id)

	if nil != err {
		return err
	}

	if !slices.Contains(e.Tags, tag) {
		return errors.New("employer has not tag with given id: " + tag)

	}
	q := fmt.Sprintf("update employers set tags -= %s where id = %s", tag, e.Id)
	_, err = a.Db.Query(q, nil)

	if nil != err {
		return err
	}
	return nil
}
