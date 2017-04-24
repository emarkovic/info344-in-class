package tasks

import (
	"time"

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
	if sID, ok := ID.(string); ok {
		ID = bson.ObjectIdHex(sID)
	}
	task := &Task{}
	err := ms.Session.DB(ms.DatabaseName).C(ms.CollectionName).FindId(ID).One(task)
	return task, err
}

func (ms *MongoStore) GetAll() ([]*Task, error) {
	tasks := []*Task{}
	err := ms.Session.DB(ms.DatabaseName).C(ms.CollectionName).Find(nil).All(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ms *MongoStore) Update(task *Task) error {
	if sID, ok := task.ID.(string); ok {
		task.ID = bson.ObjectIdHex(sID)
	}

	task.ModifiedAt = time.Now()
	// only allow modifiedAt and completed to be modified in the db
	col := ms.Session.DB(ms.DatabaseName).C(ms.CollectionName)
	updates := bson.M{"$set": bson.M{"complete": task.Complete, "modifiedAt": task.ModifiedAt}}
	return col.UpdateId(task.ID, updates)
}
