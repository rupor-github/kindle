package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Delete struct {
	Uuid string `json:"uuid"`
}

type CmdDelete struct {
	Cmd Delete `json:"delete"`
}

type UpdateBook struct {
	CollectionCount int    `json:"collectionCount"`
	Type            string `json:"type"`
	Uuid            string `json:"uuid"`
}

type CmdUpdateBook struct {
	Cmd UpdateBook `json:"update"`
}

type UpdateCollection struct {
	Members []string `json:"members"`
	Type    string   `json:"type"`
	Uuid    string   `json:"uuid"`
}

type CmdUpdateCollection struct {
	Cmd UpdateCollection `json:"update"`
}

type Title struct {
	Direction string `json:"direction"`
	Display   string `json:"display"`
	Language  string `json:"language"`
}

type Insert struct {
	Collectioncount       interface{} `json:"collectionCount"`
	Collectiondatasetname string      `json:"collectionDataSetName"`
	Collections           interface{} `json:"collections"`
	IsArchived            int         `json:"isArchived"`
	IsVisibleInHome       int         `json:"isVisibleInHome"`
	LastAccess            int64       `json:"lastAccess"`
	Titles                []Title     `json:"titles"`
	Type                  string      `json:"type"`
	Uuid                  string      `json:"uuid"`
}

type CmdInsert struct {
	Cmd Insert `json:"insert"`
}

type Queue struct {
	Commands []interface{} `json:"commands"`
	Type     string        `json:"type"`
	Id       int           `json:"id"`
}

func NewQueue() *Queue {
	return &Queue{Commands: make([]interface{}, 0, 1), Id: 1, Type: "ChangeRequest"}
}

func (q *Queue) DeleteCollection(uuid string) {
	q.Commands = append(q.Commands, CmdDelete{Cmd: Delete{Uuid: uuid}})
}

func (q *Queue) UpdateBook(uuid string, count int) {
	q.Commands = append(q.Commands, CmdUpdateBook{Cmd: UpdateBook{Uuid: uuid, Type: "Entry:Item", CollectionCount: count}})
}

func (q *Queue) UpdateCollection(uuid string, books map[string]*book) {
	bus := make([]string, 0, len(books))
	for k, _ := range books {
		bus = append(bus, k)
	}
	q.Commands = append(q.Commands, CmdUpdateCollection{Cmd: UpdateCollection{Uuid: uuid, Type: "Collection", Members: bus}})
}

func (q *Queue) InsertCollection(uuid, title string) {
	q.Commands = append(q.Commands, CmdInsert{Cmd: Insert{
		Collectiondatasetname: uuid,
		IsVisibleInHome:       1,
		LastAccess:            time.Now().Unix(),
		Titles:                []Title{Title{Direction: "LTR", Display: title, Language: locale}},
		Type:                  "Collection",
		Uuid:                  uuid}})
}

type Response struct {
	Ok      bool   `json:"ok"`
	Type    string `json:"type"`
	Changes int    `json:"changes"`
	ID      int    `json:"id"`
}

func (q *Queue) Size() int {
	return len(q.Commands)
}

func (q *Queue) Execute() (err error) {
	if b, err := json.Marshal(q); err == nil {
		if resp, err := http.Post("http://localhost:9101/change", "application/json", bytes.NewBuffer(b)); err == nil {
			defer resp.Body.Close()
			if body, err := ioutil.ReadAll(resp.Body); err == nil {
				if resp.StatusCode != http.StatusOK {
					err = errors.New("HTTP status is not OK")
				} else {
					var res Response
					err = json.Unmarshal(body, &res)
					if err == nil && res.Ok {
						log.Printf("Processed - %d changes...\n", res.Changes)
					} else {
						err = errors.New(fmt.Sprintf("Error - <%s>", body))
					}
				}
			}
		}
	}
	return
}

func (q *Queue) Dump() (err error) {
	if b, err := json.MarshalIndent(q, "", "    "); err == nil {
		err = ioutil.WriteFile("test_"+strconv.FormatInt(time.Now().UnixNano(), 16)+".json", b, 0)
	}
	return
}
