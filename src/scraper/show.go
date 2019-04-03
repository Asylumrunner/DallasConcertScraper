package main

import "time"

type show struct {
  headliner string
  venue string
  doors time.Time
  show time.Time
  openers []string
  description string
}
