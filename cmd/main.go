package main

import (
	"context"
	"fmt"
	"github.com/eviccari/simple-sns-producer/adapters"
	"github.com/eviccari/simple-sns-producer/configs"
	"github.com/eviccari/simple-sns-producer/handlers"
	"github.com/eviccari/simple-sns-producer/infra"
	"github.com/eviccari/simple-sns-producer/services"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func main() {
	ctx := context.Background()
	snsMessageProducerAdapter := adapters.NewSNSMessageProducerAdapter(infra.GetAWSSession())
	messageService := services.NewBasicMessageService(snsMessageProducerAdapter)
	logger := adapters.NewBasicLogger()

	r := chi.NewRouter()
	r.Post("/", handlers.NewPublisherHandler(ctx, messageService, logger).HandlePOST)

	logger.Info("App running ;)")
	http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(configs.GetAppPort())), r)
}
