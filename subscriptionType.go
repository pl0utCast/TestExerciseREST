package main

import "time"

type subscription struct {
	Name              string
	Price             float64
	UserId            float64
	SubscriptionStart time.Time
	SubscriptionEnd   time.Time
}
