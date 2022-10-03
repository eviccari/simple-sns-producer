package dtos

// ResponseBodyDTO - Describe ResponseBodyDTO that will be returned by handler
type ResponseBodyDTO struct {
	ProtocolID string `json:"protocol_id"`
}

// NewResponseBodyDTO - Create new ResponseBodyDTO instance
func NewResponseBodyDTO(protocolID string) *ResponseBodyDTO {
	return &ResponseBodyDTO{
		ProtocolID: protocolID,
	}
}
