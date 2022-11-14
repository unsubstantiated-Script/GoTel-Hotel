package handlers

import (
	"GoTel/internal/config"
	"GoTel/internal/forms"
	"GoTel/internal/models"
	"GoTel/internal/render"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// TemplateData holds data sent from handlers to templates used in render.go Render templates

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	//Grabbing their IP
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Yeet this!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	//Send some data
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and displays form

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		//Giving access to the form object on the page
		Form: forms.New(nil),
	})
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {

}

// Renders Availability Page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// Renders PostAvailability Page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	//Getting Form inputs
	start := r.Form.Get("start_date")
	end := r.Form.Get("end_date")

	//Casting string message to bytes
	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
}

//Setting up a JSON struct with specific types
type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// Renders AvailabilityJSON handles requests for Availability and sends JSON
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	//Creating a json response and error handling
	//assigning both out and err at the same time.
	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	//Header
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
