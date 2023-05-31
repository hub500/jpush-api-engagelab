package jpushclient

type Option struct {
	TimeLive        int    `json:"time_to_live,omitempty"`
	OverrideMsgId   int64  `json:"override_msg_id,omitempty"`
	ApnsProduction  bool   `json:"apns_production"`
	ApnsCollapseId  string `json:"apns_collapse_id,omitempty"`
	BigPushDuration int    `json:"big_push_duration,omitempty"`
}

func (this *Option) SetApnsCollapseId(id string) {
	this.ApnsCollapseId = id
}

func (this *Option) SetTimelive(timelive int) {
	this.TimeLive = timelive
}

func (this *Option) SetOverrideMsgId(id int64) {
	this.OverrideMsgId = id
}

func (this *Option) SetApns(apns bool) {
	this.ApnsProduction = apns
}

func (this *Option) SetBigPushDuration(bigPushDuration int) {
	this.BigPushDuration = bigPushDuration
}
