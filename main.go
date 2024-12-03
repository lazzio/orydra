package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"orydra/models"
	"orydra/pkg/dao"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	templates *template.Template
)

func init() {
	// Templates loading
	templates = template.Must(template.ParseFiles("templates/index.html"))
}

func main() {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Get("/", handleIndex)
	r.Get("/api/clients", handleGetClients)
	r.Get("/api/client/{id}", handleGetClientByID)
	r.Post("/api/client/update", handleUpdateClient)

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, "Error rendering the page", http.StatusInternalServerError)
	}
}

func handleGetClients(w http.ResponseWriter, r *http.Request) {
	var clients []models.Client
	// Get clients from database
	err := dao.PgDb.Select("id", "client_name").Table(dao.DbTable).Find(&clients).Error
	if err != nil {
		http.Error(w, "Error fetching clients", http.StatusInternalServerError)
		return
	}

	// Generate HTML options dynamically
	var options string = `<option value="">Select a client</option>`

	for _, client := range clients {
		options += fmt.Sprintf(`<option value="%s">%s</option>`, client.ID, client.ClientName)
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(options))
}

func handleGetClientByID(w http.ResponseWriter, r *http.Request) {
	// Get client ID from URL
	clientID := chi.URLParam(r, "id")
	if clientID == "" {
		http.Error(w, "Client ID manquant", http.StatusBadRequest)
		return
	}

	// Get client from database
	var client models.Client
	err := dao.PgDb.Table(dao.DbTable).Where("id = ?", clientID).First(&client).Error
	if err != nil {
		http.Error(w, "Client non trouvé", http.StatusNotFound)
		return
	}

	// Generate HTML details of the client
	formHTML := `
		<h2 class="subtitle">Détails du client</h2>
		<form hx-post="/api/client/update"
	`

	// For each client field, add an input in the formHTML
	clientType := reflect.TypeOf(client)
	clientValue := reflect.ValueOf(client)

	for i := 0; i < clientType.NumField(); i++ {
		field := clientType.Field(i)
		value := clientValue.Field(i)

		// Handle string fields
		if field.Type.Kind() == reflect.String {
			formHTML += fmt.Sprintf(`
				<div class="field">
					<label class="checkbox"><strong>%s</strong> (%s)</label>
					<div class="control">
						<input id="%s" name="%s" class="input" type="text" value="%s">
					</div>
				</div>
			`, field.Name, field.Type.String(), field.Name, field.Name, value.String())
		}

		// Handle string slices
		if field.Type.Kind() == reflect.Slice && field.Type.Elem().Kind() == reflect.String {
			slice := value.Interface().([]string)
			formHTML += fmt.Sprintf(`
				<div class="field">
					<label class="checkbox"><strong>%s</strong> (%s)</label>
					<div class="control">
						<input id="%s" name="%s" class="input" type="text" value="%s">
					</div>
				</div>
			`, field.Name, field.Type.String(), field.Name, field.Name, strings.Join(slice, ","))
		}

		// Handle time.Time fields
		if field.Type == reflect.TypeOf(time.Time{}) {
			timeValue := value.Interface().(time.Time)
			formHTML += fmt.Sprintf(`
				<div class="field">
					<label class="checkbox"><strong>%s</strong> (%s)</label>
					<div class="control">
						<input id="%s" name="%s" class="input" type="text" value="%s">
					</div>
				</div>
			`, field.Name, field.Type.String(), field.Name, field.Name, timeValue.Format(time.RFC3339))
		}

		// Handle uuid.UUID fields
		if field.Type == reflect.TypeOf(uuid.UUID{}) {
			uuidValue := value.Interface().(uuid.UUID)
			formHTML += fmt.Sprintf(`
				<div class="field">
					<label class="checkbox"><strong>%s</strong> (%s)</label>
					<div class="control">
						<input id="%s" name="%s" class="input" type="text" value="%s">
					</div>
				</div>
			`, field.Name, field.Type.String(), field.Name, field.Name, uuidValue.String())
		}

		// Handle sql.NullInt64 fields
		if field.Type == reflect.TypeOf(sql.NullInt64{}) {
			nullIntValue := value.Interface().(sql.NullInt64)
			formHTML += fmt.Sprintf(`
				<div class="field">
					<label class="checkbox"><strong>%s</strong> (%s)</label>
					<div class="control">
						<input id="%s" name="%s" class="input" type="text" value="%d">
					</div>
				</div>
			`, field.Name, field.Type.String(), field.Name, field.Name, nullIntValue.Int64)
		}

		// Handle []byte fields
		if field.Type == reflect.TypeOf([]byte{}) {
			var data []string
			json.Unmarshal(value.Bytes(), &data)
			formHTML += fmt.Sprintf(`
				<div class="field">
					<label class="checkbox"><strong>%s</strong> (%s)</label>
					<div class="control">
						<input id="%s" name="%s" class="input" type="text" value="%s">
					</div>
				</div>
			`, field.Name, field.Type.String(), field.Name, field.Name, strings.Join(data, ","))
		}

		// Handle bool or *bool fields
		if field.Type == reflect.TypeOf((*bool)(nil)).Elem() || field.Type.Kind() == reflect.Bool {
			checked := ""
			if value.Bool() {
				checked = "checked"
			}
			formHTML += fmt.Sprintf(`
				<div class="field">
					<label class="checkbox"><p><strong>%s</strong> (%s)</p>
						<input id="%s" name="%s" type="checkbox" %s>
					</label>
				</div>
			`, field.Name, field.Type.String(), field.Name, field.Name, checked)
		}
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(formHTML))
}

func handleUpdateClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// _, err := db.Exec(
	// 	"UPDATE hydra_client SET client_name=$2, grant_types=$3, redirect_uris=$4 WHERE id=$1",
	// 	client.ID, client.ClientName, pq.Array(client.GrantTypes), pq.Array(client.RedirectURIs),
	// )
	// if err != nil {
	// 	http.Error(w, "Error updating the client", http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Client updated successfully"))
}
