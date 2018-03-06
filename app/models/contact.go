package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/norand94/contacts/app/models/mongodb"
)

type Contact struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Email     string        `json:"email" bson:"email"`
	Phone     string        `json:"phone" bson:"phone"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

func newContactCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("contacts")
}

// AddContact insert a new Contact into database and returns
// last inserted contact on success.
func AddContact(m Contact) (contact Contact, err error) {
	c := newContactCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	return m, c.Session.Insert(m)
}

// UpdateContact update a Contact into database and returns
// last nil on success.
func (m Contact) UpdateContact() error {
	c := newContactCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": m.ID,
	}, bson.M{
		"$set": bson.M{
			"name": m.Name, "email": m.Email, "phone": m.Phone, "updatedAt": time.Now()},
	})
	return err
}

// DeleteContact Delete Contact from database and returns
// last nil on success.
func (m Contact) DeleteContact() error {
	c := newContactCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

// GetContacts Get all Contact from database and returns
// list of Contact on success
func GetContacts() ([]Contact, error) {
	var (
		contacts []Contact
		err      error
	)

	c := newContactCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-createdAt").All(&contacts)
	return contacts, err
}

// GetContact Get a Contact from database and returns
// a Contact on success
func GetContact(id bson.ObjectId) (Contact, error) {
	var (
		contact Contact
		err     error
	)

	c := newContactCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&contact)
	return contact, err
}
