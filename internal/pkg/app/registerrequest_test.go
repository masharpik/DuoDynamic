package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/ds"
)

type psevdoRepo struct {
	Users     *[]ds.User
	Cities    *[]ds.City
	UserToken *map[uint]string
}

func convertToString(val interface{}) (strVal string) {
	switch val.(type) {
	case string:
		strVal = val.(string)
	case float64:
		strVal = fmt.Sprint(val.(float64))
	case int:
		strVal = fmt.Sprint(val.(int))
	}
	return strVal
}

func mapEqual(got, expected map[string]interface{}) bool {
	for keyGot, valueGot := range got {
		valueExpected := expected[keyGot]

		var (
			strValueExpected string = convertToString(valueExpected)
			strValueGot      string = convertToString(valueGot)
		)
		if strValueExpected != strValueGot {
			return false
		}
	}
	return true
}

func TestApplication_Register(t *testing.T) {
	testCases := []struct {
		inpJson    string
		inpMethod  string
		outputJson map[string]interface{}
	}{
		{
			inpJson:    `{"username": "masharpik", "email": "masharpik@gmail.com", "password": "masharpik2004", "age": 19}`,
			inpMethod:  "POST",
			outputJson: map[string]interface{}{"err": "", "status": http.StatusOK},
		},
		{
			inpJson:    `{"username": "masharpik", "email": "masharpik@gmail.com", "password": "masharpik2004", "age": 19}`,
			inpMethod:  "GET",
			outputJson: map[string]interface{}{"err": "only POST method is supported for this route", "status": http.StatusNotFound},
		},
	}
	r := NewRepo()
	a := New(r)
	ts := httptest.NewServer(http.HandlerFunc(a.Logout))
	handler := http.HandlerFunc(a.Register)

	for _, tCase := range testCases {

		req, err := http.NewRequest(tCase.inpMethod, ts.URL, bytes.NewBuffer([]byte(tCase.inpJson)))
		if err != nil {
			t.Errorf(err.Error())
		}

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		if err != nil {
			t.Errorf(err.Error())
		}

		var jsonStr []byte
		jsonStr, err = ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Errorf(err.Error())
		}

		var result map[string]interface{}
		err = json.Unmarshal([]byte(jsonStr), &result)
		if err != nil {
			t.Errorf(err.Error())
		}
		if err != nil {
			t.Errorf("unexpected error: %#v", err)
		}

		if !mapEqual(result, tCase.outputJson) {
			t.Errorf(" wrong result, expected %#v, got %#v", tCase.outputJson, result)
		}
	}
}

func TestApplication_Login(t *testing.T) {
	testCases := []struct {
		inpJson    string
		inpMethod  string
		outputJson map[string]interface{}
	}{
		{
			inpJson:    `{"input": "masharpik", "password": "masharpik2004"}`,
			inpMethod:  "POST",
			outputJson: map[string]interface{}{"err": "", "status": http.StatusOK},
		},
		{
			inpJson:    `{"input": "", "password": "masharpik2004"}`,
			inpMethod:  "POST",
			outputJson: map[string]interface{}{"err": "empty login", "status": http.StatusBadRequest},
		},
		{
			inpJson:    `{"input": "", "password": "masharpik2004"}`,
			inpMethod:  "GET",
			outputJson: map[string]interface{}{"err": "only POST method is supported for this route", "status": http.StatusNotFound},
		},
	}
	r := NewRepo()
	a := New(r)
	ts := httptest.NewServer(http.HandlerFunc(a.Login))
	handler := http.HandlerFunc(a.Login)
	for _, tCase := range testCases {

		req, err := http.NewRequest(tCase.inpMethod, ts.URL, bytes.NewBuffer([]byte(tCase.inpJson)))
		if err != nil {
			t.Errorf(err.Error())
		}

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		if err != nil {
			t.Errorf(err.Error())
		}

		var jsonStr []byte
		jsonStr, err = ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Errorf(err.Error())
		}

		var result map[string]interface{}
		err = json.Unmarshal([]byte(jsonStr), &result)
		if err != nil {
			t.Errorf(err.Error())
		}
		if err != nil {
			t.Errorf("unexpected error: %#v", err)
		}

		if !mapEqual(result, tCase.outputJson) {
			t.Errorf(" wrong result, expected %#v, got %#v", tCase.outputJson, result)
		}
	}
}

