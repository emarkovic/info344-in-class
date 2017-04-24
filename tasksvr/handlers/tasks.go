package handlers

import (
	"encoding/json"
	"net/http"

	"path"

	"github.com/emarkovic/info344-in-class/tasksvr/models/tasks"
)

//HandleTasks will handle requests for the /v1/tasks resource
func (ctx *Context) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		newTask := &tasks.NewTask{}
		if err := decoder.Decode(newTask); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		if err := newTask.Validate(); err != nil {
			http.Error(w, "error validating task: "+err.Error(), http.StatusBadRequest)
			return
		}

		task, err := ctx.TasksStore.Insert(newTask)
		if err != nil {
			http.Error(w, "error inserting task: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add(headerContentType, contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(task)
	case "GET":
		tasks, err := ctx.TasksStore.GetAll()
		if err != nil {
			http.Error(w, "error getting tasks: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add(headerContentType, contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(tasks)
	}
}

//HandleSpecificTask will handle requests for the /v1/tasks/some-task-id resource
func (ctx *Context) HandleSpecificTask(w http.ResponseWriter, r *http.Request) {
	// split returns "directory" and the "file" which is the id of the task in our case
	_, id := path.Split(r.URL.Path)

	switch r.Method {
	case "GET":
		task, err := ctx.TasksStore.Get(id)
		if err != nil {
			http.Error(w, "error finding task: "+err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Add(headerContentType, contentTypeJSONUTF8)
		encoder := json.NewEncoder(w)
		encoder.Encode(task)
	case "PATCH":
		decoder := json.NewDecoder(r.Body)
		task := &tasks.Task{}
		if err := decoder.Decode(task); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		task.ID = id

		err := ctx.TasksStore.Update(task)
		if err != nil {
			http.Error(w, "error updating task: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("update successful"))
	}
}

/*
write a function called "respond" that sends the encoded json to the user. set the headers in here too
*/

// func respond()
