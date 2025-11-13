package main

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall/js"

	// Packages
	jsutil "github.com/djthorpe/go-wasmbuild/pkg/js"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Model struct {
	jsutil.EventTarget

	view    Component
	path    string
	columns []string
	rows    []Employee
}

type Employee struct {
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
	Location string  `json:"location"`
}

func NewModel(path string, columns ...string) *Model {
	model := new(Model)
	model.EventTarget = *jsutil.NewEventTarget()
	model.columns = columns
	model.path = path

	// Return the controller
	return model
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	EventLoad = "wasmbuild-table-model-load"
)

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (model *Model) String() string {
	data, err := json.MarshalIndent(model.rows, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (model *Employee) String() string {
	data, err := json.MarshalIndent(model, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (model *Model) Columns() []string {
	return model.columns
}

func (model *Model) Count() int {
	return len(model.rows)
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (model *Model) Load() {
	// Use the new Fetch API
	promise := jsutil.Fetch(model.path)

	// Handle the response with the new Promise bindings
	promise.Then(func(value jsutil.Value) interface{} {
		// Get response object
		response := jsutil.NewResponseFromValue(value)

		// Check if response is OK
		if !response.Ok() {
			fmt.Println("Failed to fetch data:", response.Status(), response.StatusText())
			return nil
		}

		// Parse JSON and chain another promise
		jsonPromise := response.JSON()
		jsonPromise.Then(func(jsonData jsutil.Value) interface{} {
			// Convert JS value to JSON string, then unmarshal to Go struct
			jsonString := js.Global().Get("JSON").Call("stringify", jsonData).String()

			// Parse data
			if err := json.Unmarshal([]byte(jsonString), &model.rows); err != nil {
				fmt.Fprintln(os.Stderr, "Failed to parse JSON:", err.Error())
				return nil
			} else {
				model.DispatchEvent(jsutil.NewEvent(EventLoad))
			}

			// Return success
			return nil
		}).Catch(func(reason jsutil.Value) interface{} {
			fmt.Println("JSON parse error:", reason.String())
			return nil
		})

		return nil
	}).Catch(func(reason jsutil.Value) interface{} {
		fmt.Println("Fetch error:", reason.String())
		return nil
	})
}

func (model *Model) Get(offset, limit int) []Employee {
	if offset < 0 {
		offset = 0
	}
	if limit < 1 {
		limit = 1
	}
	if offset >= len(model.rows) {
		return []Employee{}
	}

	end := offset + limit
	if end > len(model.rows) {
		end = len(model.rows)
	}
	return model.rows[offset:end]
}
