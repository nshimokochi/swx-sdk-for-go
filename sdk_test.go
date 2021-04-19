package openapi

import (
	"fmt"
	"context"
	"encoding/json"
	"os"
	"testing"
)

// Get space from env
var space string = os.Getenv("SPACE")

var collName, modelName, thingTitle string = "testcollection1", "testmodel1", "testthing1"

var thingUID string

// thing properties
type properties struct {
	Description string `json:"description,omitempty"`
	ReadOnly    bool   `json:"readOnly,omitempty"`
	Title       string `json:"title,omitempty"`
	Type        string `json:"type,omitempty"`
	Unit        string `json:"unit,omitempty"`
}

type actionInput struct {
	Maximum int    `json:"maximum,omitempty"`
	Minimum int    `json:"minimum,omitempty"`
	Type    string `json:"type,omitempty"`
}
type actionProperties struct {
	Input actionInput `json:"input,omitempty"`
}
type actInput struct {
	Properties actionProperties `json:"properties,omitempty"`
}

// thing actions
type actions struct {
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Input       actInput `json:"input,omitempty"`
}

type eventData struct {
	Type string `json:"type,omitempty"`
	Unit string `json:"unit,omitempty"`
}

// thing events
type events struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Data        eventData `json:"data,omitempty"`
}

// error message
type errorMessage struct {
	Message string `json:"message,omitempty"`
}

// error response
type errorResponse struct {
	Error errorMessage `json:"error,omitempty"`
}

// ****************************************
// Positive Tests for collection APIs
// ****************************************

