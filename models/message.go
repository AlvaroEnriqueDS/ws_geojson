package models

type Error struct {
        ErrorCode string `json:"error_code"`
        Message string `json:"message"`
        Error string `json:"error,omitempty"`
}
