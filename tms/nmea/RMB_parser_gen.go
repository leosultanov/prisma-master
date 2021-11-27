// Code generated by parse_nmea; DO NOT EDIT
package nmea

import "fmt"
import "strings"
import "strconv"

const (
	// PrefixRMB prefix
	PrefixRMB = "RMB"
)

// RMB represents fix data.
type CoreRMB struct {
	Status string

	StatusValidity bool

	CrossTrackError float64

	CrossTrackErrorValidity bool

	DirectionSteer string

	DirectionSteerValidity bool

	ToWaypointID float64

	ToWaypointIDValidity bool

	FromWaypointID float64

	FromWaypointIDValidity bool

	DestinationWaypointLatitude float64

	DestinationWaypointLatitudeValidity bool

	NorS string

	NorSValidity bool

	DestinationWaypointLongitude float64

	DestinationWaypointLongitudeValidity bool

	EorW string

	EorWValidity bool

	RangeToDestination float64

	RangeToDestinationValidity bool

	BearingToDestination float64

	BearingToDestinationValidity bool

	DestinationClosingVelocity float64

	DestinationClosingVelocityValidity bool

	ArrivalStatus string

	ArrivalStatusValidity bool

	Ffa string

	FfaValidity bool
}

type RMB struct {
	BaseSentence
	CoreRMB
}

func NewRMB(sentence BaseSentence) *RMB {
	s := new(RMB)
	s.BaseSentence = sentence

	s.StatusValidity = false

	s.CrossTrackErrorValidity = false

	s.DirectionSteerValidity = false

	s.ToWaypointIDValidity = false

	s.FromWaypointIDValidity = false

	s.DestinationWaypointLatitudeValidity = false

	s.NorSValidity = false

	s.DestinationWaypointLongitudeValidity = false

	s.EorWValidity = false

	s.RangeToDestinationValidity = false

	s.BearingToDestinationValidity = false

	s.DestinationClosingVelocityValidity = false

	s.ArrivalStatusValidity = false

	s.FfaValidity = false

	return s
}

func (s *RMB) parse() error {
	var err error

	if s.Format != PrefixRMB {
		err = fmt.Errorf("%s is not a %s", s.Format, PrefixRMB)
		return err
	}

	if len(s.Fields) == 0 {
		return nil
	} else {
		if s.Fields[0] != "" {
			s.Status = s.Fields[0]
			s.StatusValidity = true
		}
	}

	if len(s.Fields) == 1 {
		return nil
	} else {
		if s.Fields[1] != "" {
			i, err := strconv.ParseFloat(s.Fields[1], 64)
			if err != nil {
				return fmt.Errorf("RMB decode variation error: %s", s.Fields[1])
			} else {
				s.CoreRMB.CrossTrackError = float64(i)
				s.CoreRMB.CrossTrackErrorValidity = true
			}

		}
	}

	if len(s.Fields) == 2 {
		return nil
	} else {
		if s.Fields[2] != "" {
			s.DirectionSteer = s.Fields[2]
			s.DirectionSteerValidity = true
		}
	}

	if len(s.Fields) == 3 {
		return nil
	} else {
		if s.Fields[3] != "" {
			i, err := strconv.ParseFloat(s.Fields[3], 64)
			if err != nil {
				return fmt.Errorf("RMB decode variation error: %s", s.Fields[3])
			} else {
				s.CoreRMB.ToWaypointID = float64(i)
				s.CoreRMB.ToWaypointIDValidity = true
			}

		}
	}

	if len(s.Fields) == 4 {
		return nil
	} else {
		if s.Fields[4] != "" {
			i, err := strconv.ParseFloat(s.Fields[4], 64)
			if err != nil {
				return fmt.Errorf("RMB decode variation error: %s", s.Fields[4])
			} else {
				s.CoreRMB.FromWaypointID = float64(i)
				s.CoreRMB.FromWaypointIDValidity = true
			}

		}
	}

	if len(s.Fields) == 5 {
		return nil
	} else {
		if s.Fields[5] != "" {
			i, err := strconv.ParseFloat(s.Fields[5], 64)
			if err != nil {
				return fmt.Errorf("RMB decode variation error: %s", s.Fields[5])
			} else {
				s.CoreRMB.DestinationWaypointLatitude = float64(i)
				s.CoreRMB.DestinationWaypointLatitudeValidity = true
			}

		}
	}

	if len(s.Fields) == 6 {
		return nil
	} else {
		if s.Fields[6] != "" {
			s.NorS = s.Fields[6]
			s.NorSValidity = true
		}
	}

	if len(s.Fields) == 7 {
		return nil
	} else {
		if s.Fields[7] != "" {
			i, err := strconv.ParseFloat(s.Fields[7], 64)
			if err != nil {
				return fmt.Errorf("RMB decode variation error: %s", s.Fields[7])
			} else {
				s.CoreRMB.DestinationWaypointLongitude = float64(i)
				s.CoreRMB.DestinationWaypointLongitudeValidity = true
			}

		}
	}

	if len(s.Fields) == 8 {
		return nil
	} else {
		if s.Fields[8] != "" {
			s.EorW = s.Fields[8]
			s.EorWValidity = true
		}
	}

	if len(s.Fields) == 9 {
		return nil
	} else {
		if s.Fields[9] != "" {
			i, err := strconv.ParseFloat(s.Fields[9], 64)
			if err != nil {
				return fmt.Errorf("RMB decode variation error: %s", s.Fields[9])
			} else {
				s.CoreRMB.RangeToDestination = float64(i)
				s.CoreRMB.RangeToDestinationValidity = true
			}

		}
	}

	if len(s.Fields) == 10 {
		return nil
	} else {
		if s.Fields[10] != "" {
			i, err := strconv.ParseFloat(s.Fields[10], 64)
			if err != nil {
				return fmt.Errorf("RMB decode variation error: %s", s.Fields[10])
			} else {
				s.CoreRMB.BearingToDestination = float64(i)
				s.CoreRMB.BearingToDestinationValidity = true
			}

		}
	}

	if len(s.Fields) == 11 {
		return nil
	} else {
		if s.Fields[11] != "" {
			i, err := strconv.ParseFloat(s.Fields[11], 64)
			if err != nil {
				return fmt.Errorf("RMB decode variation error: %s", s.Fields[11])
			} else {
				s.CoreRMB.DestinationClosingVelocity = float64(i)
				s.CoreRMB.DestinationClosingVelocityValidity = true
			}

		}
	}

	if len(s.Fields) == 12 {
		return nil
	} else {
		if s.Fields[12] != "" {
			s.ArrivalStatus = s.Fields[12]
			s.ArrivalStatusValidity = true
		}
	}

	if len(s.Fields) == 13 {
		return nil
	} else {
		if s.Fields[13] != "" {
			s.Ffa = s.Fields[13]
			s.FfaValidity = true
		}
	}

	return nil
}

