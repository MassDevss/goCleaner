package helpers

import "log"

func ManageErr(toManageError error) {
	if toManageError != nil {
		log.Println("Error exist..", toManageError)
	}
}
