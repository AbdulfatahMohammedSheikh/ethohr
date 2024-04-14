package surreal

import (
	"encoding/json"
	"fmt"

	"github.com/AbdulfatahMohammedSheikh/backend/core/config"
	logger "github.com/sirupsen/logrus"
	"github.com/surrealdb/surrealdb.go"
	surreal "github.com/surrealdb/surrealdb.go"
)

type App struct {
	Version string
	Mode    string
	DB      DB
}

type DB struct {
	Address   string
	User      string
	Pass      string
	Namespace string
	Database  string
}

type ConfigRepository struct {
	Address   string
	Pass      string
	User      string
	Database  string
	Namespace string
}

type AppRepository struct {
	Db *surreal.DB
}

type Result[T interface{}] struct {
	Data   []T    `json:"result"`
	Status string `json:"status"`
	Time   string `json:"time"`
}

var log = logger.New()

func NewApp() *App {

	mode := config.GetConigVirable("mode")
	var app *App
	var db DB

	if "dev" == mode {

		db = DB{
			Address:   config.GetConigVirable("address"),
			User:      config.GetConigVirable("username"),
			Pass:      config.GetConigVirable("pass"),
			Namespace: config.GetConigVirable("namespace"),
			Database:  config.GetConigVirable("database"),
		}
		app = &App{
			Mode:    config.GetConigVirable("mode"),
			Version: config.GetConigVirable("version"),
			DB:      db,
		}

		return app
	}

	db = DB{
		Address:   config.GetConigVirable("t_address"),
		User:      config.GetConigVirable("t_username"),
		Pass:      config.GetConigVirable("t_pass"),
		Namespace: config.GetConigVirable("t_namespace"),
		Database:  config.GetConigVirable("t_database"),
	}
	app = &App{
		Mode:    config.GetConigVirable("mode"),
		Version: config.GetConigVirable("version"),
		DB:      db,
	}
	return app

}

func NewAppRepository(config DB) (*AppRepository, error) {

	db, err := surreal.New(config.Address)

	if nil != err {

		return nil, fmt.Errorf("Faid to connect to address : %s", err)
	}

	_, err = db.Signin(map[string]interface{}{

		"user": config.User,
		"pass": config.Pass,
	})

	if nil != err {
		return nil, fmt.Errorf("Faid to signin : %s", err)
	}

	_, err = db.Use(config.Database, config.Namespace)

	if nil != err {
		return nil, err
	}

	return &AppRepository{Db: db}, nil
}

// TODO: check if this [Close] is actually called
func (a AppRepository) Close() {
	a.Db.Close()
}

func Decode[T interface{}](data interface{}) ([]T, error) {

	jsonBytes, err := json.Marshal(data)
	if nil != err {
		return []T{}, err
	}

	var r []Result[T]

	err = json.Unmarshal(jsonBytes, &r)

	if nil != err {
		return []T{}, err
	}

	return r[0].Data, nil
}

func Find[T interface{}](a *AppRepository, q string, m map[string]interface{}, input []T) ([]T, error) {

	data, err := a.Db.Query(q, m)
	if nil != err {
		logger.Fatal(err)
		return input, err
	}

	r, err := Decode[T](data)

	if nil != err {
		return input, err
	}
	return r, nil
}

func All[T interface{}](a *AppRepository, table string) ([]T, error) {

	data, err := a.Db.Select(table)

	if nil != err {
		return []T{}, err
	}

	result := []T{}

	err = surrealdb.Unmarshal(data, &result)
	if nil != err {
		return []T{}, err
	}

	return result, nil
}

func Delete(a *AppRepository, id string) error {
	_, err := a.Db.Delete(id)

	if nil != err {
		logger.Fatal("error ", err)
		return err
	}

	// TODO: log the name of the deleted tag
	logger.Warn("deleted item: ", id)
	return err
}

func Update(a *AppRepository, id string, recored interface{}) error {
	// TODO: make this more appealing this mathod does not check if the recored exitst. It just calls [a.Db.Update] and hopes things work well.
	// instead its it better to check if the record in db else return custom error
	_, err := a.Db.Update(id, recored)
	return err
}
