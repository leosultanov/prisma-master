// Code generated by parse_nmea; DO NOT EDIT
package nmea

import "fmt"
import "strconv"
import "strings"

// M137116 represents fix data.
type CoreM137116 struct {
	MessageID uint8

	RepeatIndicator uint32

	Mmsi uint32

	Spare1 uint32 //Supposed to be Unknown

	DestinationAID uint32

	OffsetA uint32

	IncrementA uint32

	DestinationBID uint32

	OffsetB uint32

	IncrementB uint32

	Spare2 uint32 //Supposed to be Unknown
}
type M137116 struct {
	VDMO
	CoreM137116
}

func NewM137116(sentence VDMO) *M137116 {
	s := new(M137116)
	s.VDMO = sentence
	return s
}

func (s *M137116) parse() error {
	var err error

	if MessageType(s.EncapData) != 16 {
		err = fmt.Errorf("message %d is not a M137116", MessageType(s.EncapData))
		return err
	}

	data := []byte(s.EncapData)

	//if len(data)*6 > 148 {
	//	err = fmt.Errorf("Message lenght is larger than it should be [%d!=148]", len(data)*6)
	//	return err
	//}

	s.MessageID = MessageType(s.EncapData)

	s.CoreM137116.RepeatIndicator = BitsToInt(6, 7, data)

	s.CoreM137116.Mmsi = BitsToInt(8, 37, data)

	s.CoreM137116.Spare1 = BitsToInt(38, 39, data)

	s.CoreM137116.DestinationAID = BitsToInt(40, 69, data)

	s.CoreM137116.OffsetA = BitsToInt(70, 81, data)

	s.CoreM137116.IncrementA = BitsToInt(82, 91, data)

	s.CoreM137116.DestinationBID = BitsToInt(92, 121, data)

	s.CoreM137116.OffsetB = BitsToInt(122, 133, data)

	s.CoreM137116.IncrementB = BitsToInt(134, 143, data)

	s.CoreM137116.Spare2 = BitsToInt(144, 6*(len(data)-1), data)

	return nil
}

func (s *M137116) Encode() (string, error) {
	var Raw string
	var Sbinary string

	if s.MessageID != 16 {
		err := fmt.Errorf("message %d is not a M137116", s.MessageID)
		return "", err
	}

	Raw = s.SOS + s.Talker + s.Format + ","

	if s.SentenceCountValidity == true {
		Raw = Raw + strconv.FormatInt(int64(s.SentenceCount), 10) + ","
	} else {
		Raw = Raw + ","
	}

	if s.SentenceIndexValidity == true {
		Raw = Raw + strconv.FormatInt(int64(s.SentenceIndex), 10) + ","
	} else {
		Raw = Raw + ","
	}

	if s.SeqMsgIDValidity == true {
		Raw = Raw + strconv.FormatInt(int64(s.SeqMsgID), 10) + ","
	} else {
		Raw = Raw + ","
	}

	if s.ChannelValidity == true {
		Raw = Raw + s.Channel
	}

	str := strconv.FormatInt(int64(s.CoreM137116.MessageID), 2)
	for len(str) < 6 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137116.RepeatIndicator), 2)
	for len(str) < 2 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137116.Mmsi), 2)
	for len(str) < 30 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137116.Spare1), 2)
	for len(str) < 2 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137116.DestinationAID), 2)
	for len(str) < 30 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137116.OffsetA), 2)
	for len(str) < 12 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137116.IncrementA), 2)
	for len(str) < 10 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137116.DestinationBID), 2)
	for len(str) < 30 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137116.OffsetB), 2)
	for len(str) < 12 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137116.IncrementB), 2)
	for len(str) < 10 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	str = strconv.FormatInt(int64(s.CoreM137116.Spare2), 2)
	for len(str) < 4 {
		str = "0" + str
	}
	Sbinary = Sbinary + str

	field := strings.SplitN(Sbinary, "", len(Sbinary))

	var encdata = make([]string, int((len(Sbinary)+int(s.FillBits))/6))

	j := 0
	for i := 0; i < int((len(Sbinary)+int(s.FillBits))/6); i++ {

		if i == (int((len(Sbinary)+int(s.FillBits))/6) - 1) {
			for j < len(Sbinary) {
				encdata[i] = encdata[i] + field[j]
				j = j + 1
			}
			for h := 0; h < int(s.FillBits); h++ {
				encdata[i] = encdata[i] + "0" // fill bits
			}
		} else {
			encdata[i] = field[j] + field[j+1] + field[j+2] + field[j+3] + field[j+4] + field[j+5]
			j = j + 6
		}
	}

	var data string
	for j := 0; j < int((len(Sbinary)+int(s.FillBits))/6); j++ {
		i, _ := strconv.ParseInt(encdata[j], 2, 8)
		if i < 40 {
			i = i + 48
		} else {
			i = i + 8 + 48
		}
		data = data + string(rune(i))
	}

	Raw = Raw + "," + data + ","

	if s.FillBitsValidity == true {
		Raw = Raw + strconv.FormatInt(int64(s.FillBits), 10)
	}

	check := Checksum(Raw)

	Raw = Raw + check

	return Raw, nil

}
