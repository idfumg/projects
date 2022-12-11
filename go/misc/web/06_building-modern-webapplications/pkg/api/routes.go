package api

import (
	"myapp/pkg/config"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(appConfig *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	if appConfig.InProduction {
		mux.Use(LoadCSRF(appConfig))
	}
	mux.Use(LoadSession(appConfig))

	mux.Get("/", Home(appConfig))
	mux.Get("/about", About(appConfig))
	mux.Get("/contact", Contact(appConfig))
	mux.Get("/room1", Room1(appConfig))
	mux.Get("/room2", Room2(appConfig))
	mux.Get("/reservation", Reservation(appConfig))
	mux.Post("/reservation", PostReservation(appConfig))
	mux.Get("/reservation-summary", ReservationSummary(appConfig))
	mux.Get("/availability", Availability(appConfig))
	mux.Post("/availability", PostAvailability(appConfig))
	mux.Post("/availability-json", AvailabilityJson(appConfig))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
