package tasks

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoStore struct {
	Session        *mgo.Session
	DatabaseName   string
	CollectionName string
}

func (ms *MongoStore) Insert(newTask *NewTask) (*Task, error) {
	t := newTask.ToTask()
	t.ID = bson.NewObjectId()
	// assuming that whoever created the struct has an active session
	err := ms.Session.DB(ms.DatabaseName).C(ms.CollectionName).Insert(t)
	return t, err
}

func (ms *MongoStore) Get(ID interface{}) (*Task, error) {
	task := &Task{}
	err := ms.Session.DB(ms.DatabaseName).C(ms.CollectionName).FindId(ID).One(task)
	return task, err
}
