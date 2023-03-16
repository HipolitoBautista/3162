package main

import (
	"net/http"
)

func (app *application) AffidavitForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("form"))
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("signup"))
}

func (app *application) dashboard(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("dashboard"))
}

func (app *application) TableView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TableView"))
}
func (app *application) ArchiveView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ArchiveView"))
}

func (app *application) SubmittedAffidavit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SubmittedAffidavit"))
}

func (app *application) affidavitAccept(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Affidavit Accept"))
}

func (app *application) affidavitDeny(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Affidavit Deny"))
}

func (app *application) ArchiveForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Archive Form"))
}

func (app *application) UnarchiveForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Unarchive Form"))
}

func (app *application) DownloadPDF(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Download PDF"))
}

func (app *application) RestoreForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Restore Form"))
}
