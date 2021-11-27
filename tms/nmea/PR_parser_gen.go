// Code generated by parse_nmea; DO NOT EDIT
package nmea

import "fmt"
import "strings"
import "strconv"

const (
	// PrefixPR prefix
	PrefixPR = "PR"
)

// PR represents fix data.
type CorePR struct {
	UnusedBackwardCompatibilityFieldNumber1 uint32

	UnusedBackwardCompatibilityFieldNumber1Validity bool

	UnusedBackwardCompatibilityFieldNumber2 uint32

	UnusedBackwardCompatibilityFieldNumber2Validity bool

	CurrentProfileNumber uint32

	CurrentProfileNumberValidity bool

	UnusedBackwardCompatibilityFieldNumber4 uint32

	UnusedBackwardCompatibilityFieldNumber4Validity bool
}

type PR struct {
	BaseSentence
	CorePR
}

func NewPR(sentence BaseSentence) *PR {
	s := new(PR)
	s.BaseSentence = sentence

	s.UnusedBackwardCompatibilityFieldNumber1Validity = false

	s.UnusedBackwardCompatibilityFieldNumber2Validity = false

	s.CurrentProfileNumberValidity = false

	s.UnusedBackwardCompatibilityFieldNumber4Validity = false

	return s
}

func (s *PR) parse() error {
	var err error

	if s.Format != PrefixPR {
		err = fmt.Errorf("%s is not a %s", s.Format, PrefixPR)
		return err
	}

	if len(s.Fields) == 0 {
		return nil
	} else {
		if s.Fields[0] != "" {
			i, err := strconv.ParseUint(s.Fields[0], 10, 32)
			if err != nil {
				return fmt.Errorf("PR decode variation error: %s", s.Fields[0])
			} else {
				s.CorePR.UnusedBackwardCompatibilityFieldNumber1 = uint32(i)
				s.CorePR.UnusedBackwardCompatibilityFieldNumber1Validity = true
			}

		}
	}

	if len(s.Fields) == 1 {
		return nil
	} else {
		if s.Fields[1] != "" {
			i, err := strconv.ParseUint(s.Fields[1], 10, 32)
			if err != nil {
				return fmt.Errorf("PR decode variation error: %s", s.Fields[1])
			} else {
				s.CorePR.UnusedBackwardCompatibilityFieldNumber2 = uint32(i)
				s.CorePR.UnusedBackwardCompatibilityFieldNumber2Validity = true
			}

		}
	}

	if len(s.Fields) == 2 {
		return nil
	} else {
		if s.Fields[2] != "" {
			i, err := strconv.ParseUint(s.Fields[2], 10, 32)
			if err != nil {
				return fmt.Errorf("PR decode variation error: %s", s.Fields[2])
			} else {
				s.CorePR.CurrentProfileNumber = uint32(i)
				s.CorePR.CurrentProfileNumberValidity = true
			}

		}
	}

	if len(s.Fields) == 3 {
		return nil
	} else {
		if s.Fields[3] != "" {
			i, err := strconv.ParseUint(s.Fields[3], 10, 32)
			if err != nil {
				return fmt.Errorf("PR decode variation error: %s", s.Fields[3])
			} else {
				s.CorePR.UnusedBackwardCompatibilityFieldNumber4 = uint32(i)
				s.CorePR.UnusedBackwardCompatibilityFieldNumber4Validity = true
			}

		}
	}

	return nil
}

func (s *PR) Encode() (string, error) {
	var Raw string

	if s.Format != PrefixPR {
		err := fmt.Errorf("Sentence format %s is not a PR sentence", s.Format)
		return "", err
	}

	Raw = s.SOS + s.Talker + s.Format

	if s.UnusedBackwardCompatibilityFieldNumber1Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CorePR.UnusedBackwardCompatibilityFieldNumber1), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CorePR.UnusedBackwardCompatibilityFieldNumber1), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.UnusedBackwardCompatibilityFieldNumber2Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CorePR.UnusedBackwardCompatibilityFieldNumber2), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CorePR.UnusedBackwardCompatibilityFieldNumber2), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.CurrentProfileNumberValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CorePR.CurrentProfileNumber), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CorePR.CurrentProfileNumber), 10)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.UnusedBackwardCompatibilityFieldNumber4Validity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + "," + strconv.FormatUint(uint64(s.CorePR.UnusedBackwardCompatibilityFieldNumber4), 10)

		} else {
			Raw = Raw + "," + strconv.FormatUint(uint64(s.CorePR.UnusedBackwardCompatibilityFieldNumber4), 10)
		}

	}

	check := Checksum(Raw)

	Raw = Raw + check

	return Raw, nil

}
