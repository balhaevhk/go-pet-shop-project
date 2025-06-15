package handlers

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"go-pet-shop/models"
	"log/slog"
	"net/http"
)

type Histories interface {
	GetUserOrderHistory(email string) ([]models.OrderDetail, error)
	GetPopularProducts() ([]models.PopularProduct, error)
}

func GetUserOrderHistory(log *slog.Logger, histories Histories) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "handlers.histories.GetUserOrderHistory"

		log = log.With(
			slog.String("fn", fn),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		email := r.URL.Query().Get("email")

		if email == "" {
			log.Error("missing email parameter")
			http.Error(w, "email is required", http.StatusBadRequest)
			return
		}

		history, err := histories.GetUserOrderHistory(email)
		if err != nil {
			log.Error("failed to get user order history", slog.Any("error", err))
			http.Error(w, "failed to get history", http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, history)
	}
}

func GetPopularProducts(log *slog.Logger, histories Histories) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "handlers.histories.GetPopularProducts"

		log = log.With(
			slog.String("fn", fn),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		products, err := histories.GetPopularProducts()
		if err != nil {
			log.Error("failed to get popular products", slog.Any("error", err))
			http.Error(w, "failed to get popular products", http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, products)
	}
}
