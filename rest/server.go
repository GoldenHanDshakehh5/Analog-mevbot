/*
	Copyright 2019 whiteblock Inc.
	This file is a part of the genesis.

	Genesis is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    Genesis is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/whiteblock/genesis/db"
	"log"
	"net/http"
	"strconv"
)

func getAllServerInfo(w http.ResponseWriter, r *http.Request) {
	servers, err := db.GetAllServers()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 204)
		return
	}
	json.NewEncoder(w).Encode(servers)
}

func addNewServer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var server db.Server
	err := json.NewDecoder(r.Body).Decode(&server)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = server.Validate()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	log.Println(fmt.Sprintf("Adding server: %+v", server))

	id, err := db.InsertServer(params["name"], server)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(strconv.Itoa(id)))
}

func getServerInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	server, _, err := db.GetServer(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(server)
	if err != nil {
		log.Println(err.Error())
	}
}

func deleteServer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid id", 400)
		return
	}
	db.DeleteServer(id)
	w.Write([]byte("Success"))
}

func updateServerInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var server db.Server

	err := json.NewDecoder(r.Body).Decode(&server)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err = server.Validate()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid id", 400)
		return
	}

	err = db.UpdateServer(id, server)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte("Success"))
}
