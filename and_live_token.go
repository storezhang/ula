package ula

import (
	`time`
)

type andLiveToken struct {
	accessToken string
	expiresIn   time.Time
}

func (at *andLiveToken) validate() (token string, validate bool) {
	return at.accessToken, time.Now().After(at.expiresIn.Add(5 * time.Minute))
}
