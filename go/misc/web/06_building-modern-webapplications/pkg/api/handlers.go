package api

import (
	"encoding/json"
	"log"
	"myapp/pkg/config"
	"myapp/pkg/forms"
	"myapp/pkg/models"
	"myapp/pkg/render"
	"net/http"
)

func renderWithFatal(appConfig *config.AppConfig, w http.ResponseWriter, r *http.Request, page string, td *render.TemplateData) {
	if err := render.RenderTemplate(appConfig, w, r, page, td); err != nil {
		appConfig.InfoLog.Fatalln(err)
	}
}

func Home(appConfig *config.AppConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		remoteIP := r.RemoteAddr
		appConfig.Session.Put(r.Context(), "remote_ip", remoteIP)
		renderWithFatal(appConfig, w, r, "home.page.gohtml", &render.TemplateData{})
	}
}

func About(appConfig *config.AppConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		stringMap := map[string]string{}
		stringMap["test"] = "Hello, again."

		remoteIP := appConfig.Session.GetString(r.Context(), "remote_ip")
		stringMap["remote_ip"] = remoteIP

		renderWithFatal(appConfig, w, r, "about.page.gohtml", &render.TemplateData{StringMap: stringMap})
	}
}

func Contact(appConfig *config.AppConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		renderWithFatal(appConfig, w, r, "contact.page.gohtml", &render.TemplateData{})
	}
}

func Room1(appConfig *config.AppConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		renderWithFatal(appConfig, w, r, "room1.page.gohtml", &render.TemplateData{})
	}
}

func Room2(appConfig *config.AppConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		renderWithFatal(appConfig, w, r, "room2.page.gohtml", &render.TemplateData{})
	}
}

func Reservation(appConfig *config.AppConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var emptyReservation models.Reservation
		data := map[string]interface{}{}
		data["reservation"] = emptyReservation
		renderWithFatal(appConfig, w, r, "reservation.page.gohtml", &render.TemplateData{
			Data: data,
			Form: forms.New(nil),
		})
	}
}

func PostReservation(appConfig *config.AppConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			appConfig.InfoLog.Println()
			return
		}
		reservation := models.Reservation{
			FirstName: r.Form.Get("first_name"),
			LastName:  r.Form.Get("last_name"),
			Email:     r.Form.Get("email"),
			Phone:     r.Form.Get("phone"),
		}
		form := forms.New(r.PostForm)

		form.Required("first_name", "last_name", "email")
		form.MinLength("first_name", 3, r)
		form.MinLength("last_name", 3, r)
		form.IsEmail("email")

		if !form.Valid() {
			data := map[string]interface{}{}
			data["reservation"] = reservation

			renderWithFatal(appConfig, w, r, "reservation.page.gohtml", &render.TemplateData{
				Data: data,
				Form: form,
			})

			return
		}

		appConfig.Session.Put(r.Context(), "reservation", reservation)
		http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
	}
}

func Availability(appConfig *config.AppConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		renderWithFatal(appConfig, w, r, "availability.page.gohtml", &render.TemplateData{})
	}
}

func PostAvailability(appConfig *config.AppConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := r.Form.Get("start")
		end := r.Form.Get("end")
		log.Println("Posted to the search availability", start, end)
		w.Write([]byte("Posted to the search availability"))
	}
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func AvailabilityJson(appConfig *config.AppConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := jsonResponse{
			OK:      true,
			Message: "hi there",
		}
		output, err := json.MarshalIndent(&resp, "", "    ")
		if err != nil {
			appConfig.InfoLog.Fatalln(err)
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err = w.Write(output); err != nil {
			appConfig.InfoLog.Fatalln(err)
		}
	}
}

func ReservationSummary(appConfig *config.AppConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reservation, ok := appConfig.Session.Get(r.Context(), "reservation").(models.Reservation)
		if !ok {
			appConfig.InfoLog.Println("Can not get an item from a session")
			appConfig.Session.Put(r.Context(), "error", "Can't get a reservation from a session")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		appConfig.Session.Remove(r.Context(), "reservation")
		data := map[string]interface{}{}
		data["reservation"] = reservation
		renderWithFatal(appConfig, w, r, "reservation-summary.page.gohtml", &render.TemplateData{
			Data: data,
		})
	}
}
