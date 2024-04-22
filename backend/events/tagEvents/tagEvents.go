package tagevents

import (
	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
)

func OnTagCreated(a *surreal.AppRepository) {

	name := "tag_created"
	event := " \"CREATE\"  "
	action := "update tags set  created_at = time::now() where id = $after.id   "

	err := surreal.CreateEvent(a, name, "tags", event, action)

	if nil != err {
		panic(err.Error())
	}

}

// when tag deleted from tags
func OnTagDeleted(a *surreal.AppRepository) {

	name := "tag_deleted "
	event := " \"DELETE\"  "

	action := "update users set tags -= $before.id "

	err := surreal.CreateEvent(a, name, "tags", event, action)

	if nil != err {
		panic(err.Error())
	}

}
