package app

import "time"

func Notifier() {
	config := Get().Config()
	ntime, _ := time.Parse("15:04", config.NTime)
	ctime, _ := time.Parse("15:04", config.CTime)
	var tgm TGMessage
	var trigger bool
	tgm.Keyboard = CompileQueryKeyboard()
	tgm.TGToken = config.TGToken
	tgm.UserID = config.NotificationID
	for {
		now := time.Now()
		currentTime := time.Date(0, 1, 1, now.Hour(), now.Minute(), 0, 0, time.UTC)
		if currentTime.Equal(ntime) && !trigger {
			tgm.Text = config.NotificationText
			trigger = true
			SendTGMessage(tgm, true)
		} else if currentTime.Equal(ctime) && !trigger {
			tgm.Text = config.ConfirmationText
			trigger = true
			SendTGMessage(tgm, false)
		} else {
			trigger = false
		}
		time.Sleep(TIMEOUT * time.Second)
	}
}
