package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/xprazak2/ventrata/internal/controller"
	"github.com/xprazak2/ventrata/internal/repository/memory"
	"github.com/xprazak2/ventrata/internal/service"
	"github.com/xprazak2/ventrata/internal/utils"
)

func main() {
	r := chi.NewRouter()

	repo := memory.NewMemoryRepository()

	repo.SeedData()

	productService := service.NewProductService(repo)
	bookingService := service.NewBookingService(repo)

	dateProvider := utils.NewDateProvider()

	availabilityService := service.NewAvailabilityService(dateProvider, repo)

	r.Use(controller.WithCapability)

	r.Route("/products", func(r chi.Router) {
		r.Get("/", controller.GetProducts(productService))
		r.Get("/{id}", controller.GetProduct(productService))
	})

	r.Post("/availability", controller.GetAvailability(availabilityService))

	r.Route("/bookings", func(r chi.Router) {
		r.Post("/{id}/confirm", controller.ConfirmBooking(bookingService))
		r.Post("/", controller.CreateBooking(bookingService))
		r.Get("/{id}", controller.GetBooking(bookingService))
	})

	port := utils.PortEnvVar()

	sigsChan := make(chan os.Signal, 1)
	signal.Notify(sigsChan, syscall.SIGINT, syscall.SIGTERM)

	server := http.Server{Addr: ":" + port, Handler: r}

	go func() {
		fmt.Printf("Server started at port %s\n", port)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Println(err)
		}
	}()

	<-sigsChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("Failed to gracefully shut down HTTP server: %v", err)
	}

	fmt.Println("Shutting down")
}
