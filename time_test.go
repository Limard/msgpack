package msgpack

import (
	"io/ioutil"
	"testing"
	"time"
)

type QueryMessageRes struct {
	Record []Message `msgpack:"record"`
}

type Message struct {
	ReadTime time.Time `msgpack:"readTime"` // 读取时间
}

func Test111(t *testing.T) {
	buf, e := ioutil.ReadFile(`D:\temp\buffer`)
	if e != nil {
		t.Fatal(e)
	}

	res := &QueryMessageRes{}
	e = Unmarshal(buf, res)
	if e != nil {
		t.Fatal(e)
	}
	t.Log(res)
}
