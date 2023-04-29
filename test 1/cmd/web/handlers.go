package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"text/template"
	"time"

	"github.com/HipolitoBautista/internal/models"
)

// shared data store
var dataStore = struct {
	sync.RWMutex
	data map[string]int64
}{data: make(map[string]int64)}

func (app *application) AffidavitForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("form"))
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "./ui/html/Login.UI.tmpl")
}

func (app *application) UserUI(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "./ui/html/User.UI.tmpl")
}

func (app *application) TableView(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "./ui/html/Table.View.tmpl")
}

func (app *application) ArchiveView(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "./ui/html/Archive.View.tmpl")
}

func (app *application) LoadForm(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "./ui/html/Form.UI.tmpl")

}

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("signup"))
}

func (app *application) SubmittedAffidavit(w http.ResponseWriter, r *http.Request) {

	//Code for INSERT
	w.Write([]byte("submitted"))

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	//get the form data

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return

	}
	//collecting values from form
	name := r.PostForm.Get("full-name")
	other_name := r.PostForm.Get("other-names")
	name_change, _ := strconv.ParseBool(r.PostForm.Get("name-changed"))
	previous_name := r.PostForm.Get("previous-names")
	reason_for_change := r.PostForm.Get("reason-for-change")
	social_num, _ := strconv.ParseInt(r.PostForm.Get("SS-number"), 10, 64)
	social_date, _ := time.Parse("2006-01-02", r.PostForm.Get("SS-issue-date"))
	social_country := r.PostForm.Get("SS-country-value")
	passport_num, _ := strconv.ParseInt(r.PostForm.Get("PP-number"), 10, 64)
	passport_date, _ := time.Parse("2006-01-02", r.PostForm.Get("PP-issue-date"))
	passport_country := r.PostForm.Get("PP-country-value")
	NHI_num, _ := strconv.ParseInt(r.PostForm.Get("NHI-number"), 10, 64)
	NHI_date, _ := time.Parse("2006-01-02", r.PostForm.Get("NHI-issue-date"))
	NHI_country := r.PostForm.Get("NHI-country-value")
	dob, _ := time.Parse("2006-01-02", r.PostForm.Get("dateOfBirth"))
	pob := r.PostForm.Get("placeOfBirth")
	nationality := r.PostForm.Get("nationality")
	acq_nationality := r.PostForm.Get("acq-nationality")
	spouse := r.PostForm.Get("spouse-name")
	address := r.PostForm.Get("AF-address")
	phone, _ := strconv.ParseInt(r.PostForm.Get("AF-number"), 10, 64)
	fax, _ := strconv.ParseInt(r.PostForm.Get("Fax-Number"), 10, 64)
	email := r.PostForm.Get("email-address")
	//create instance of form model
	data := &models.Form{
		Form_status:              "unverified",
		Archive_status:           false,
		Affiant_full_name:        name,
		Other_names:              other_name,
		Name_change_status:       name_change,
		Previous_name:            previous_name,
		Reason_for_Change:        reason_for_change,
		Social_security_num:      social_num,
		Social_security_date:     social_date,
		Social_security_country:  social_country,
		Passport_number:          passport_num,
		Passport_date:            passport_date,
		Passport_country:         passport_country,
		NHI_number:               NHI_num,
		NHI_date:                 NHI_date,
		NHI_country:              NHI_country,
		Dob:                      dob,
		Place_of_birth:           pob,
		Nationality:              nationality,
		Acquired_nationality:     acq_nationality,
		Spouse_name:              spouse,
		Affiants_address:         address,
		Residencial_phone_number: phone,
		Residenceial_fax_num:     fax,
		Residencial_email:        email,
	}
	//inserting data into DB using the insert method
	form_id, err := app.form.Insert(data)
	fmt.Println(form_id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}

