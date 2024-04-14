package offermodal

import (
	"errors"

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
)

type Offer struct {
	Id       string `json:"id,omitempty"`
	Employer string `json:"employer"`
	Title    string `json:"title"`
	// path of the offer, this one depends of editor js
	OfferPath string   `json:"offer"`
	PostDate  string   `json:"postDate"`
	Deadline  string   `json:"deadline"`
	Tags      []string `json:"tags"`
}

// new
func New(employer, title, postDate, deadline string, tags []string) Offer {
	// TODO: add the offerPath by using static files
	return Offer{
		Employer: employer,
		Title:    title,
		PostDate: postDate,
		Deadline: deadline,
		Tags:     tags,
	}
}

// create
func (o *Offer) Create(a *surreal.AppRepository) error {
    // TODO: implement this
	return nil
}

// index
func Index(a *surreal.AppRepository, id string) (Offer, error) {
	q := "select * from offers where id = $id limit 1;"
	m := map[string]interface {
	}{
		"id": id,
	}

	offer, err := surreal.Find[Offer](a, q, m, []Offer{})
	if nil != err {
		return Offer{}, err
	}
	if 0 == len(offer) || offer[0].Id == "" {
		return Offer{}, errors.New("no offer with given id")
	}
	return offer[0], nil
}

// update
// delete
func Delete(a *surreal.AppRepository, id string) error {
	_, err := Index(a, id)
	if nil != err {
		return err
	}

	err = surreal.Delete(a, id)
	return err
}
