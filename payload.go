package jpushclient

import (
	"encoding/json"
)

type PayLoad struct {
	From     string      `json:"from"`
	Audience interface{} `json:"to"`
	Body     struct {
		Platform     interface{} `json:"platform"`
		Notification interface{} `json:"notification,omitempty"`
		Message      interface{} `json:"message,omitempty"`
		Options      *Option     `json:"options,omitempty"`
	} `json:"body"`
}

func NewPushPayLoad() *PayLoad {
	pl := &PayLoad{
		From: "push",
	}
	o := &Option{
		ApnsProduction: true,
		TimeLive:       86400,
	}
	pl.Body.Options = o
	return pl
}

func (this *PayLoad) SetPlatform(pf *Platform) {
	this.Body.Platform = pf.Os
}

func (this *PayLoad) SetAudience(ad *Audience) {
	this.Audience = ad.Object
}

func (this *PayLoad) SetOptions(o *Option) {
	this.Body.Options = o
}

func (this *PayLoad) SetMessage(m *Message) {
	this.Body.Message = m
}

func (this *PayLoad) SetNotice(notice *Notice) {
	this.Body.Notification = notice
}

func (this *PayLoad) SetApnsProduction(isProd bool) {
	this.Body.Options.ApnsProduction = isProd
}

func (this *PayLoad) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