func TestApplication_Logout(t *testing.T) {
	testCases := []struct {
		inpJson    string
		inpMethod  string
		inpCookies http.Cookie
		outputJson map[string]interface{}
	}{
		{
			inpJson:   `{}`,
			inpMethod: "POST",
			inpCookies: http.Cookie{
				Name:     SessionTokenCookieName,
				Value:    "",
				Expires:  time.Now().Add(120 * time.Second),
				HttpOnly: true,
			},
			outputJson: map[string]interface{}{"err": "", "status": http.StatusOK},
		},
		{
			inpJson:    `{}`,
			inpMethod:  "POST",
			inpCookies: http.Cookie{},
			outputJson: map[string]interface{}{"err": "http: named cookie not present", "status": http.StatusUnauthorized},
		},
		{
			inpJson:    `{}`,
			inpMethod:  "GET",
			inpCookies: http.Cookie{},
			outputJson: map[string]interface{}{"err": "only POST method is supported for this route", "status": http.StatusNotFound},
		},
	}
	r := NewRepo()
	a := New(r)
	ts := httptest.NewServer(http.HandlerFunc(a.Logout))
	handler := http.HandlerFunc(a.Logout)
	for _, tCase := range testCases {

		req, err := http.NewRequest(tCase.inpMethod, ts.URL, bytes.NewBuffer([]byte(tCase.inpJson)))
		if err != nil {
			t.Errorf(err.Error())
		}

		req.AddCookie(&tCase.inpCookies)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		if err != nil {
			t.Errorf(err.Error())
		}

		var jsonStr []byte
		jsonStr, err = ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Errorf(err.Error())
		}

		var result map[string]interface{}
		err = json.Unmarshal([]byte(jsonStr), &result)
		if err != nil {
			t.Errorf(err.Error())
		}

		if !mapEqual(result, tCase.outputJson) {
			t.Errorf(" wrong result, expected %#v, got %#v", tCase.outputJson, result)
		}
	}
}

func TestApplication_GetCities(t *testing.T) {
	testCases := []struct {
		inpMethod  string
		outputJson map[string]interface{}
	}{
		{
			inpMethod:  "GET",
			outputJson: map[string]interface{}{"err": "", "status": http.StatusOK},
		},
		{
			inpMethod:  "POST",
			outputJson: map[string]interface{}{"err": "only GET method is supported for this route", "status": http.StatusNotFound},
		},
	}
	r := NewRepo()
	a := New(r)
	ts := httptest.NewServer(http.HandlerFunc(a.GetCities))
	handler := http.HandlerFunc(a.GetCities)
	for _, tCase := range testCases {

		req, err := http.NewRequest(tCase.inpMethod, ts.URL, nil)
		if err != nil {
			t.Errorf(err.Error())
		}

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		if err != nil {
			t.Errorf(err.Error())
		}

		var jsonStr []byte
		jsonStr, err = ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Errorf(err.Error())
		}

		var result map[string]interface{}
		err = json.Unmarshal([]byte(jsonStr), &result)
		if err != nil {
			t.Errorf(err.Error())
		}

		if !mapEqual(result, tCase.outputJson) {
			t.Errorf(" wrong result, expected %#v, got %#v", tCase.outputJson, result)
		}
	}
}

