package util

import "fmt"

func NewPostgresDS(host, user, password, db string, port int) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", host, user, password, db, port)
}
