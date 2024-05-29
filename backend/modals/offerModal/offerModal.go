package offermodal

import (
	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
)

type Offer struct {
	Id           string `json:"id,omitempty"`
	EmployerId   string `json:"employer_id"`
	EmployerName string `json:"employer_name"`
	Title        string `json:"title"`
	Requirements string `json:"requirements"`
	// path of the offer, this one depends of editor js
	Duty     string    `json:"duty"`
	PostDate string    `json:"postDate"`
	Deadline string    `json:"deadline"`
	Tags     *[]string `json:"tags"`
}

// TODO: add and remove tags

// TODO: create migration
func New(employerId, employerName, title, requirements, duty, postdate, deadline string) *Offer {
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

func Index(a *surreal.AppRepository, id string) (*[]Offer, error) {
	query := "select * from offers where employer_id = $id "
	m := map[string]interface{}{
		"id": id,
	}
	offers, err := surreal.Find(a, query, m, []Offer{})

	if nil != err {
		return nil, err
	}

	return &offers, nil
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
