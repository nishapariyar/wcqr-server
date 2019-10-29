package model

import "github.com/jinzhu/gorm"

type Attendee struct {
	gorm.Model
	AttendeeID         int    `json:"attendee_id"`
	FirstName          string `json:"firstname"`
	LastName           string `json:"lastname"`
	Email              string `json:"email"`
	AttendedEvent      bool   `json:"attended_event"`
	AttendedAfterparty bool   `json:"attended_afterparty"`
}
