package dto

// GetBasicMessage DTO
type GetBasicMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// PlayRequest dto to play
type PlayRequest struct {
	Player int `json:"player"`
}

// PlayResponse dto to play request response
type PlayResponse struct {
	Results string `json:"results"`
	Player  int    `json:"player"`
}

// GetResponse dto to get all choices
type GetResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"choices"`
}

// GetChoices dto
type GetChoices struct {
	Choices interface{} `json:"name"`
	Token   string
}

type GetScoreboard struct {
	PlayerScore   int
	ComputerScore int
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

// GetRandomNumber struct
type GetRandomNumber struct {
	RandomNumber int `json:"random_number"`
}

type Choices struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}
