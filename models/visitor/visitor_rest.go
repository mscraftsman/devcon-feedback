package visitor
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type responseSingle struct {
	Status   bool      `json:"status"`
	Messages []message `json:"messages"`
	Entity   *Visitor `json:"entity"`
}

type responseList struct {
	Status   bool                `json:"status"`
	Messages []message           `json:"messages"`
	Entities []*Visitor `json:"entities"`
}

type message struct {
	Type    rune   `json:"type"`
	Message string `json:"message"`
}

//RegisterRoutes registers routes with a mux Router
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("//visitors/{id}", RestGet).Methods("GET")
	router.HandleFunc("//visitors", RestList).Methods("GET")
	router.HandleFunc("//visitors", RestCreate).Methods("POST")
	router.HandleFunc("//visitors/{id}", RestUpdate).Methods("PUT")
	router.HandleFunc("//visitors/{id}", RestDelete).Methods("DELETE")
}


//RestGet is a REST endpoint for GET //visitors/{id}
func RestGet(w http.ResponseWriter, r *http.Request) {
	var (
		id       int64
		err      error
		response responseSingle
		stop     bool
	)

	vars := mux.Vars(r)
	valid := false
	if _, ok := vars["id"]; ok {
		id, err = strconv.ParseInt(vars["id"], 10, 64)
		valid = err == nil && id > 0
	}

	if !valid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Invalid ID"}]}`
		return
	}

	
    if stop, err = restPreGet(w, r, id); err != nil || stop {
        return
    }
    

	response.Entity, err = Get(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "An error occurred"}]}`
		return
	}

	if response.Entity == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Entity not found"}]}`
		return
	}

	
    if stop, err = restPostGet(w, r, response.Entity); err != nil || stop {
        return
    }
    

	response.Status = true
	output, err := json.Marshal(response)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "JSON encoding failed"}]}`
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(output))
}



//RestList is a REST endpoint for GET //visitors
func RestList(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		response responseList
		filters  []models.ListFilter
		stop     bool
	)
	

	
    if filters, stop, err = restPreList(w, r, filters); err != nil || stop {
        return
    }
    

	response.Entities, err = List(filters)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "An error occurred"}]}`
		return
	}

	
    if response.Entities, stop, err = restPostList(w, r, response.Entities); err != nil || stop {
        return
    }
    

	response.Status = true
	output, err := json.Marshal(response)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "JSON encoding failed"}]}`
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(output))
}



//RestCreate is a REST endpoint for POST //visitors
func RestCreate(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		rawbody  []byte
		response responseSingle
		tx       *sql.Tx
		stop     bool
	)

	rawbody, err = ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Failed to read body"}]}`
		return
	}

	response.Entity = New()
	err = json.Unmarshal(rawbody, response.Entity)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Failed to decode body"}]}`
		return
	}
	response.Entity.ID = nil

	tx, err = db.Begin()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Failed to process"}]}`
		return
	}

	
	if stop, err = restPreCreate(w, r, response.Entity, tx); err != nil {
		tx.Rollback()
		return
	} else if stop {
		return
	}
    

	err = response.Entity.Save(tx, false)
	if err != nil {
		tx.Rollback()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Save failed"}]}`
		return
	}

	
	if stop, err = restPostCreate(w, r, response.Entity, tx); err != nil {
		tx.Rollback()
		return
	} else if stop {
		return
	}
	
	
	if err = tx.Commit(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "E", "message": "RestCreate could not commit transaction"}]}`
		return
	}

	output, err := json.Marshal(response)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "JSON encoding failed"}]}`
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(output))
}



//RestUpdate is a REST endpoint for PUT //visitors/{id}
func RestUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		rawbody  []byte
		id       int64
		response responseSingle
		tx       *sql.Tx
		stop     bool
	)

	vars := mux.Vars(r)
	valid := false
	if _, ok := vars["id"]; ok {
		id, err = strconv.ParseInt(vars["id"], 10, 64)
		valid = err == nil && id > 0
	}

	if !valid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Invalid ID"}]}`
		return
	}

	response.Entity, err = Get(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "An error occurred"}]}`
		return
	}

	if response.Entity == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Entity not found"}]}`
		return
	}

	rawbody, err = ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Failed to read body"}]}`
		return
	}

	err = json.Unmarshal(rawbody, response.Entity)
	if err != nil {
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Failed to decode body"}]}`
			return
		}
	}
	response.Entity.ID = &id

	tx, err = db.Begin()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Failed to process"}]}`
		return
	}

	
    if stop, err = restPreUpdate(w, r, response.Entity, tx); err != nil {
		tx.Rollback()
        return
    } else if stop {
		return
	}
    

	err = response.Entity.Save(tx, false)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Save failed"}]}`
		return
	}

	
    if stop, err = restPostUpdate(w, r, response.Entity, tx); err != nil {
		tx.Rollback()
        return
    } else if stop {
		return
	}
	
	
	if err = tx.Commit(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "E", "message": "RestUpdate could not commit transaction"}]}`
		return
	}

	output, err := json.Marshal(response)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "JSON encoding failed"}]}`
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(output))
}



//RestDelete is a REST endpoint for DELETE //visitors/{id}
func RestDelete(w http.ResponseWriter, r *http.Request) {
	var (
		id       int64
		err      error
		response responseSingle
		tx       *sql.Tx
		stop     bool
	)

	vars := mux.Vars(r)
	valid := false
	if _, ok := vars["id"]; ok {
		id, err = strconv.ParseInt(vars["id"], 10, 64)
		valid = err == nil && id > 0
	}

	if !valid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Invalid ID"}]}`
		return
	}

	response.Entity, err = Get(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "An error occurred"}]}`
		return
	}

	if response.Entity == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Entity not found"}]}`
		return
	}

	tx, err = db.Begin()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Failed to process"}]}`
		return
	}
	
	if stop, err = restPreDelete(w, r, id, tx); err != nil {
		tx.Rollback()
		return
	} else if stop {
		return
	}
    
	err = response.Entity.Delete(tx, false)
	if err != nil {
		tx.Rollback()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "Delete failed"}]}`
		return
	}
	
	if stop, err = restPostDelete(w, r, id, tx); err != nil {
		tx.Rollback()
		return
	} else if stop {
		return
	}
	
	if err = tx.Commit(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "E", "message": "RestDelete could not commit transaction"}]}`
		return
	}

	output, err := json.Marshal(response)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status": false, "messages": [{"type": "error", "text": "JSON encoding failed"}]}`
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(output))
}
