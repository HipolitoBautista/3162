package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// ROUTES: 10
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.login)
	router.HandlerFunc(http.MethodGet, "/signup", app.signup)
	router.HandlerFunc(http.MethodGet, "/userui", app.UserUI)
	router.HandlerFunc(http.MethodGet, "/AffidavitForm", app.LoadForm)
	router.HandlerFunc(http.MethodPost, "/AffidavitForm", app.SubmittedAffidavit)
	router.HandlerFunc(http.MethodGet, "/TableView", app.TableView)
	router.HandlerFunc(http.MethodGet, "/ArchiveView", app.ArchiveView)
	router.HandlerFunc(http.MethodGet, "/affidavitSubmitted", app.SubmittedAffidavit)
	router.HandlerFunc(http.MethodGet, "/updateForm", app.updateForm)
	router.HandlerFunc(http.MethodPost, "/updateForm", app.updateFormQuery)
	router.HandlerFunc(http.MethodGet, "/deleteForm", app.deleteForm)
	router.HandlerFunc(http.MethodGet, "/affidavit", app.AffidavitForm)
	router.HandlerFunc(http.MethodGet, "/affidavitAccept", app.affidavitAccept)
	router.HandlerFunc(http.MethodGet, "/affidavitDeny", app.affidavitDeny)
	router.HandlerFunc(http.MethodGet, "/ArchiveForm", app.ArchiveForm)
	router.HandlerFunc(http.MethodGet, "/UnarchiveForm", app.UnarchiveForm)
	router.HandlerFunc(http.MethodGet, "/DownloadPDF", app.DownloadPDF)
	router.HandlerFunc(http.MethodGet, "/RestoreForm", app.RestoreForm)
	router.HandlerFunc(http.MethodGet, "/admin", app.AdminLogin)

	return router
}
