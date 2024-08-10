package cmd

type User struct {
	ID            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Age           int    `json:"age"`
	RecordingDate int64  `json:"recording_date"`
}
