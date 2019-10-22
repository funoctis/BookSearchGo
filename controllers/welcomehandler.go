package controllers

import (
	"BookSearchGo/parsers"
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

//WelcomeHandler is the handler function for the /welcome page.
//This page is accessible only if the user has signed in.
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			log.Println("Error while reading cookie: ", err.Error())
			http.Redirect(w, r, "", http.StatusPermanentRedirect)
		} else {
			sessionToken := cookie.Value
			db, err := sql.Open("sqlite3", "database.db")
			if err != nil {
				log.Println(err.Error())
				http.Redirect(w, r, "", http.StatusPermanentRedirect)
			} else {

				defer db.Close()

				query := fmt.Sprintf("select * from sessions where session_token = '%s'", sessionToken)
				row := db.QueryRow(query)
				var storedSessionID int
				var storedSessionToken, storedUsername string
				err := row.Scan(&storedSessionID, &storedSessionToken, &storedUsername)
				if err != nil {
					log.Println("Error while scanning row: ", err.Error())
					http.Redirect(w, r, "", http.StatusPermanentRedirect)
				}
				if storedSessionToken != "" {
					data, err := parsers.ParseNYTResult()
					if err != nil {
						log.Println("Couldn't fetch NYT results: ", err.Error())
					}

					err = templates.ExecuteTemplate(w, "welcome", data)
					if err != nil {
						log.Println(err.Error())
						http.Error(w, "Could not serve page", http.StatusNotFound)
					}
				}
			}
		}

	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println("Error while parsing form: ", err.Error())
			http.Redirect(w, r, "", http.StatusPermanentRedirect)
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

			sqlStmt := fmt.Sprintf("select * from users where username = '%s'", username)
			row := db.QueryRow(sqlStmt)
			var storedId int
			var storedUsername, storedPassword string
			err = row.Scan(&storedId, &storedUsername, &storedPassword)
			if err != nil {
				if err == sql.ErrNoRows {
					http.Redirect(w, r, "", http.StatusPermanentRedirect)
				} else {
					log.Println("Error while scanning rows: ", err.Error())
					http.Redirect(w, r, "", http.StatusPermanentRedirect)
				}
			}

			if encryptedPassword == storedPassword {
				sessionToken := uuid.NewV4().String()
				query := fmt.Sprintf("insert into sessions(session_token,username) values('%s', '%s')",
					sessionToken, username)
				_, err := db.Exec(query)
				if err != nil {
					log.Println("Error while inserting session: ", err.Error())
					http.Redirect(w, r, "", http.StatusPermanentRedirect)
				}

				cookie := http.Cookie{
					Name:   "session_token",
					Value:  sessionToken,
					MaxAge: 3600,
				}
				http.SetCookie(w, &cookie)

				data, err := parsers.ParseNYTResult()
				if err != nil {
					log.Println("Couldn't fetch NYT results: ", err.Error())
				}

				err = templates.ExecuteTemplate(w, "welcome", data)
				if err != nil {
					log.Println(err.Error())
					http.Error(w, "Could not serve page", http.StatusNotFound)
				}
			}
		}
	}
}
