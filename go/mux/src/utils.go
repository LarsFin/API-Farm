package apifarm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"time"
)

// Custom Time

type CustomTime struct {
	time.Time
}

const layout = "02/01/2006"

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	if s == "null" {
		ct.Time = time.Time{}
		return nil
	}

	t, err := time.Parse(layout, s)
	ct.Time = t

	return err
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("\"%s\"", ct.Time.Format(layout))

	return []byte(s), nil
}

// Data Utilities

type DataUtils interface {
	Serialize(interface{}) ([]byte, error)
	DeserializeConfiguration([]byte) (*Configuration, error)
	DeserializeVideoGame([]byte) (*VideoGame, error)
	DeserializeVideoGames([]byte) (*[]VideoGame, error)
}

type JSON struct {
}

func (*JSON) Serialize(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func (*JSON) DeserializeConfiguration(data []byte) (*Configuration, error) {
	var c Configuration
	err := json.Unmarshal(data, &c)
	return &c, err
}

func (*JSON) DeserializeVideoGame(data []byte) (*VideoGame, error) {
	var vg VideoGame

	var m map[string]json.RawMessage
	err := json.Unmarshal(data, &m)

	if err != nil {
		return &vg, err
	}

	t := reflect.TypeOf(vg)

	for k := range m {
		if k == "id" || !hasFieldWithJSONTag(t, k) {
			return &vg, &InvalidAttributeError{Attribute: k}
		}
	}

	err = json.Unmarshal(data, &vg)

	return &vg, err
}

func (*JSON) DeserializeVideoGames(data []byte) (*[]VideoGame, error) {
	var vgs []VideoGame
	err := json.Unmarshal(data, &vgs)
	return &vgs, err
}

func hasFieldWithJSONTag(t reflect.Type, jt string) bool {
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("json") == jt {
			return true
		}
	}

	return false
}

// Errors

type InvalidAttributeError struct {
	Attribute string
}

func (err *InvalidAttributeError) Error() string {
	return VideoGameInvalidAttribute(err.Attribute)
}

// Factories

type QueryFactory interface {
	BuildResult([]byte, uint) Query
	BuildMessage(string, uint) Query
	Error(error) Query
}

type queryFactory struct {
}

func (*queryFactory) BuildResult(result []byte, code uint) Query {
	return Query{
		result,
		"Successfully built.",
		code,
		nil,
	}
}

func (*queryFactory) BuildMessage(msg string, code uint) Query {
	return Query{
		nil,
		msg,
		code,
		nil,
	}
}

func (*queryFactory) Error(err error) Query {
	return Query{
		nil,
		"An unforeseen error occurred.",
		500,
		err,
	}
}

type Query struct {
	Result  []byte
	Message string
	Code    uint
	Error   error
}

// File Reading

type FileUtils interface {
	Read(string) ([]byte, error)
}

type fileUtils struct {
}

func (*fileUtils) Read(p string) ([]byte, error) {
	return ioutil.ReadFile(p)
}

// Messages

const InvalidJSON = "Invalid JSON in body."
const SuccessfullyLoadedData = "Successfully loaded data."
const VideoGameDateRequired = "A date_released is required for a video game."
const VideoGameNameRequired = "A name is required for a video game."

func ParamInvalidID(invalidID string) string {
	return fmt.Sprintf("The provided id '%s' is invalid.", invalidID)
}

func VideoGameInvalidAttribute(invalidAttribute string) string {
	return fmt.Sprintf("The provided data has an invalid attribute '%s'.", invalidAttribute)
}

func VideoGameInvalidDate(invalidDate string) string {
	return fmt.Sprintf("The provided date_released '%s' is invalid.", invalidDate)
}

func VideoGameNotFound(id uint) string {
	return fmt.Sprintf("No video game with id '%d' could be found.", id)
}
