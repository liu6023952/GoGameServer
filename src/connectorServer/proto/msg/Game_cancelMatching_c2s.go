package msg

import "connectorServer/proto"
import "bytes"



type Game_cancelMatching_c2s struct {
	MsgId	uint16
	UserId	string

}


func NewGame_cancelMatching_c2s() *Game_cancelMatching_c2s {
	return &Game_cancelMatching_c2s{
		MsgId: 	ID_Game_cancelMatching_c2s,
	}
}


func (this *Game_cancelMatching_c2s) Encode() []byte {
	buf := new(bytes.Buffer)
	proto.SetUint16(buf, this.MsgId)
	proto.SetString(buf, this.UserId)

	return buf.Bytes()
}

func (this *Game_cancelMatching_c2s) Decode(msg []byte) {
	buf := bytes.NewBuffer(msg)
	this.MsgId = proto.GetUint16(buf)
	this.UserId = proto.GetString(buf)

}