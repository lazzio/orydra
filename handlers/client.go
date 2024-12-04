package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"orydra/config"
	"orydra/models"
	"orydra/pkg/dao"
	"orydra/pkg/logger"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetClients(w http.ResponseWriter, r *http.Request) {
	envVars := config.SetEnv()

	var clients []models.Client
	// Get clients from database
	err := dao.PgDb.Select("id", "client_name").Table(envVars.POSTGRES_CLIENT_TABLE).Find(&clients).Error
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

func GetClientByID(w http.ResponseWriter, r *http.Request) {
	envVars := config.SetEnv()

	// Get client ID from URL
	clientID := chi.URLParam(r, "id")
	if clientID == "" {
		http.Error(w, "Client ID manquant", http.StatusBadRequest)
		return
	}

	// Get client from database
	var client models.Client
	err := dao.PgDb.Table(envVars.POSTGRES_CLIENT_TABLE).Where("id = ?", clientID).First(&client).Error
	if err != nil {
		http.Error(w, "Client non trouvé", http.StatusNotFound)
		return
	}

	// Generate HTML details of the client
	formHTML := fmt.Sprintf(`
		<h2 class="subtitle">Détails du client</h2>
		<form id="clientForm" hx-post="/api/client/%s/update" hx-trigger="submit">
			<input type="hidden" name="clientId" value="%s">
	`, clientID, clientID)

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

		// Handle int32 fields
		if field.Type.Kind() == reflect.Int32 {
			formHTML += fmt.Sprintf(`
				<div class="field">
					<label class="checkbox"><p><strong>%s</strong> (%s)</p></label>
					<div class="control">
						<input id="%s" name="%s" class="input" type="text" value="%d">
					</div>
				</div>
				`, field.Name, field.Type.String(), field.Name, field.Name, value.Int())
		}
	}

	// Add a submit button to the form that call the UpdateClient function
	// Add a cancel button to the form that redirect to the index page
	formHTML += `<div class="field is-grouped">`
	formHTML += `<p class="control">`
	formHTML += `<button class="button is-primary is-rounded" type="submit">Update</button>`
	formHTML += `</p>`
	formHTML += `<p class="control">`
	formHTML += `<a class="button is-danger is-rounded" href="/">Cancel</a>`
	formHTML += `</p>`
	formHTML += `</div>`
	formHTML += `</form>`

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(formHTML))
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	envVars := config.SetEnv()

	// Get client ID from URL
	clientID := chi.URLParam(r, "id")
	if clientID == "" {
		logger.Logger.Error("Client ID missing", "error", "No client ID provided")
		http.Error(w, "Client ID missing", http.StatusBadRequest)
		return
	}

	// Get existing client
	var client models.Client
	if err := dao.PgDb.Table(envVars.POSTGRES_CLIENT_TABLE).Where("id = ?", clientID).First(&client).Error; err != nil {
		logger.Logger.Error("Client not found", "error", err)
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	// Parse the form
	if err := r.ParseForm(); err != nil {
		logger.Logger.Error("Error parsing the form", "error", err)
		http.Error(w, "Error processing the form", http.StatusBadRequest)
		return
	}

	// Update the client fields with the values from the form
	clientType := reflect.TypeOf(client)
	clientValue := reflect.ValueOf(&client).Elem()

	for i := 0; i < clientType.NumField(); i++ {
		field := clientType.Field(i)
		formValue := r.FormValue(field.Name)

		if formValue != "" {
			fieldValue := clientValue.Field(i)

			switch field.Type.Kind() {
			case reflect.String:
				fieldValue.SetString(formValue)
			case reflect.Bool:
				fieldValue.SetBool(formValue == "on" || formValue == "true")
			case reflect.Slice:
				if field.Type.Elem().Kind() == reflect.String {
					values := strings.Split(formValue, ",")
					fieldValue.Set(reflect.ValueOf(values))
				}
			case reflect.Int32:
				intValue, err := strconv.ParseInt(formValue, 10, 32)
				if err != nil {
					logger.Logger.Error("Error parsing int32", "error", err)
				}
				fieldValue.SetInt(intValue)
			}

			// Manage time.Time fields
			if field.Type == reflect.TypeOf(time.Time{}) {
				timeValue, err := time.Parse(time.RFC3339, formValue)
				if err != nil {
					logger.Logger.Error("Error parsing time", "error", err)
				}
				fieldValue.Set(reflect.ValueOf(timeValue))
			}

			// Manage uuid.UUID fields
			if field.Type == reflect.TypeOf(uuid.UUID{}) {
				uuidValue, err := uuid.Parse(formValue)
				if err != nil {
					logger.Logger.Error("Error parsing UUID", "error", err)
				}
				fieldValue.Set(reflect.ValueOf(uuidValue))
			}

			// Manage sql.NullInt64 fields
			if field.Type == reflect.TypeOf(sql.NullInt64{}) {
				nullIntValue, err := strconv.ParseInt(formValue, 10, 64)
				if err != nil {
					logger.Logger.Error("Error parsing SQL Null Int64", "error", err)
					continue
				}
				// Create a new sql.NullInt64 structure
				newNullInt := sql.NullInt64{
					Int64: nullIntValue,
					Valid: true,
				}
				fieldValue.Set(reflect.ValueOf(newNullInt))
			}

			// Manage []byte fields
			if field.Type == reflect.TypeOf([]byte{}) {
				// Convert the string to a slice
				values := strings.Split(formValue, ",")
				// Serialize to JSON
				jsonData, err := json.Marshal(values)
				if err != nil {
					logger.Logger.Error("Error marshaling JSON", "error", err)
					continue
				}
				fieldValue.SetBytes(jsonData)
			}
		}
	}

	// Save the changes to the database
	if err := dao.PgDb.Table(envVars.POSTGRES_CLIENT_TABLE).Save(&client).Error; err != nil {
		logger.Logger.Error("Error updating the client", "error", err)
		http.Error(w, "Error updating the client", http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(fmt.Sprintf(`<div class="notification is-danger">Error updating the client %s with ID %s: %s</div>`, client.ClientName, clientID, err)))
		return
	}

	// Redirect to the index page and display a success message
	http.Redirect(w, r, "/", http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`<div class="notification is-success">Client %s with ID %s updated successfully</div>`, client.ClientName, clientID)))
}
