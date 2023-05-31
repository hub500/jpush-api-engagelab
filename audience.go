package jpushclient

const (
	TAG     = "tag"
	TAG_AND = "tag_and"
	ALIAS   = "alias"
	ID      = "registration_id"
)

type Audience struct {
	Object   interface{}
	audience map[string][]string
}

func (this *Audience) All() {
	this.Object = "all"
}

// 设备标识 一次推送最多 1000 个
func (this *Audience) SetID(ids []string) {
	_ids := []string{}
	if len(ids) > 1000 {
		_ids = append(_ids, ids[:1000]...)
	} else {
		_ids = ids
	}
	this.set(ID, _ids)
}

func (this *Audience) SetTag(tags []string) {
	this.set(TAG, tags)
}

func (this *Audience) SetTagAnd(tags []string) {
	this.set(TAG_AND, tags)
}

func (this *Audience) SetAlias(alias []string) {
	this.set(ALIAS, alias)
}

func (this *Audience) set(key string, v []string) {
	if this.audience == nil {
		this.audience = make(map[string][]string)
		this.Object = this.audience
	}

	//v, ok = this.audience[key]
	//if ok {
	//	return
	//}
	this.audience[key] = v
}
