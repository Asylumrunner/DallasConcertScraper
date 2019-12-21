package main

import "time"

type show struct {
  headliner string
  venue string
  date string
  doors time.Time
  show time.Time
  openers string
  description string
  spotify_url string
}

type ByDate []show

func (a ByDate) Len() int {
  return len(a)
}

func (a ByDate) Less(i, j int) bool {
  return a[i].date < a[j].date
}

func (a ByDate) Swap(i, j int) {
  a[i], a[j] = a[j], a[i]
}
