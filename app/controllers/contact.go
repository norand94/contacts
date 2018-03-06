package controllers

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/norand94/contacts/app/models"
	"github.com/revel/revel"
)

type Contact struct {
	*revel.Controller
}

func (c Contact) Index() revel.Result {
	var (
		contacts []models.Contact
		err      error
	)
	contacts, err = models.GetContacts()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200
	return c.RenderJSON(contacts)
}

func (c Contact) Show(id string) revel.Result {
	var (
		contact   models.Contact
		err       error
		contactID bson.ObjectId
	)
	c.Log.Info("id = " + id)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid contact id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	contactID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid contact id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	contact, err = models.GetContact(contactID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	c.Response.Status = 200
	return c.RenderJSON(contact)
}

func (c Contact) Create() revel.Result {
	var (
		contact models.Contact
		err     error
	)

	err = c.Params.BindJSON(&contact)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJSON(errResp)
	}

	contact, err = models.AddContact(contact)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 201
	return c.RenderJSON(contact)
}

func (c Contact) Update() revel.Result {
	c.Log.Info("update!!")
	var (
		contact models.Contact
		err     error
	)
	err = c.Params.BindJSON(&contact)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	err = contact.UpdateContact()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	return c.RenderJSON(contact)
}

func (c Contact) Delete(id string) revel.Result {
	var (
		err       error
		contact   models.Contact
		contactID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid contact id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	contactID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid contact id format"), "400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	contact, err = models.GetContact(contactID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	err = contact.DeleteContact()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 204
	return c.RenderJSON(nil)
}
