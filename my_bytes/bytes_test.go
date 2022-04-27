package my_bytes

import (
	"bytes"
	logs "github.com/danbai225/go-logs"
	"go-rustdesk-server/model/model_proto"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"testing"
)

var testData = []byte{100, 66, 23, 10, 14, 49, 50, 51, 49, 50, 51, 113, 113, 113, 113, 113, 113, 113, 113, 16, 2, 26, 3, 49, 50, 51}
var testData2 = []byte{122, 82, 10, 8, 52, 50, 50, 50, 55, 51, 56, 51, 18, 36, 70, 49, 57, 49, 57, 57, 57, 52, 45, 66, 52, 51, 65, 45, 53, 57, 55, 48, 45, 66, 55, 68, 53, 45, 54, 53, 65, 66, 53, 53, 49, 68, 53, 56, 50, 69, 26, 32, 162, 208, 255, 250, 164, 248, 203, 65, 39, 191, 175, 216, 113, 41, 196, 148, 131, 230, 60, 101, 77, 17, 79, 148, 236, 119, 201, 242, 238, 162, 22, 217}

func TestDecode(t *testing.T) {
	data, err2 := Decode(testData)
	if err2 != nil {
		logs.Err(err2)
		return
	}
	msg := &model_proto.RendezvousMessage{}
	err := proto.Unmarshal(data, msg)
	if err != nil {
		t.Error(err)
		return
	}
	_, err2 = Encoder(data)
	if err != nil {
		t.Error(err)
		return
	}
}
func TestDecode2(t *testing.T) {
	msg := &model_proto.RendezvousMessage{}
	err := proto.Unmarshal(testData2, msg)
	if err != nil {
		t.Error(err)
		return
	}
}
func TestDecodeHead(t *testing.T) {
	b := bytes.NewBufferString("TEST")
	encoder, _ := Encoder(b.Bytes())
	head, err := DecodeHead(encoder)
	if err != nil {
		t.Error(err)
	}
	if int(head) != len(encoder) {
		t.Error("!=")
	}
	head, err = DecodeHead(testData)
	if err != nil {
		t.Error(err)
	}
	if int(head) != len(testData) {
		t.Error("!=")
	}
}
func TestEncoder(t *testing.T) {
	b := bytes.NewBufferString("")
	for b.Len() < 0x3F {
		b.WriteByte(byte(rand.Int31n(255)))
	}
	cp := make([]byte, b.Len())
	copy(cp, b.Bytes())
	encode, err := Encoder(cp)
	if err != nil {
		t.Error(err)
	}
	data, err := Decode(encode)
	if err != nil {
		t.Error(err)
	}
	if b.String() != string(data) {
		t.Error("!=")
	}
	///-----------
	b = bytes.NewBufferString("")
	for b.Len() < 0x3FF {
		b.WriteByte(byte(rand.Int31n(255)))
	}
	cp = make([]byte, b.Len())
	copy(cp, b.Bytes())
	encode, err = Encoder(cp)
	if err != nil {
		t.Error(err)
	}
	data, err = Decode(encode)
	if err != nil {
		t.Error(err)
	}
	if b.String() != string(data) {
		t.Error("!=")
	}
	///-----------
	b = bytes.NewBufferString("")
	for b.Len() < 0x3FFF {
		b.WriteByte(byte(rand.Int31n(255)))
	}
	cp = make([]byte, b.Len())
	copy(cp, b.Bytes())
	encode, err = Encoder(cp)
	if err != nil {
		t.Error(err)
	}
	data, err = Decode(encode)
	if err != nil {
		t.Error(err)
	}
	if b.String() != string(data) {
		t.Error("!=")
	}
	///-----------
	b = bytes.NewBufferString("")
	for b.Len() < 0x3FFFF {
		b.WriteByte(byte(rand.Int31n(255)))
	}
	cp = make([]byte, b.Len())
	copy(cp, b.Bytes())
	encode, err = Encoder(cp)
	if err != nil {
		t.Error(err)
	}
	data, err = Decode(encode)
	if err != nil {
		t.Error(err)
	}
	if b.String() != string(data) {
		t.Error("!=")
	}
	///-----------
	b = bytes.NewBufferString("")
	for b.Len() < 0x3FFFFF {
		b.WriteByte(byte(rand.Int31n(255)))
	}
	cp = make([]byte, b.Len())
	copy(cp, b.Bytes())
	encode, err = Encoder(cp)
	if err != nil {
		t.Error(err)
	}
	data, err = Decode(encode)
	if err != nil {
		t.Error(err)
	}
	if b.String() != string(data) {
		t.Error("!=")
	}
}
