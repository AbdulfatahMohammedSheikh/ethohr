package offermodal

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	tagmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/tagModal"
)

type Offer struct {
	Id           string   `json:"id,omitempty"`
	EmployerId   string   `json:"employer_id"`
	EmployerName string   `json:"employer_name"`
	Title        string   `json:"title"`
	Requirements []string `json:"requirements"`
	Description  []string `json:"description"`
	// path of the offer, this one depends of editor js
	Duty     []string  `json:"duty"`
	PostDate string    `json:"postDate"`
	Deadline string    `json:"deadline"`
	Tags     *[]string `json:"tags"`
}

// TODO: create migration
func New(employerId, employerName, title, postdate, deadline string, requirements, duty []string) *Offer {
	return &Offer{
		EmployerId:   employerId,
		EmployerName: employerName,
		Title:        title,
		Requirements: requirements,
		Duty:         duty,
		PostDate:     postdate,
		Deadline:     deadline,
		Tags:         nil,
	}
}

func (o *Offer) Create(a *surreal.AppRepository) error {
	// TODO: check if the use is ngo or not
	_, err := a.Db.Create("offers", o)
	if nil != err {
		return err
	}
	return nil
}

func All(a *surreal.AppRepository) (*[]Offer, error) {

	q := "select * from offers"
	m := map[string]interface{}{}

	offers, err := surreal.Find(a, q, m, []Offer{})

	if nil != err {
		return nil, err
	}

	return &offers, nil
}

func Index(a *surreal.AppRepository, id string) (*Offer, error) {
	query := "select * from offers where id = $id "
	m := map[string]interface{}{
		"id": id,
	}
	offers, err := surreal.Find(a, query, m, []Offer{})

	if nil != err {
		return nil, err
	}

	if len(offers) == 0 {
		err = errors.New("no offers yet")
		return nil, err

	}

	return &offers[0], nil
}

func ShowOffer(a *surreal.AppRepository, id string) (*Offer, error) {

	query := "select * from offers where id = $id "
	m := map[string]interface{}{
		"id": id,
	}
	offers, err := surreal.Find(a, query, m, []Offer{})

	if nil != err {
		return nil, err
	}

	return &offers[0], nil
}

func Update(a *surreal.AppRepository, offer Offer) error {
	o, err := Index(a, offer.Id)

	if nil != err {
		return err
	}

	if reflect.DeepEqual(*o, offer) {
		return nil
	}

	err = surreal.Update(a, o.Id, offer)

	if nil != err {
		return err
	}

	return nil
}

// from -> the person who requested the offer to be deleted
// id -> offer id
func Delete(a *surreal.AppRepository, id, from string) error {

	query := "select * from offers where id = $id "
	m := map[string]interface{}{
		"id": id,
	}
	offer, err := surreal.Find(a, query, m, []Offer{})

	if nil != err {
		return err
	}

	if offer[0].EmployerId != from {
		return errors.New("unautherized request")
	}

	err = surreal.Delete(a, id)
	return err

}

func AddTag(a *surreal.AppRepository, offerId, tagId string) error {

	_, err := Index(a, offerId)

	if nil != err {
		return err
	}

	_, err = tagmodal.Index(a, tagId)
	if nil != err {
		return err
	}
	// tag

	q := fmt.Sprintf("update offers set tags += %s where id = %s ;", tagId, offerId)
	_, err = a.Db.Query(q, nil)

	if nil != err {
		return err
	}

	return nil
}

func RemoveTag(a *surreal.AppRepository, offerId, tagId string) error {

	_, err := Index(a, offerId)

	if nil != err {
		return err
	}

	_, err = tagmodal.Index(a, tagId)
	if nil != err {
		return err
	}

	q := fmt.Sprintf("update offers set tags -= %s where id = %s ;", tagId, offerId)
	_, err = a.Db.Query(q, nil)

	if nil != err {
		return err
	}

	return nil
}
