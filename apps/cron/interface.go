package cron

const AppName = "cronManager"

type Service interface {
	Run()
}
