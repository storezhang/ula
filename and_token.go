package ula

import (
	`time`
)

type andToken struct {
	accessToken string
	expiresIn   time.Time
}

func (at *andToken) validate() (token string, validate bool) {
	return at.accessToken, time.Now().After(at.expiresIn.Add(5 * time.Minute))
}
