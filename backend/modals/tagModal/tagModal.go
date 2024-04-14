package tagmodal

import (
	"errors"
	"net/url"

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
)

type Tag struct {
	// TODO: add more fields to store the count of times that a tag was mentions by ngo
	// TODO: add more feilds to store the count of times that a tag was mentions by users
	Id   string `json:"id,omitempty"`
	Name string `json:"name"`
}

func MapIt(id, name string) string {
	if id == "" {
		return url.Values{
			"name": {name},
		}.Encode()
	}
	return url.Values{
		"id":   {id},
		"name": {name},
	}.Encode()

}

func HasTagWithName(a *surreal.AppRepository, name string) error {

	q := "select * from tags where name = $name limit 1;"
	m := map[string]interface{}{
		"name": name,
	}
	tags, err := surreal.Find[Tag](a, q, m, []Tag{})
	if nil != err {
		return err
	}

	if 0 == len(tags) {
		return nil
	}

	if 0 != len(tags) {
		return errors.New("element already in db")
	}
	return nil
}

// create
func Create(a *surreal.AppRepository, name string) error {
	var tag Tag = Tag{
		Name: name,
	}

	_, err := a.Db.Create("tags", tag)

	if nil != err {
		return err
	}

	return nil
}

func All(a *surreal.AppRepository) ([]Tag, error) {

	q := "select * from tags ;"
	m := map[string]interface{}{}
	tags, err := surreal.Find[Tag](a, q, m, []Tag{})
	if nil != err {
		return []Tag{}, err
	}
	if 0 == len(tags) {
		return []Tag{}, errors.New("no tags has been created yet")
	}

	return tags, nil
}

// index
func Index(a *surreal.AppRepository, id string) (Tag, error) {
	q := "select * from tags where id = $id limit 1"
	m := map[string]interface{}{
		"id": id,
	}
	tag, err := surreal.Find[Tag](a, q, m, []Tag{})
	if nil != err {
		return Tag{}, err
	}
	if 0 == len(tag) || tag[0].Id == "" {
		return Tag{}, errors.New("no tag with given id")
	}

	return tag[0], nil
}

// update
func Update(a *surreal.AppRepository, t Tag) error {

	tag, err := Index(a, t.Id)
	if nil != err {
		return err
	}

	if t == tag {
		return errors.New("there is nothing to update")
	}

	err = surreal.Update(a, t.Id, t)

	if nil != err {
		return err
	}

	return nil
}

// delete
func Delete(a *surreal.AppRepository, id string) error {
	_, err := Index(a, id)
	if nil != err {
		return err
	}

	err = surreal.Delete(a, id)

	return err
}
