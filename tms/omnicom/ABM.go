//Ack binary message(0x09), Iridium/3G
// Code generated by omnicom; DO NOT EDIT
package omnicom

import (
	"errors"
	"strconv"
)

type ABM struct {
	Header     byte
	Date       Date
	ID_Msg     uint32
	Error_Type uint32
	Padding    uint32
	CRC        uint32
}

func (ABM *ABM) Parse(input string) error {
	var err error
	if len(input) > 72 {
		err = errors.New("Input message is longer than limit")
		return err
	}
	var count int = 0
	if count+8 > len(input) {
		err = errors.New("Input message length of ABM is shorter than required")
		return err
	}
	count = count + 8
	if count+27 > len(input) {
		err = errors.New("Input message length of ABM is shorter than required")
		return err
	}
	ABM.Date = *new(Date)
	err = ABM.Date.parse(input[count : count+27])
	count = count + 27
	if err != nil {
		return err
	}
	var num uint64
	if count+16 > len(input) {
		err = errors.New("Input message length of ABM is shorter than required")
		return err
	}
	num, err = strconv.ParseUint(input[count:count+16], 2, 32)
	ABM.ID_Msg = uint32(num*1 - 0)
	count = count + 16
	if err != nil {
		return err
	}
	if count+8 > len(input) {
		err = errors.New("Input message length of ABM is shorter than required")
		return err
	}
	num, err = strconv.ParseUint(input[count:count+8], 2, 32)
	ABM.Error_Type = uint32(num*1 - 0)
	count = count + 8
	if err != nil {
		return err
	}
	if count+(8-count%8) > len(input) {
		err = errors.New("Input message length of ABM is shorter than required")
		return err
	}
	num, err = strconv.ParseUint(input[count:count+(8-count%8)], 2, 32)
	ABM.Padding = uint32(num*1 - 0)
	count = count + (8 - count%8)
	if err != nil {
		return err
	}
	if count+8 > len(input) {
		err = errors.New("Input message length of ABM is shorter than required")
		return err
	}
	num, err = strconv.ParseUint(input[count:count+8], 2, 32)
	ABM.CRC = uint32(num*1 - 0)
	count = count + 8
	if err != nil {
		return err
	}
	return err
}
func (ABM *ABM) Encode() ([]byte, error) {
	var str string
	var s string
	var err error
	str += pad(strconv.FormatUint(uint64(ABM.Header), 2), 8)
	s, err = ABM.Date.encode()
	if err != nil {
		return []byte{}, err
	}
	str += s

	s = pad(strconv.FormatUint(uint64((ABM.ID_Msg+0)/1), 2), 16)
	if len(s) > 16 {
		err = errors.New("Value assigned for ABM.ID_Msg exceeds limit")
		return []byte{}, err
	}
	str += s

	s = pad(strconv.FormatUint(uint64((ABM.Error_Type+0)/1), 2), 8)
	if len(s) > 8 {
		err = errors.New("Value assigned for ABM.Error_Type exceeds limit")
		return []byte{}, err
	}
	str += s

	s = pad(strconv.FormatUint(uint64((ABM.Padding+0)/1), 2), 8-len(str)%8)
	if len(s) > (8 - len(str)%8) {
		err = errors.New("Value assigned for ABM.Padding exceeds limit")
		return []byte{}, err
	}
	str += s

	s = pad(strconv.FormatUint(uint64((ABM.CRC+0)/1), 2), 8)
	if len(s) > 8 {
		err = errors.New("Value assigned for ABM.CRC exceeds limit")
		return []byte{}, err
	}
	str += s

	str = attachCRC(str)
	byteList, err := decToByte(str)
	return byteList, err
}
func (ABM *ABM) getHeader() byte {
	return ABM.Header
}
func (ABM *ABM) getCRC() uint32 {
	return ABM.CRC
}
func (ABM *ABM) setCRC(crc uint32) {
	ABM.CRC = crc
}
