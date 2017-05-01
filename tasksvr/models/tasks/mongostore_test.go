package tasks

import (
	"testing"

	mgo "gopkg.in/mgo.v2"
)

func TestCRUD(t *testing.T) {
	sess, err := mgo.Dial("localhost:27017")
	if err != nil {
		// stops automated tests
		t.Fatalf("error dialing Mongo: %v", err)
	}
	defer sess.Close()

	store := &MongoStore{
		Session:        sess,
		DatabaseName:   "test",
		CollectionName: "tasks",
	}

	newTask := &NewTask{
		Title: "Learn MongoDB",
		Tags:  []string{"mongo", "info344"},
	}
	task, err := store.Insert(newTask)
	if err != nil {
		t.Errorf("error inserting new task: %v", err)
	}

	task2, err := store.Get(task.ID)
	if err != nil {
		t.Errorf("error fetching task: %v", err)
	}
	if task2.Title != task.Title {
		t.Errorf("task title didnt match, expected %s but got %s", task.Title, task2.Title)
	}

	// sess.DB(store.DatabaseName).C(store.CollectionName).RemoveAll(nil)
}
