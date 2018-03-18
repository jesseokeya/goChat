package junction

import (
	"strings"

	"gopkg.in/mgo.v2"

	"github.com/jesseokeya/goChat/database"
	"github.com/jesseokeya/goChat/models"
)

var (
	api API
)

// API creates an api struct which connects logics with the database
type API struct {
	db      database.DB
	session *mgo.Session
}

func init() {
	api.session = api.db.ConnectToDatabase()
}

// CreateUser creates a new user in the database
func (a API) CreateUser(u models.User) {
	dbName := strings.ToLower(api.db.GetDbName())
	a.session.DB(dbName).C("Users").Insert(u)
}