func TestAddCollection(t *testing.T) {
	collDesc := "My collection"

	// CollectionRequest | Create a new collection in the platform
	cr := *NewCollectionRequest()
	cr.Name = &collName
	cr.Description = &collDesc

	resp, r, err := apiClient.CollectionsApi.AddCollection(
		context.Background(),
		space).CollectionRequest(cr).Execute()
	if err != nil {
		t.Errorf("Error when calling `CollectionsApi.AddCollection``: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `AddCollection`: CollectionResponse
	if r.StatusCode != 201 || *resp.Name != collName || *resp.Description != collDesc {
		t.Errorf("%v failed; expected 201, got %v with name = %v & description = %v", 
			t.Name(), r.StatusCode, *resp.Name, *resp.Description)
	}
}

func TestListCollections(t *testing.T) {
	resp, r, err := apiClient.CollectionsApi.ListCollections(
		context.Background(),
		space).Execute()
	if err != nil {
		t.Errorf("Error when calling `CollectionsApi.ListCollections``: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from response from `ListCollections`: CollectionListResponse
	if r.StatusCode != 200 || len(*resp.Data) == 0 {
		t.Errorf("%v failed; expected 200, got %v with http response body = %v", t.Name(), r.StatusCode, r.Body)
	}
}

func TestUpdateCollection(t *testing.T) {
	collDesc := "My test collection"

	// CollectionUpdateRequest | Update an existent collection by name
	cur := *NewCollectionUpdateRequest()
	cur.Description = &collDesc

	resp, r, err := apiClient.CollectionsApi.UpdateCollection(
		context.Background(),
		space,
		collName).CollectionUpdateRequest(cur).Execute()
	if err != nil {
		t.Errorf("Error when calling `CollectionsApi.UpdateCollection``: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `UpdateCollection`: CollectionUpdateResponse
	if r.StatusCode != 200 || *resp.Name != collName || *resp.Description != collDesc {
		t.Errorf("%v failed; expected 200, got %v with name = %v & description = %v", 
			t.Name(), r.StatusCode, *resp.Name, *resp.Description)
	}
}

func TestShowCollection(t *testing.T) {
	collDesc := "My test collection"

	resp, r, err := apiClient.CollectionsApi.ShowCollection(
		context.Background(),
		space, collName).Execute()
	if err != nil {
		t.Errorf("Error when calling `CollectionsApi.ShowCollection``: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `ShowCollection`: CollectionResponse
	if r.StatusCode != 200 || *resp.Name != collName || *resp.Description != collDesc {
		t.Errorf("%v failed; expected 200, got %v with name = %v & description = %v", 
			t.Name(), r.StatusCode, *resp.Name, *resp.Description)
	}
}

// ****************************************
// Negative Tests for collection APIs
// ****************************************

func TestListCollectionSpaceNotFound(t *testing.T) {
	_, r, _ := apiClient.CollectionsApi.ListCollections(
		context.Background(),
		"space_not_found_1").Execute()

	// Check response from `ListCollections`: CollectionListResponse
	er := errorResponse{}
	json.NewDecoder(r.Body).Decode(&er)
	if r.StatusCode == 403 || r.StatusCode == 404 {
		if er.Error.Message == "" {
			t.Errorf("%v failed: received empty error message", t.Name())
		}
	} else {
		t.Errorf("%v failed: expected 403 or 404, got %v with error message = %v", 
			t.Name(), r.StatusCode, er.Error.Message)
	}
}

func TestAddCollectionAlreadyExists(t *testing.T) {
	collDesc := "My collection"

	// CollectionRequest | Create a new collection in the platform
	cr := *NewCollectionRequest()
	cr.Name = &collName
	cr.Description = &collDesc

	_, r, _ := apiClient.CollectionsApi.AddCollection(
		context.Background(),
		space).CollectionRequest(cr).Execute()

	// Check response from `AddCollection`: CollectionResponse
	er := errorResponse{}
	json.NewDecoder(r.Body).Decode(&er)
	if r.StatusCode != 400 || er.Error.Message == "" {
		t.Errorf("%v failed: expected 400 with non-empty error message, got %v with error message = %v", 
			t.Name(), r.StatusCode, er.Error.Message)
	}
}

func TestUpdateCollectionNotFound(t *testing.T) {
	collDesc := "My test collection"

	// CollectionUpdateRequest | Update an existent collection by name
	cur := *NewCollectionUpdateRequest()
	cur.Description = &collDesc

	_, r, _ := apiClient.CollectionsApi.UpdateCollection(
		context.Background(),
		space,
		"collection_not_found_1").CollectionUpdateRequest(cur).Execute()

	// Check response from `UpdateCollection`: CollectionUpdateResponse
	if r.StatusCode != 404 {
		t.Errorf("%v failed: got %v, expected 404", t.Name(), r.StatusCode)
		t.Errorf("HTTP response body: %v\n", r.Body)
	}
}

// ****************************************
// Positive Tests for model APIs
// ****************************************

func TestAddModel(t *testing.T) {
	modelDesc := "My model"

	// ModelRequest | Create a new collection in the platform
	mr := map[string]interface{}{
		"name":        modelName,
		"description": modelDesc,
	}

	resp, r, err := apiClient.ModelsApi.AddModel(
		context.Background(),
		space,
		collName).RequestBody(mr).Execute()
	if err != nil {
		t.Errorf("Error when calling `ModelsApi.AddModel``: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `AddModel`: ModelResponse
	if r.StatusCode != 201 || *resp.Collection != collName || *resp.Name != modelName || *resp.Description != modelDesc {
		t.Errorf("%v failed; expected 201, got %v with collection = %v, modelName = %v & modelDesc = %v",
			t.Name(), r.StatusCode, *resp.Collection, *resp.Name, *resp.Description)
	}
}

func TestListModels(t *testing.T) {
	resp, r, err := apiClient.ModelsApi.ListModels(
		context.Background(),
		space,
		collName).Execute()
	if err != nil {
		t.Errorf("Error when calling `ModelsApi.ListModels``: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `ListModels`: ModelListResponse
	if r.StatusCode != 200 || len(*resp.Data) == 0 {
		t.Errorf("%v failed; expected 200, got %v with http response body = %v", t.Name(), r.StatusCode, r.Body)
	}
}

func TestUpdateModel(t *testing.T) {
	modelDesc := "My test model"

	// CollectionUpdateRequest | Update an existent collection by name
	mur := map[string]interface{}{
		"name":        modelName,
		"description": modelDesc,
	}

	resp, r, err := apiClient.ModelsApi.UpdateModel(
		context.Background(),
		space,
		collName,
		modelName).RequestBody(mur).Execute()
	if err != nil {
		t.Errorf("rror when calling `ModelsApi.UpdateModel``: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `UpdateModel`: ModelUpdateResponse
	if r.StatusCode != 200 || *resp.Collection != collName || *resp.Name != modelName || *resp.Description != modelDesc {
		t.Errorf("%v failed; expected 200, got %v with collection = %v, modelName = %v & modelDesc = %v",
			t.Name(), r.StatusCode, *resp.Collection, *resp.Name, *resp.Description)
	}
}

func TestShowModel(t *testing.T) {
	modelDesc := "My test model"

	resp, r, err := apiClient.ModelsApi.ShowModel(
		context.Background(),
		space,
		collName,
		modelName).Execute()
	if err != nil {
		t.Errorf("Error when calling `ModelsApi.ShowModel``: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `ShowModel`: ModelResponse
	if r.StatusCode != 200 || *resp.Collection != collName || *resp.Name != modelName || *resp.Description != modelDesc {
		t.Errorf("%v failed; expected 200, got %v with collection = %v, modelName = %v & modelDesc = %v",
			t.Name(), r.StatusCode, *resp.Collection, *resp.Name, *resp.Description)
	}
}

// ****************************************
// Negative Tests for model APIs
// ****************************************

func TestAddModelAlreadyExists(t *testing.T) {
	modelDesc := "My model"

	// ModelRequest | Create a new collection in the platform
	mr := map[string]interface{}{
		"name":        modelName,
		"description": modelDesc,
	}

	_, r, _ := apiClient.ModelsApi.AddModel(
		context.Background(),
		space,
		collName).RequestBody(mr).Execute()

	// Check response from `AddModel`: ModelResponse
	er := errorResponse{}
	json.NewDecoder(r.Body).Decode(&er)
	if r.StatusCode != 400 || er.Error.Message == "" {
		t.Errorf("%v failed: expected 400 with non-empty error message, got %v with error message = %v", 
			t.Name(), r.StatusCode, er.Error.Message)
	}
}

func TestUpdateModelNotFound(t *testing.T) {
	modelDesc := "My test model"

	// CollectionUpdateRequest | Update an existent collection by name
	mur := map[string]interface{}{
		"name":        modelName,
		"description": modelDesc,
	}

	_, r, _ := apiClient.ModelsApi.UpdateModel(
		context.Background(),
		space,
		collName,
		"model_not_found_1").RequestBody(mur).Execute()

	// Check response from `UpdateModel`: ModelUpdateResponse
	if r.StatusCode != 404 {
		t.Errorf("%v failed: got %v, expected 404", t.Name(), r.StatusCode)
		t.Errorf("HTTP response body: %v\n", r.Body)
	}
}

// ****************************************
// Positive Tests for thing APIs
// ****************************************

func TestAddThing(t *testing.T) {
	thingDesc := "My thing"

	// Create thing request with property
	thingProperty1 := properties{
		Description: "This is Test Property 1",
		ReadOnly:    false,
		Title:       "Test Property 1",
		Unit:        "percent",
		Type:        "number",
	}
	thingProperties := map[string]interface{}{
		"property1": thingProperty1,
	}

	// ThingRequest | Create a new thing description in the platform
	tr := *NewThingRequest()
	tr.Title = &thingTitle
	tr.Description = &thingDesc
	tr.Properties = &thingProperties

	resp, r, err := apiClient.ThingsApi.AddThing(
		context.Background(),
		space,
		collName).ThingRequest(tr).Execute()
	if err != nil {
		t.Errorf("Error when calling `ThingsApi.AddThing``: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `AddThing`: ThingCreateResponse
	if r.StatusCode != 201 || *resp.Collection != collName || *resp.Title != thingTitle || *resp.Description != thingDesc {
		t.Errorf("%v failed; expected 201, got %v with collection = %v, thingTitle = %v & thingDesc = %v", 
			t.Name(), r.StatusCode, *resp.Collection, *resp.Title, *resp.Description)
	}
	thingUID = *resp.Uid
}

func TestListThings(t *testing.T) {
	resp, r, err := apiClient.ThingsApi.ListThings(
		context.Background(),
		space,
		collName).Execute()
	if err != nil {
		t.Errorf("Error when calling `ThingsApi.ListThings``: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `ListThings`: ThingListResponse
	if r.StatusCode != 200 || len(*resp.Data) == 0 {
		t.Errorf("%v failed; expected 200, got %v with http response body = %v", t.Name(), r.StatusCode, r.Body)
	}
}

func TestUpdateThingDesc(t *testing.T) {
	thingDesc := "My test thing"

	// ThingUpdateRequest | Update an existing thing description by Id
	tur := *NewThingUpdateRequest()
	tur.Title = &thingTitle
	tur.Description = &thingDesc

	resp, r, err := apiClient.ThingsApi.UpdateThing(
		context.Background(),
		space,
		collName,
		thingUID).ThingUpdateRequest(tur).Execute()
	if err != nil {
		t.Errorf("Error when calling `ThingsApi.UpdateThing``: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `UpdateThing`: ThingUpdateResponse
	if r.StatusCode != 200 || *resp.Collection != collName || *resp.Title != thingTitle || 
		*resp.Description != thingDesc || *resp.Uid != thingUID {
			t.Errorf("%v failed; expected 200, got %v with collection = %v, thingTitle = %v, thingDesc = %v & thingUID = %v", 
				t.Name(), r.StatusCode, *resp.Collection, *resp.Title, *resp.Description, *resp.Uid)
	}
}

func TestUpdateThingProperty(t *testing.T) {
	thingDesc := "My test thing"
	thingProperty1 := properties{
		Description: "Type changed for Test Property 1",
		ReadOnly:    false,
		Title:       "Test Property 1",
		Type:        "string",
	}
	thingProperty2 := properties{
		Description: "This is Test Property 2",
		ReadOnly:    false,
		Title:       "Test Property 2",
		Unit:        "percent",
		Type:        "number",
	}

	thingProperties := map[string]interface{}{
		"property1": thingProperty1,
		"property2": thingProperty2,
	}

	actionInput1 := actionInput{
		Maximum: 100,
		Minimum: 3,
		Type:    "number",
	}
	actionProperties1 := actionProperties{
		Input: actionInput1,
	}
	actInput1 := actInput{
		Properties: actionProperties1,
	}
	thingAction1 := actions{
		Title:       "Test Action 1",
		Description: "This is Test Action 1",
		Input:       actInput1,
	}
	thingActions := map[string]interface{}{
		"action1": thingAction1,
	}

	eventData1 := eventData{
		Type: "number",
		Unit: "percent",
	}
	thingEvent1 := events{
		Title:       "Test Event 1",
		Description: "This is Test Event 1",
		Data:        eventData1,
	}
	thingEvents := map[string]interface{}{
		"event1": thingEvent1,
	}

	// ThingUpdateRequest | Update an existing thing description and properties by Id
	tur := *NewThingUpdateRequest()
	tur.Title = &thingTitle
	tur.Description = &thingDesc
	tur.Properties = &thingProperties
	tur.Actions = &thingActions
	tur.Events = &thingEvents

	resp, r, err := apiClient.ThingsApi.UpdateThing(
		context.Background(),
		space,
		collName,
		thingUID).ThingUpdateRequest(tur).Execute()
	if err != nil {
		t.Errorf("Error when calling `ThingsApi.UpdateThing``: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `UpdateThing`
	if r.StatusCode != 200 || *resp.Collection != collName || *resp.Title != thingTitle || 
		*resp.Description != thingDesc || *resp.Uid != thingUID {
			t.Errorf("%v failed; expected 200, got %v with collection = %v, thingTitle = %v, thingDesc = %v & thingUID = %v", 
				t.Name(), r.StatusCode, *resp.Collection, *resp.Title, *resp.Description, *resp.Uid)
	}
}

func TestShowThing(t *testing.T) {
	thingDesc := "My test thing"
	resp, r, err := apiClient.ThingsApi.ShowThing(
		context.Background(),
		space,
		collName,
		thingUID).Execute()
	if err != nil {
		t.Errorf("Error when calling `ThingsApi.ShowThing``: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `ShowThing`: ThingResponse
	if r.StatusCode != 200 || *resp.Collection != collName || *resp.Title != thingTitle || 
		*resp.Description != thingDesc || *resp.Uid != thingUID {
			t.Errorf("%v failed; expected 200, got %v with collection = %v, thingTitle = %v, thingDesc = %v & thingUID = %v", 
				t.Name(), r.StatusCode, *resp.Collection, *resp.Title, *resp.Description, *resp.Uid)
	}
}

// ****************************************
// Negative Tests for thing APIs
// ****************************************

func TestAddThingCollectionNotFound(t *testing.T) {
	thingDesc := "My thing"

	// ThingRequest | Create a new thing description in the platform
	tr := *NewThingRequest()
	tr.Title = &thingTitle
	tr.Description = &thingDesc

	_, r, _ := apiClient.ThingsApi.AddThing(
		context.Background(),
		space,
		"collection_not_found_1").ThingRequest(tr).Execute()

	// Check response from `AddThing`: ThingCreateResponse
	er := errorResponse{}
	json.NewDecoder(r.Body).Decode(&er)
	if r.StatusCode != 400 || er.Error.Message == "" {
		t.Errorf("%v failed: expected 400 with non-empty error message, got %v with error message = %v", 
			t.Name(), r.StatusCode, er.Error.Message)
	}
}

func TestUpdateThingNotFound(t *testing.T) {
	thingDesc := "My test thing"

	// ThingUpdateRequest | Update an existent thing description by Id
	tur := *NewThingUpdateRequest()
	tur.Title = &thingTitle
	tur.Description = &thingDesc

	_, r, _ := apiClient.ThingsApi.UpdateThing(
		context.Background(),
		space,
		"thingUid_not_found_1",
		collName).ThingUpdateRequest(tur).Execute()

	// Check response from `UpdateThing`: ThingUpdateResponse
	er := errorResponse{}
	json.NewDecoder(r.Body).Decode(&er)
	if r.StatusCode != 404 || er.Error.Message == "" {
		t.Errorf("%v failed: expected 404 with non-empty error message, got %v with error message = %v", 
			t.Name(), r.StatusCode, er.Error.Message)
	}
}

// ****************************************
// Positive Tests for property APIs
// ****************************************

func TestUpdateProperty(t *testing.T) {
	thingPropertyBody := map[string]interface{}{
		"property2": 100,
	}

	// Update an existing thing property
	resp, r, err := apiClient.PropertiesApi.UpdateProperty(
		context.Background(),
		space,
		collName,
		thingUID,
		"property2").RequestBody(thingPropertyBody).Execute()
	if err != nil {
		t.Errorf("Error when calling `PropertiesApi.UpdateProperty`: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	if r.StatusCode != 201 || resp["property2"] != 100.0 {
		t.Errorf("%v failed; expected 201 with property2 = 100, got %v with property2 = %v", 
			t.Name(), r.StatusCode, resp["property2"])
	}
}

func TestListProperties(t *testing.T) {
	resp, r, err := apiClient.PropertiesApi.ListProperties(
		context.Background(),
		space,
		collName,
		thingUID).Execute()
	if err != nil {
		t.Errorf("Error when calling `PropertiesApi.ListProperties`: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `ListProperties`
	fmt.Println(r.Body)
	if r.StatusCode != 200 || resp["property2"] != 100.0 {
		t.Errorf("%v failed; expected 200 with property2 = 100, got %v with property2 = %v", 
			t.Name(), r.StatusCode, resp["property2"])
	}
}

func TestShowProperty(t *testing.T) {
	resp, r, err := apiClient.PropertiesApi.ShowProperty(
		context.Background(),
		space,
		collName,
		thingUID,
		"property2").Execute()
	if err != nil {
		t.Errorf("Error when calling `PropertiesApi.ShowProperty`: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `ShowProperty`
	fmt.Println(r.Body)
	if r.StatusCode != 200 || resp["property2"] != 100.0 {
		t.Errorf("%v failed; expected 200 with property2 = 100, got %v with property2 = %v", 
			t.Name(), r.StatusCode, resp["property2"])
	}
}

// ****************************************
// Negative Tests for property APIs
// ****************************************

func TestUpdatePropertyNotFound(t *testing.T) {

	thingPropertyBody := map[string]interface{}{
		"property2": 100,
	}

	// Update an existing thing property
	_, r, _ := apiClient.PropertiesApi.UpdateProperty(
		context.Background(),
		space,
		collName,
		thingUID,
		"property_not_found_1").RequestBody(thingPropertyBody).Execute()

	// Check response from `UpdateProperty`
	er := errorResponse{}
	json.NewDecoder(r.Body).Decode(&er)
	if r.StatusCode != 400 || er.Error.Message == "" {
		t.Errorf("%v failed: expected 400 with non-empty error message, got %v with error message = %v", 
			t.Name(), r.StatusCode, er.Error.Message)
	}
}

// ****************************************
// Positive Tests for model, thing & collection Delete APIs
// ****************************************

func TestDeleteModel(t *testing.T) {
	r, err := apiClient.ModelsApi.DeleteModel(
		context.Background(),
		space,
		collName,
		modelName).Execute()
	if err != nil {
		t.Errorf("Error when calling `ModelsApi.DeleteModel``: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `DeleteModel`: ModelDeleteResponse
	if r.StatusCode != 204 {
		t.Errorf("%v failed; expected 204, got %v with http response body = %v", t.Name(), r.StatusCode, r.Body)
	}
}

func TestDeleteThing(t *testing.T) {
	r, err := apiClient.ThingsApi.DeleteThing(
		context.Background(),
		space,
		collName,
		thingUID).Execute()
	if err != nil {
		t.Errorf("Error when calling `ThingsApi.DeleteThing``: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `DeleteThing`: ThingDeleteResponse
	if r.StatusCode != 204 {
		t.Errorf("%v failed; expected 204, got %v with http response body = %v", t.Name(), r.StatusCode, r.Body)
	}
}

func TestDeleteCollection(t *testing.T) {
	r, err := apiClient.CollectionsApi.DeleteCollection(
		context.Background(),
		space,
		collName).Execute()
	if err != nil {
		t.Errorf("Error when calling `CollectionsApi.DeleteCollection``: %v\n", err)
		t.Errorf("Full HTTP response: %v\n", r)
		return
	}

	// Check response from `DeleteCollection`: CollectionDeleteResponse
	if r.StatusCode != 204 {
		t.Errorf("%v failed; expected 204, got %v with http response body = %v", t.Name(), r.StatusCode, r.Body)
	}
}
