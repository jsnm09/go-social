<<<<<<< HEAD
package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
=======
package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
>>>>>>> faed817c323ba73a1311cc7c20827d731826c10e
