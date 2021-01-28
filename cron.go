package main

import (
	"github.com/robfig/cron"
	"go-admin/api/models"
	"time"
)

func main() {


	c:=cron.New()

	c.AddFunc("* * * * * *", func() {
		models.DeleteRole()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <- t1.C:
			t1.Reset(time.Second*10)
		}
	}
}