package mongo

import (
	"testing"
	"github.com/globalsign/mgo/bson"
	"github.com/json-iterator/go/assert"
	"prisma/tms"
)

func TestCoder_DecodeToPositions(t *testing.T) {
	// get mongo session
	get := bson.Raw{
		Data: []byte{23, 4, 0, 0, 7, 95, 105, 100, 0, 91, 85, 199, 2, 214, 128, 5, 115, 162, 107, 101, 254, 2, 114, 101, 103, 105, 115, 116, 114, 121, 95, 105, 100, 0, 33, 0, 0, 0, 49, 57, 57, 50, 49, 97, 48, 98, 101, 102, 49, 49, 50, 51, 53, 54, 100, 50, 57, 49, 54, 97, 97, 97, 49, 100, 100, 48, 54, 50, 99, 97, 0, 3, 116, 103, 116, 0, 120, 3, 0, 0, 1, 99, 111, 117, 114, 115, 101, 0, 0, 0, 0, 0, 0, 232, 116, 64, 1, 104, 101, 97, 100, 105, 110, 103, 0, 0, 0, 0, 0, 0, 64, 81, 64, 3, 105, 100, 0, 111, 0, 0, 0, 3, 112, 114, 111, 100, 117, 99, 101, 114, 0, 32, 0, 0, 0, 18, 101, 105, 100, 0, 240, 210, 0, 0, 0, 0, 0, 0, 18, 115, 105, 116, 101, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 3, 116, 105, 109, 101, 83, 101, 114, 105, 97, 108, 0, 52, 0, 0, 0, 3, 116, 105, 109, 101, 83, 101, 114, 105, 97, 108, 0, 35, 0, 0, 0, 16, 99, 111, 117, 110, 116, 101, 114, 0, 5, 0, 0, 0, 18, 115, 101, 99, 111, 110, 100, 115, 0, 2, 199, 85, 91, 0, 0, 0, 0, 0, 0, 0, 3, 105, 110, 103, 101, 115, 116, 95, 116, 105, 109, 101, 0, 33, 0, 0, 0, 16, 110, 97, 110, 111, 115, 0, 150, 171, 211, 6, 18, 115, 101, 99, 111, 110, 100, 115, 0, 2, 199, 85, 91, 0, 0, 0, 0, 0, 3, 110, 109, 101, 97, 0, 211, 1, 0, 0, 2, 111, 114, 105, 103, 105, 110, 97, 108, 95, 115, 116, 114, 105, 110, 103, 0, 48, 0, 0, 0, 33, 66, 83, 86, 68, 77, 44, 49, 44, 49, 44, 44, 65, 44, 51, 56, 74, 82, 48, 62, 49, 48, 48, 49, 55, 76, 119, 52, 54, 48, 103, 100, 57, 77, 52, 66, 58, 98, 48, 48, 48, 48, 44, 48, 42, 52, 53, 0, 2, 115, 111, 115, 0, 2, 0, 0, 0, 33, 0, 2, 116, 97, 108, 107, 101, 114, 0, 3, 0, 0, 0, 66, 83, 0, 3, 118, 100, 109, 0, 106, 1, 0, 0, 2, 99, 104, 97, 110, 110, 101, 108, 0, 2, 0, 0, 0, 65, 0, 3, 109, 49, 51, 55, 49, 0, 31, 1, 0, 0, 18, 109, 101, 115, 115, 97, 103, 101, 95, 105, 100, 0, 3, 0, 0, 0, 0, 0, 0, 0, 18, 109, 109, 115, 105, 0, 56, 128, 168, 33, 0, 0, 0, 0, 3, 112, 111, 115, 0, 243, 0, 0, 0, 4, 99, 111, 109, 109, 95, 115, 116, 97, 116, 101, 0, 12, 0, 0, 0, 16, 48, 0, 0, 0, 0, 0, 0, 18, 99, 111, 117, 114, 115, 101, 95, 111, 118, 101, 114, 95, 103, 114, 111, 117, 110, 100, 0, 17, 13, 0, 0, 0, 0, 0, 0, 1, 108, 97, 116, 105, 116, 117, 100, 101, 0, 0, 0, 0, 0, 74, 216, 39, 65, 1, 108, 111, 110, 103, 105, 116, 117, 100, 101, 0, 0, 0, 0, 24, 196, 207, 141, 65, 18, 110, 97, 118, 105, 103, 97, 116, 105, 111, 110, 97, 108, 95, 115, 116, 97, 116, 117, 115, 0, 1, 0, 0, 0, 0, 0, 0, 0, 8, 112, 111, 115, 105, 116, 105, 111, 110, 95, 97, 99, 99, 117, 114, 97, 99, 121, 0, 0, 8, 114, 97, 105, 109, 95, 102, 108, 97, 103, 0, 0, 4, 115, 112, 97, 114, 101, 0, 12, 0, 0, 0, 16, 48, 0, 0, 0, 0, 0, 0, 18, 115, 112, 101, 101, 100, 95, 111, 118, 101, 114, 95, 103, 114, 111, 117, 110, 100, 0, 1, 0, 0, 0, 0, 0, 0, 0, 18, 116, 105, 109, 101, 95, 115, 116, 97, 109, 112, 0, 21, 0, 0, 0, 0, 0, 0, 0, 18, 116, 114, 117, 101, 95, 104, 101, 97, 100, 105, 110, 103, 0, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 18, 115, 101, 110, 116, 101, 110, 99, 101, 95, 99, 111, 117, 110, 116, 0, 1, 0, 0, 0, 0, 0, 0, 0, 18, 115, 101, 110, 116, 101, 110, 99, 101, 95, 105, 110, 100, 101, 120, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 112, 111, 115, 105, 116, 105, 111, 110, 0, 61, 0, 0, 0, 2, 116, 121, 112, 101, 0, 6, 0, 0, 0, 80, 111, 105, 110, 116, 0, 4, 99, 111, 111, 114, 100, 105, 110, 97, 116, 101, 115, 0, 27, 0, 0, 0, 1, 48, 0, 171, 91, 61, 39, 189, 12, 90, 64, 1, 49, 0, 169, 254, 46, 89, 2, 214, 244, 63, 0, 0, 1, 114, 97, 116, 101, 95, 111, 102, 95, 116, 117, 114, 110, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 114, 101, 112, 101, 97, 116, 0, 0, 1, 115, 112, 101, 101, 100, 0, 154, 153, 153, 153, 153, 153, 185, 63, 3, 116, 105, 109, 101, 0, 33, 0, 0, 0, 16, 110, 97, 110, 111, 115, 0, 43, 169, 211, 6, 18, 115, 101, 99, 111, 110, 100, 115, 0, 2, 199, 85, 91, 0, 0, 0, 0, 0, 2, 116, 121, 112, 101, 0, 4, 0, 0, 0, 65, 73, 83, 0, 3, 117, 112, 100, 97, 116, 101, 95, 116, 105, 109, 101, 0, 33, 0, 0, 0, 16, 110, 97, 110, 111, 115, 0, 43, 169, 211, 6, 18, 115, 101, 99, 111, 110, 100, 115, 0, 2, 199, 85, 91, 0, 0, 0, 0, 0, 0, 9, 116, 105, 109, 101, 0, 66, 96, 17, 199, 100, 1, 0, 0, 2, 116, 114, 97, 99, 107, 95, 105, 100, 0, 33, 0, 0, 0, 52, 53, 49, 101, 98, 97, 98, 55, 98, 102, 97, 56, 54, 53, 55, 53, 56, 101, 53, 97, 52, 48, 100, 52, 54, 99, 57, 51, 101, 50, 55, 50, 0, 9, 117, 112, 100, 97, 116, 101, 95, 116, 105, 109, 101, 0, 66, 96, 17, 199, 100, 1, 0, 0, 0},
	}
	track := new(DBTrack)
	Decode(&track, get)
	assert.Equal(t, &tms.Point{
		Latitude:1.3022483333333332,
		Longitude:104.199045,
	}, track.Target.Position)
}