func TestApplication_GetCurrentUser(t *testing.T) {
	testCases := []struct {
		inpMethod  string
		inpCookies http.Cookie
		outputJson map[string]interface{}
	}{
		{
			inpMethod: "GET",
			inpCookies: http.Cookie{
				Name:     SessionTokenCookieName,
				Value:    "",
				Expires:  time.Now().Add(120 * time.Second),
				HttpOnly: true,
			},
			outputJson: map[string]interface{}{"Age": 20, "Avatar": "", "City": "", "Description": "", "Email": "", "Sex": 0, "Username": "", "err": "", "status": http.StatusOK},
		},
		{
			inpMethod:  "POST",
			inpCookies: http.Cookie{},
			outputJson: map[string]interface{}{"err": "only GET method is supported for this route", "status": http.StatusNotFound},
		},
		{
			inpMethod:  "GET",
			inpCookies: http.Cookie{},
			outputJson: map[string]interface{}{"err": "http: named cookie not present", "status": http.StatusUnauthorized},
		},
	}
	r := NewRepo()
	a := New(r)
	ts := httptest.NewServer(http.HandlerFunc(a.GetCurrentUser))
	handler := http.HandlerFunc(a.GetCurrentUser)
	for _, tCase := range testCases {

		req, err := http.NewRequest(tCase.inpMethod, ts.URL, nil)
		if err != nil {
			t.Errorf(err.Error())
		}

		req.AddCookie(&tCase.inpCookies)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		if err != nil {
			t.Errorf(err.Error())
		}

		var jsonStr []byte
		jsonStr, err = ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Errorf(err.Error())
		}

		var result map[string]interface{}
		err = json.Unmarshal([]byte(jsonStr), &result)
		if err != nil {
			t.Errorf(err.Error())
		}

		if !mapEqual(result, tCase.outputJson) {
			t.Errorf(" wrong result, expected %#v, got %#v", tCase.outputJson, result)
		}
	}
}

func TestApplication_GetRecommendations(t *testing.T) {
	testCases := []struct {
		inpMethod  string
		inpCookies http.Cookie
		outputJson map[string]interface{}
	}{
		{
			inpMethod: "GET",
			inpCookies: http.Cookie{
				Name:     SessionTokenCookieName,
				Value:    "",
				Expires:  time.Now().Add(120 * time.Second),
				HttpOnly: true,
			},
			outputJson: map[string]interface{}{"err": "", "status": http.StatusOK},
		},
		{
			inpMethod:  "GET",
			inpCookies: http.Cookie{},
			outputJson: map[string]interface{}{"err": "http: named cookie not present", "status": http.StatusUnauthorized},
		},
		{
			inpMethod:  "POST",
			inpCookies: http.Cookie{},
			outputJson: map[string]interface{}{"err": "only GET method is supported for this route", "status": http.StatusNotFound},
		},
	}
	r := NewRepo()
	a := New(r)
	ts := httptest.NewServer(http.HandlerFunc(a.GetRecommendations))
	handler := http.HandlerFunc(a.GetRecommendations)
	for _, tCase := range testCases {

		req, err := http.NewRequest(tCase.inpMethod, ts.URL, nil)
		if err != nil {
			t.Errorf(err.Error())
		}

		req.AddCookie(&tCase.inpCookies)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		if err != nil {
			t.Errorf(err.Error())
		}

		var jsonStr []byte
		jsonStr, err = ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Errorf(err.Error())
		}

		var result map[string]interface{}
		err = json.Unmarshal([]byte(jsonStr), &result)
		if err != nil {
			t.Errorf(err.Error())
		}

		if !mapEqual(result, tCase.outputJson) {
			t.Errorf(" wrong result, expected %#v, got %#v", tCase.outputJson, result)
		}
	}
}

///PsevdoRepo

func NewRepo() *psevdoRepo {
	var userDS []ds.User
	var cityDS []ds.City
	tokenDS := make(map[uint]string)
	r := psevdoRepo{&userDS, &cityDS, &tokenDS}

	return &r
}

func (pr *psevdoRepo) AddUser(_ *ds.User) error {
	return nil
}

func (pr *psevdoRepo) DeleteToken(_ string) error {
	return nil
}

func (pr *psevdoRepo) Login(_, _ string) (userId uint, err error) {
	return 0, nil
}

func (pr *psevdoRepo) SaveToken(_ uint, _ string) {

}

func (pr *psevdoRepo) GetCities() ([]string, error) {
	return nil, nil
}

func (pr *psevdoRepo) GetUserIdByToken(string) (uint, error) {
	return 0, nil
}

func (pr *psevdoRepo) GetUserById(uint) (*UserRes, error) {
	ur := UserRes{
		Username:    "",
		Age:         20,
		Sex:         0,
		City:        "",
		Description: "",
		Email:       "",
		Avatar:      "",
	}
	return &ur, nil
}

func (pr *psevdoRepo) GetRecommendation(uint) ([]*Recommendation, error) {
	return nil, nil
}
