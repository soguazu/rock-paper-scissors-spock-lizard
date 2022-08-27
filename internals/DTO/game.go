package dto

// GetBasicMessage DTO
type GetBasicMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// PlayRequest dto to play
type PlayRequest struct {
	Choice string `json:"choice"`
}

// PlayResponse dto to play request response
type PlayResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// GetChoices dto to get all choices
type GetChoices struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Choices []string `json:"choices"`
}

// GetChoice dto to get random choice
type GetChoice struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Choice  string `json:"choice"`
}

// Error struct
type Error struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}
