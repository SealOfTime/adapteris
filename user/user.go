// Package user provides domain model for the User aggregate, including data storage (e.g. Postgres, in-memory)
// and delivery method (e.g. REST, RPC) agnostic implementation of business logic.
package user

import "time"

// User represents the user of AdapterIS
type User struct {
	//Id is a unique auto-incremented identifier of a User
	Id int
	//FullName is a full name of the User (e.g. Vdovitsyn Matvei Valentinovich, Sultan bin Abdulaziz Al Salud...)
	FullName string
	//ShortName is a preferred nickname or a shortened form of the FullName, that is used more commonly.
	ShortName string
	//RegisteredAt is a timestamp at which the User has registered.
	RegisteredAt time.Time
	//IsAdmin shows, whether the User has global administrative rights (i.e. to create a new AdapterSchool etc.)
	IsAdmin bool
}
