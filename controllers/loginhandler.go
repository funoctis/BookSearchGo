package controllers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

//LoginHandler is the handler function for the register page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var exec = true

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			exec = false
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		passwordByte := []byte(password)
		hash := sha256.Sum256(passwordByte)
		encryptedPassword := base32.StdEncoding.EncodeToString(hash[:])
		
		db, err := sql.Open("sqlite3", "database.db")
		if err != nil {
			log.Println(err.Error())
			http.Redirect(w, r, "", http.StatusPermanentRedirect)
		} else {

			defer db.Close()

			sqlStmt := fmt.Sprintf("SELECT username FROM users WHERE username = '%s'", username)
			rows := db.QueryRow(sqlStmt)
			var checkUsername string
			err = rows.Scan(&checkUsername)
			if err != nil {
				if err == sql.ErrNoRows {
					query := fmt.Sprintf("insert into users(username,password) values('%s', '%s')", username, encryptedPassword)
					_, err = db.Exec(query)
					if err != nil {
						log.Println(err.Error())
						http.Redirect(w, r, "", http.StatusPermanentRedirect)
					}
				} else {
					log.Println(err.Error())
					http.Redirect(w, r, "", http.StatusPermanentRedirect)
				}
			} else {
				exec = false
				_, err = fmt.Fprintf(w, "User already exists. Try another name.")
			}
		}
	}

	if exec == true {
		err := templates.ExecuteTemplate(w, "login", nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Could not serve page", http.StatusNotFound)
		}
	}
}
