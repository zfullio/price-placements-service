package entity

type StatusCheck int

const (
	UnknownStatus StatusCheck = iota
	Warning
	Error
	OK
)

type StatusFeed string

const (
	Active   StatusFeed = "активен"
	Inactive StatusFeed = "неактивен"
	Unknown  StatusFeed = "неизвестно"
	Deleted  StatusFeed = "удалено"
)