func (s *RMB) Encode() (string, error) {
	var Raw string

	if s.Format != PrefixRMB {
		err := fmt.Errorf("Sentence format %s is not a RMB sentence", s.Format)
		return "", err
	}

	Raw = s.SOS + s.Talker + s.Format

	if s.StatusValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreRMB.Status

		} else {
			Raw = Raw + "," + s.CoreRMB.Status
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.CrossTrackErrorValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatFloat(s.CoreRMB.CrossTrackError, 'f', -1, 64)

		} else {
			Raw = Raw + "," + strconv.FormatFloat(s.CoreRMB.CrossTrackError, 'f', -1, 64)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.DirectionSteerValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreRMB.DirectionSteer

		} else {
			Raw = Raw + "," + s.CoreRMB.DirectionSteer
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ToWaypointIDValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatFloat(s.CoreRMB.ToWaypointID, 'f', -1, 64)

		} else {
			Raw = Raw + "," + strconv.FormatFloat(s.CoreRMB.ToWaypointID, 'f', -1, 64)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.FromWaypointIDValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatFloat(s.CoreRMB.FromWaypointID, 'f', -1, 64)

		} else {
			Raw = Raw + "," + strconv.FormatFloat(s.CoreRMB.FromWaypointID, 'f', -1, 64)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.DestinationWaypointLatitudeValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatFloat(s.CoreRMB.DestinationWaypointLatitude, 'f', -1, 64)

		} else {
			Raw = Raw + "," + strconv.FormatFloat(s.CoreRMB.DestinationWaypointLatitude, 'f', -1, 64)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.NorSValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreRMB.NorS

		} else {
			Raw = Raw + "," + s.CoreRMB.NorS
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.DestinationWaypointLongitudeValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatFloat(s.CoreRMB.DestinationWaypointLongitude, 'f', -1, 64)

		} else {
			Raw = Raw + "," + strconv.FormatFloat(s.CoreRMB.DestinationWaypointLongitude, 'f', -1, 64)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.EorWValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreRMB.EorW

		} else {
			Raw = Raw + "," + s.CoreRMB.EorW
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.RangeToDestinationValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatFloat(s.CoreRMB.RangeToDestination, 'f', -1, 64)

		} else {
			Raw = Raw + "," + strconv.FormatFloat(s.CoreRMB.RangeToDestination, 'f', -1, 64)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.BearingToDestinationValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatFloat(s.CoreRMB.BearingToDestination, 'f', -1, 64)

		} else {
			Raw = Raw + "," + strconv.FormatFloat(s.CoreRMB.BearingToDestination, 'f', -1, 64)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.DestinationClosingVelocityValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + strconv.FormatFloat(s.CoreRMB.DestinationClosingVelocity, 'f', -1, 64)

		} else {
			Raw = Raw + "," + strconv.FormatFloat(s.CoreRMB.DestinationClosingVelocity, 'f', -1, 64)
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.ArrivalStatusValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreRMB.ArrivalStatus

		} else {
			Raw = Raw + "," + s.CoreRMB.ArrivalStatus
		}

	} else if len(Raw) > len(strings.TrimSuffix(Raw, ",,")) {
		Raw = Raw + ","
	} else {
		Raw = Raw + ",,"
	}

	if s.FfaValidity == true {

		if len(Raw) > len(strings.TrimSuffix(Raw, ",")) {

			Raw = Raw + s.CoreRMB.Ffa

		} else {
			Raw = Raw + "," + s.CoreRMB.Ffa
		}
	}

	check := Checksum(Raw)

	Raw = Raw + check

	return Raw, nil

}
