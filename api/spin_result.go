package api

import "time"

type SpinResult struct {
	ID        int       `db:"id" json:"id,omitempty"`
	Result    string    `db:"result" json:"result"`
	Timestamp time.Time `db:"time" json:"timestamp,omitempty"`
}
