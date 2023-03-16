package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	//Cresate a file server
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handler(http.MethodGet, "/resource/*filepath", http.StripPrefix("/resource", fileServer))

	router.HandlerFunc(http.MethodGet, "/affidavit", app.AffidavitForm)
	router.HandlerFunc(http.MethodGet, "/", app.login)
	router.HandlerFunc(http.MethodGet, "/signup", app.signup)
	router.HandlerFunc(http.MethodGet, "/dashboard", app.dashboard)
	router.HandlerFunc(http.MethodGet, "/TableView", app.TableView)
	router.HandlerFunc(http.MethodGet, "/ArchiveView", app.ArchiveView)
	router.HandlerFunc(http.MethodPost, "/affidavit", app.SubmittedAffidavit)
	router.HandlerFunc(http.MethodGet, "/affidavitAccept", app.affidavitAccept)
	router.HandlerFunc(http.MethodGet, "/affidavitDeny", app.affidavitDeny)
	router.HandlerFunc(http.MethodGet, "/ArchiveForm", app.ArchiveForm)
	router.HandlerFunc(http.MethodGet, "/UnarchiveForm", app.UnarchiveForm)
	router.HandlerFunc(http.MethodGet, "/DownloadPDF", app.DownloadPDF)
	router.HandlerFunc(http.MethodGet, "/RestoreForm", app.RestoreForm)

	return router
}
