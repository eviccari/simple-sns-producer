package handlers

import (
	"context"
	"encoding/json"
	"github.com/eviccari/rest-http-utils/httperrors"
	"github.com/eviccari/rest-http-utils/httputils"
	"github.com/eviccari/simple-sns-producer/adapters"
	"github.com/eviccari/simple-sns-producer/dtos"
	"github.com/eviccari/simple-sns-producer/services"
	"net/http"
	"time"
)

type PublisherHandler struct {
	ctx    context.Context
	ms     services.MessageService
	logger adapters.LoggerAdapter
}

// NewPublisherHandler - Create new PublisherHandler instance
func NewPublisherHandler(ctx context.Context, ms services.MessageService, logger adapters.LoggerAdapter) *PublisherHandler {
	return &PublisherHandler{
		ctx:    ctx,
		ms:     ms,
		logger: logger,
	}
}

// HandlePOST - Handle POST Request to send message body to topic.
// The topicName must be provided into HTTP Request Header
func (ph *PublisherHandler) HandlePOST(w http.ResponseWriter, r *http.Request) {
	topicName := httputils.GetValueFromHeader(r, "topic-name")
	var message dtos.RequestBodyDTO

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		httputils.WriteJSONErrorResponse(w, httperrors.NewBadRequestError(err))
		ph.logger.Error(err)
		return
	}

	ctxWithCancel, cancel := context.WithTimeout(ph.ctx, time.Second*10)
	defer cancel()

	protocolID, err := ph.ms.Send(ctxWithCancel, message.Message, topicName)
	if err != nil {
		httputils.WriteJSONErrorResponse(w, httperrors.NewInternalServerError(err))
		ph.logger.Error(err)
		return
	}

	httputils.WriteJSONResponse(w, dtos.NewResponseBodyDTO(protocolID), http.StatusOK)
}