func (app *application) AdminLogin(w http.ResponseWriter, r *http.Request) {
	id, err := app.admin.InsertAdmins()
	fmt.Println(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) updateForm(w http.ResponseWriter, r *http.Request) {

	//form ID to load when wanting to edit the form (READING)
	//set the form id manually since it's not integrated with the UI yet
	form_id := 3
	FD, err := app.form.Read(form_id)

	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//an instance of templateData
	data := &templateData{
		Form: FD,
	} //this allows us to pass in mutliple data into the template

	//display quotes using tmpl
	ts, err := template.ParseFiles("./ui/html/form.show.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//getting the form ID
	dataStore.Lock()
	dataStore.data["key"] = int64(form_id)
	dataStore.Unlock()

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}

}

func (app *application) updateFormQuery(w http.ResponseWriter, r *http.Request) {
	//Code to UPDATE a form

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	//get the form data

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return

	}
	//collecting values from form
	//form_id, _ := strconv.ParseInt(r.PostForm.Get("form-id"), 10, 64)
	name := r.PostForm.Get("full-name")
	other_name := r.PostForm.Get("other-names")
	name_change, _ := strconv.ParseBool(r.PostForm.Get("name-changed"))
	previous_name := r.PostForm.Get("previous-names")
	reason_for_change := r.PostForm.Get("reason-for-change")
	social_num, _ := strconv.ParseInt(r.PostForm.Get("SS-number"), 10, 64)
	social_date, _ := time.Parse("2006-01-02", r.PostForm.Get("SS-issue-date"))
	social_country := r.PostForm.Get("SS-country-value")
	passport_num, _ := strconv.ParseInt(r.PostForm.Get("PP-number"), 10, 64)
	passport_date, _ := time.Parse("2006-01-02", r.PostForm.Get("PP-issue-date"))
	passport_country := r.PostForm.Get("PP-country-value")
	NHI_num, _ := strconv.ParseInt(r.PostForm.Get("NHI-number"), 10, 64)
	NHI_date, _ := time.Parse("2006-01-02", r.PostForm.Get("NHI-issue-date"))
	NHI_country := r.PostForm.Get("NHI-country-value")
	dob, _ := time.Parse("2006-01-02", r.PostForm.Get("dateOfBirth"))
	pob := r.PostForm.Get("placeOfBirth")
	nationality := r.PostForm.Get("nationality")
	acq_nationality := r.PostForm.Get("acq-nationality")
	spouse := r.PostForm.Get("spouse-name")
	address := r.PostForm.Get("AF-address")
	phone, _ := strconv.ParseInt(r.PostForm.Get("AF-number"), 10, 64)
	fax, _ := strconv.ParseInt(r.PostForm.Get("Fax-Number"), 10, 64)
	email := r.PostForm.Get("email-address")
	dataStore.RLock()
	form_id := dataStore.data["key"]
	dataStore.RUnlock()

	data := &models.Form{
		Form_id:                  form_id,
		Form_status:              "unverified",
		Archive_status:           false,
		Affiant_full_name:        name,
		Other_names:              other_name,
		Name_change_status:       name_change,
		Previous_name:            previous_name,
		Reason_for_Change:        reason_for_change,
		Social_security_num:      social_num,
		Social_security_date:     social_date,
		Social_security_country:  social_country,
		Passport_number:          passport_num,
		Passport_date:            passport_date,
		Passport_country:         passport_country,
		NHI_number:               NHI_num,
		NHI_date:                 NHI_date,
		NHI_country:              NHI_country,
		Dob:                      dob,
		Place_of_birth:           pob,
		Nationality:              nationality,
		Acquired_nationality:     acq_nationality,
		Spouse_name:              spouse,
		Affiants_address:         address,
		Residencial_phone_number: phone,
		Residenceial_fax_num:     fax,
		Residencial_email:        email,
	}

	Test, err := app.form.Update(data)
	fmt.Println(Test)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
func (app *application) deleteForm(w http.ResponseWriter, r *http.Request) {
	//Code to delete a form
	err := r.ParseForm()

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return

	}
	//collecting values from form
	dataStore.RLock()
	form_id := dataStore.data["key"]
	dataStore.RUnlock()

	app.form.Delete(form_id)

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
