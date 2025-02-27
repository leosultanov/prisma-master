package gjson

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/buger/jsonparser"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson/jlexer"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// TestRandomData is a fuzzing test that throws random data at the Parse
// function looking for panics.
func TestRandomData(t *testing.T) {
	var lstr string
	defer func() {
		if v := recover(); v != nil {
			println("'" + hex.EncodeToString([]byte(lstr)) + "'")
			println("'" + lstr + "'")
			panic(v)
		}
	}()
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 200)
	for i := 0; i < 2000000; i++ {
		n, err := rand.Read(b[:rand.Int()%len(b)])
		if err != nil {
			t.Fatal(err)
		}
		lstr = string(b[:n])
		GetBytes([]byte(lstr), "zzzz")
		Parse(lstr)
	}
}

func TestRandomValidStrings(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 200)
	for i := 0; i < 100000; i++ {
		n, err := rand.Read(b[:rand.Int()%len(b)])
		if err != nil {
			t.Fatal(err)
		}
		sm, err := json.Marshal(string(b[:n]))
		if err != nil {
			t.Fatal(err)
		}
		var su string
		if err := json.Unmarshal([]byte(sm), &su); err != nil {
			t.Fatal(err)
		}
		token := Get(`{"str":`+string(sm)+`}`, "str")
		if token.Type != String || token.Str != su {
			println("["+token.Raw+"]", "["+token.Str+"]", "["+su+"]", "["+string(sm)+"]")
			t.Fatal("string mismatch")
		}
	}
}

func TestEmoji(t *testing.T) {
	const input = `{"utf8":"Example emoji, KO: \ud83d\udd13, \ud83c\udfc3 OK: \u2764\ufe0f "}`
	value := Get(input, "utf8")
	var s string
	json.Unmarshal([]byte(value.Raw), &s)
	if value.String() != s {
		t.Fatalf("expected '%v', got '%v'", s, value.String())
	}
}

func testEscapePath(t *testing.T, json, path, expect string) {
	if Get(json, path).String() != expect {
		t.Fatalf("expected '%v', got '%v'", expect, Get(json, path).String())
	}
}

func TestEscapePath(t *testing.T) {
	json := `{
		"test":{
			"*":"valZ",
			"*v":"val0",
			"keyv*":"val1",
			"key*v":"val2",
			"keyv?":"val3",
			"key?v":"val4",
			"keyv.":"val5",
			"key.v":"val6",
			"keyk*":{"key?":"val7"}
		}
	}`

	testEscapePath(t, json, "test.\\*", "valZ")
	testEscapePath(t, json, "test.\\*v", "val0")
	testEscapePath(t, json, "test.keyv\\*", "val1")
	testEscapePath(t, json, "test.key\\*v", "val2")
	testEscapePath(t, json, "test.keyv\\?", "val3")
	testEscapePath(t, json, "test.key\\?v", "val4")
	testEscapePath(t, json, "test.keyv\\.", "val5")
	testEscapePath(t, json, "test.key\\.v", "val6")
	testEscapePath(t, json, "test.keyk\\*.key\\?", "val7")
}

// this json block is poorly formed on purpose.
var basicJSON = `{"age":100, "name":{"here":"B\\\"R"},
	"noop":{"what is a wren?":"a bird"},
	"happy":true,"immortal":false,
	"items":[1,2,3,{"tags":[1,2,3],"points":[[1,2],[3,4]]},4,5,6,7],
	"arr":["1",2,"3",{"hello":"world"},"4",5],
	"vals":[1,2,3,{"sadf":sdf"asdf"}],"name":{"first":"tom","last":null},
	"created":"2014-05-16T08:28:06.989Z",
	"loggy":{
		"programmers": [
    	    {
    	        "firstName": "Brett",
    	        "lastName": "McLaughlin",
    	        "email": "aaaa",
				"tag": "good"
    	    },
    	    {
    	        "firstName": "Jason",
    	        "lastName": "Hunter",
    	        "email": "bbbb",
				"tag": "bad"
    	    },
    	    {
    	        "firstName": "Elliotte",
    	        "lastName": "Harold",
    	        "email": "cccc",
				"tag":, "good"
    	    },
			{
				"firstName": 1002.3,
				"age": 101
			}
    	]
	},
	"lastly":{"yay":"final"}
}`
var basicJSONB = []byte(basicJSON)

func TestTimeResult(t *testing.T) {
	assert(t, Get(basicJSON, "created").String() == Get(basicJSON, "created").Time().Format(time.RFC3339Nano))
}

func TestParseAny(t *testing.T) {
	assert(t, Parse("100").Float() == 100)
	assert(t, Parse("true").Bool())
	assert(t, Parse("valse").Bool() == false)
}

func TestManyVariousPathCounts(t *testing.T) {
	json := `{"a":"a","b":"b","c":"c"}`
	counts := []int{3, 4, 7, 8, 9, 15, 16, 17, 31, 32, 33, 63, 64, 65, 127, 128, 129, 255, 256, 257, 511, 512, 513}
	paths := []string{"a", "b", "c"}
	expects := []string{"a", "b", "c"}
	for _, count := range counts {
		var gpaths []string
		var gexpects []string
		for i := 0; i < count; i++ {
			if i < len(paths) {
				gpaths = append(gpaths, paths[i])
				gexpects = append(gexpects, expects[i])
			} else {
				gpaths = append(gpaths, fmt.Sprintf("not%d", i))
				gexpects = append(gexpects, "null")
			}
		}
		results := GetMany(json, gpaths...)
		for i := 0; i < len(paths); i++ {
			if results[i].String() != expects[i] {
				t.Fatalf("expected '%v', got '%v'", expects[i], results[i].String())
			}
		}
	}
}
func TestManyRecursion(t *testing.T) {
	var json string
	var path string
	for i := 0; i < 100; i++ {
		json += `{"a":`
		path += ".a"
	}
	json += `"b"`
	for i := 0; i < 100; i++ {
		json += `}`
	}
	path = path[1:]
	assert(t, GetMany(json, path)[0].String() == "b")
}
func TestByteSafety(t *testing.T) {
	jsonb := []byte(`{"name":"Janet","age":38}`)
	mtok := GetBytes(jsonb, "name")
	if mtok.String() != "Janet" {
		t.Fatalf("expected %v, got %v", "Jason", mtok.String())
	}
	mtok2 := GetBytes(jsonb, "age")
	if mtok2.Raw != "38" {
		t.Fatalf("expected %v, got %v", "Jason", mtok2.Raw)
	}
	jsonb[9] = 'T'
	jsonb[12] = 'd'
	jsonb[13] = 'y'
	if mtok.String() != "Janet" {
		t.Fatalf("expected %v, got %v", "Jason", mtok.String())
	}
}

func get(json, path string) Result {
	return GetBytes([]byte(json), path)
}

func TestBasic(t *testing.T) {
	var mtok Result
	mtok = get(basicJSON, `loggy.programmers.#[tag="good"].firstName`)
	if mtok.String() != "Brett" {
		t.Fatalf("expected %v, got %v", "Brett", mtok.String())
	}
	mtok = get(basicJSON, `loggy.programmers.#[tag="good"]#.firstName`)
	if mtok.String() != `["Brett","Elliotte"]` {
		t.Fatalf("expected %v, got %v", `["Brett","Elliotte"]`, mtok.String())
	}
}
func TestPlus53BitInts(t *testing.T) {
	json := `{"IdentityData":{"GameInstanceId":634866135153775564}}`
	value := Get(json, "IdentityData.GameInstanceId")
	assert(t, value.Uint() == 634866135153775564)
	assert(t, value.Int() == 634866135153775564)
	assert(t, value.Float() == 634866135153775616)

	json = `{"IdentityData":{"GameInstanceId":634866135153775564.88172}}`
	value = Get(json, "IdentityData.GameInstanceId")
	assert(t, value.Uint() == 634866135153775616)
	assert(t, value.Int() == 634866135153775616)
	assert(t, value.Float() == 634866135153775616.88172)

	json = `{
		"min_uint64": 0,
		"max_uint64": 18446744073709551615,
		"overflow_uint64": 18446744073709551616,
		"min_int64": -9223372036854775808,
		"max_int64": 9223372036854775807,
		"overflow_int64": 9223372036854775808,
		"min_uint53":  0,
		"max_uint53":  4503599627370495,
		"overflow_uint53": 4503599627370496,
		"min_int53": -2251799813685248,
		"max_int53": 2251799813685247,
		"overflow_int53": 2251799813685248
	}`

	assert(t, Get(json, "min_uint53").Uint() == 0)
	assert(t, Get(json, "max_uint53").Uint() == 4503599627370495)
	assert(t, Get(json, "overflow_uint53").Int() == 4503599627370496)
	assert(t, Get(json, "min_int53").Int() == -2251799813685248)
	assert(t, Get(json, "max_int53").Int() == 2251799813685247)
	assert(t, Get(json, "overflow_int53").Int() == 2251799813685248)
	assert(t, Get(json, "min_uint64").Uint() == 0)
	assert(t, Get(json, "max_uint64").Uint() == 18446744073709551615)
	// this next value overflows the max uint64 by one which will just
	// flip the number to zero
	assert(t, Get(json, "overflow_uint64").Int() == 0)
	assert(t, Get(json, "min_int64").Int() == -9223372036854775808)
	assert(t, Get(json, "max_int64").Int() == 9223372036854775807)
	// this next value overflows the max int64 by one which will just
	// flip the number to the negative sign.
	assert(t, Get(json, "overflow_int64").Int() == -9223372036854775808)
}

func TestTypes(t *testing.T) {
	assert(t, (Result{Type: String}).Type.String() == "String")
	assert(t, (Result{Type: Number}).Type.String() == "Number")
	assert(t, (Result{Type: Null}).Type.String() == "Null")
	assert(t, (Result{Type: False}).Type.String() == "False")
	assert(t, (Result{Type: True}).Type.String() == "True")
	assert(t, (Result{Type: JSON}).Type.String() == "JSON")
	assert(t, (Result{Type: 100}).Type.String() == "")
	// bool
	assert(t, (Result{Type: String, Str: "true"}).Bool())
	assert(t, (Result{Type: True}).Bool())
	assert(t, (Result{Type: False}).Bool() == false)
	assert(t, (Result{Type: Number, Num: 1}).Bool())
	// int
	assert(t, (Result{Type: String, Str: "1"}).Int() == 1)
	assert(t, (Result{Type: True}).Int() == 1)
	assert(t, (Result{Type: False}).Int() == 0)
	assert(t, (Result{Type: Number, Num: 1}).Int() == 1)
	// uint
	assert(t, (Result{Type: String, Str: "1"}).Uint() == 1)
	assert(t, (Result{Type: True}).Uint() == 1)
	assert(t, (Result{Type: False}).Uint() == 0)
	assert(t, (Result{Type: Number, Num: 1}).Uint() == 1)
	// float
	assert(t, (Result{Type: String, Str: "1"}).Float() == 1)
	assert(t, (Result{Type: True}).Float() == 1)
	assert(t, (Result{Type: False}).Float() == 0)
	assert(t, (Result{Type: Number, Num: 1}).Float() == 1)
}
func TestForEach(t *testing.T) {
	Result{}.ForEach(nil)
	Result{Type: String, Str: "Hello"}.ForEach(func(_, value Result) bool {
		assert(t, value.String() == "Hello")
		return false
	})
	Result{Type: JSON, Raw: "*invalid*"}.ForEach(nil)

	json := ` {"name": {"first": "Janet","last": "Prichard"},
	"asd\nf":"\ud83d\udd13","age": 47}`
	var count int
	ParseBytes([]byte(json)).ForEach(func(key, value Result) bool {
		count++
		return true
	})
	assert(t, count == 3)
	ParseBytes([]byte(`{"bad`)).ForEach(nil)
	ParseBytes([]byte(`{"ok":"bad`)).ForEach(nil)
}
func TestMap(t *testing.T) {
	assert(t, len(ParseBytes([]byte(`"asdf"`)).Map()) == 0)
	assert(t, ParseBytes([]byte(`{"asdf":"ghjk"`)).Map()["asdf"].String() == "ghjk")
	assert(t, len(Result{Type: JSON, Raw: "**invalid**"}.Map()) == 0)
	assert(t, Result{Type: JSON, Raw: "**invalid**"}.Value() == nil)
	assert(t, Result{Type: JSON, Raw: "{"}.Map() != nil)
}
func TestBasic1(t *testing.T) {
	mtok := get(basicJSON, `loggy.programmers`)
	var count int
	mtok.ForEach(func(key, value Result) bool {
		if key.Exists() {
			t.Fatalf("expected %v, got %v", false, key.Exists())
		}
		count++
		if count == 3 {
			return false
		}
		if count == 1 {
			i := 0
			value.ForEach(func(key, value Result) bool {
				switch i {
				case 0:
					if key.String() != "firstName" || value.String() != "Brett" {
						t.Fatalf("expected %v/%v got %v/%v", "firstName", "Brett", key.String(), value.String())
					}
				case 1:
					if key.String() != "lastName" || value.String() != "McLaughlin" {
						t.Fatalf("expected %v/%v got %v/%v", "lastName", "McLaughlin", key.String(), value.String())
					}
				case 2:
					if key.String() != "email" || value.String() != "aaaa" {
						t.Fatalf("expected %v/%v got %v/%v", "email", "aaaa", key.String(), value.String())
					}
				}
				i++
				return true
			})
		}
		return true
	})
	if count != 3 {
		t.Fatalf("expected %v, got %v", 3, count)
	}
}
func TestBasic2(t *testing.T) {
	mtok := get(basicJSON, `loggy.programmers.#[age=101].firstName`)
	if mtok.String() != "1002.3" {
		t.Fatalf("expected %v, got %v", "1002.3", mtok.String())
	}
	mtok = get(basicJSON, `loggy.programmers.#[firstName != "Brett"].firstName`)
	if mtok.String() != "Jason" {
		t.Fatalf("expected %v, got %v", "Jason", mtok.String())
	}
	mtok = get(basicJSON, `loggy.programmers.#[firstName % "Bre*"].email`)
	if mtok.String() != "aaaa" {
		t.Fatalf("expected %v, got %v", "aaaa", mtok.String())
	}
	mtok = get(basicJSON, `loggy.programmers.#[firstName == "Brett"].email`)
	if mtok.String() != "aaaa" {
		t.Fatalf("expected %v, got %v", "aaaa", mtok.String())
	}
	mtok = get(basicJSON, "loggy")
	if mtok.Type != JSON {
		t.Fatalf("expected %v, got %v", JSON, mtok.Type)
	}
	if len(mtok.Map()) != 1 {
		t.Fatalf("expected %v, got %v", 1, len(mtok.Map()))
	}
	programmers := mtok.Map()["programmers"]
	if programmers.Array()[1].Map()["firstName"].Str != "Jason" {
		t.Fatalf("expected %v, got %v", "Jason", mtok.Map()["programmers"].Array()[1].Map()["firstName"].Str)
	}
}
func TestBasic3(t *testing.T) {
	var mtok Result
	if Parse(basicJSON).Get("loggy.programmers").Get("1").Get("firstName").Str != "Jason" {
		t.Fatalf("expected %v, got %v", "Jason", Parse(basicJSON).Get("loggy.programmers").Get("1").Get("firstName").Str)
	}
	var token Result
	if token = Parse("-102"); token.Num != -102 {
		t.Fatalf("expected %v, got %v", -102, token.Num)
	}
	if token = Parse("102"); token.Num != 102 {
		t.Fatalf("expected %v, got %v", 102, token.Num)
	}
	if token = Parse("102.2"); token.Num != 102.2 {
		t.Fatalf("expected %v, got %v", 102.2, token.Num)
	}
	if token = Parse(`"hello"`); token.Str != "hello" {
		t.Fatalf("expected %v, got %v", "hello", token.Str)
	}
	if token = Parse(`"\"he\nllo\""`); token.Str != "\"he\nllo\"" {
		t.Fatalf("expected %v, got %v", "\"he\nllo\"", token.Str)
	}
	mtok = get(basicJSON, "loggy.programmers.#.firstName")
	if len(mtok.Array()) != 4 {
		t.Fatalf("expected 4, got %v", len(mtok.Array()))
	}
	for i, ex := range []string{"Brett", "Jason", "Elliotte", "1002.3"} {
		if mtok.Array()[i].String() != ex {
			t.Fatalf("expected '%v', got '%v'", ex, mtok.Array()[i].String())
		}
	}
	mtok = get(basicJSON, "loggy.programmers.#.asd")
	if mtok.Type != JSON {
		t.Fatalf("expected %v, got %v", JSON, mtok.Type)
	}
	if len(mtok.Array()) != 0 {
		t.Fatalf("expected 0, got %v", len(mtok.Array()))
	}
}
func TestBasic4(t *testing.T) {
	if get(basicJSON, "items.3.tags.#").Num != 3 {
		t.Fatalf("expected 3, got %v", get(basicJSON, "items.3.tags.#").Num)
	}
	if get(basicJSON, "items.3.points.1.#").Num != 2 {
		t.Fatalf("expected 2, got %v", get(basicJSON, "items.3.points.1.#").Num)
	}
	if get(basicJSON, "items.#").Num != 8 {
		t.Fatalf("expected 6, got %v", get(basicJSON, "items.#").Num)
	}
	if get(basicJSON, "vals.#").Num != 4 {
		t.Fatalf("expected 4, got %v", get(basicJSON, "vals.#").Num)
	}
	if !get(basicJSON, "name.last").Exists() {
		t.Fatal("expected true, got false")
	}
	token := get(basicJSON, "name.here")
	if token.String() != "B\\\"R" {
		t.Fatal("expecting 'B\\\"R'", "got", token.String())
	}
	token = get(basicJSON, "arr.#")
	if token.String() != "6" {
		t.Fatal("expecting '6'", "got", token.String())
	}
	token = get(basicJSON, "arr.3.hello")
	if token.String() != "world" {
		t.Fatal("expecting 'world'", "got", token.String())
	}
	_ = token.Value().(string)
	token = get(basicJSON, "name.first")
	if token.String() != "tom" {
		t.Fatal("expecting 'tom'", "got", token.String())
	}
	_ = token.Value().(string)
	token = get(basicJSON, "name.last")
	if token.String() != "" {
		t.Fatal("expecting ''", "got", token.String())
	}
	if token.Value() != nil {
		t.Fatal("should be nil")
	}
}
func TestBasic5(t *testing.T) {
	token := get(basicJSON, "age")
	if token.String() != "100" {
		t.Fatal("expecting '100'", "got", token.String())
	}
	_ = token.Value().(float64)
	token = get(basicJSON, "happy")
	if token.String() != "true" {
		t.Fatal("expecting 'true'", "got", token.String())
	}
	_ = token.Value().(bool)
	token = get(basicJSON, "immortal")
	if token.String() != "false" {
		t.Fatal("expecting 'false'", "got", token.String())
	}
	_ = token.Value().(bool)
	token = get(basicJSON, "noop")
	if token.String() != `{"what is a wren?":"a bird"}` {
		t.Fatal("expecting '"+`{"what is a wren?":"a bird"}`+"'", "got", token.String())
	}
	_ = token.Value().(map[string]interface{})

	if get(basicJSON, "").Value() != nil {
		t.Fatal("should be nil")
	}

	get(basicJSON, "vals.hello")

	mm := Parse(basicJSON).Value().(map[string]interface{})
	fn := mm["loggy"].(map[string]interface{})["programmers"].([]interface{})[1].(map[string]interface{})["firstName"].(string)
	if fn != "Jason" {
		t.Fatalf("expecting %v, got %v", "Jason", fn)
	}
}
func TestUnicode(t *testing.T) {
	var json = `{"key":0,"的情况下解":{"key":1,"的情况":2}}`
	if Get(json, "的情况下解.key").Num != 1 {
		t.Fatal("fail")
	}
	if Get(json, "的情况下解.的情况").Num != 2 {
		t.Fatal("fail")
	}
	if Get(json, "的情况下解.的?况").Num != 2 {
		t.Fatal("fail")
	}
	if Get(json, "的情况下解.的?*").Num != 2 {
		t.Fatal("fail")
	}
	if Get(json, "的情况下解.*?况").Num != 2 {
		t.Fatal("fail")
	}
	if Get(json, "的情?下解.*?况").Num != 2 {
		t.Fatal("fail")
	}
	if Get(json, "的情下解.*?况").Num != 0 {
		t.Fatal("fail")
	}
}

func TestUnescape(t *testing.T) {
	unescape(string([]byte{'\\', '\\', 0}))
	unescape(string([]byte{'\\', '/', '\\', 'b', '\\', 'f'}))
}
func assert(t testing.TB, cond bool) {
	if !cond {
		panic("assert failed")
	}
}
func TestLess(t *testing.T) {
	assert(t, !Result{Type: Null}.Less(Result{Type: Null}, true))
	assert(t, Result{Type: Null}.Less(Result{Type: False}, true))
	assert(t, Result{Type: Null}.Less(Result{Type: True}, true))
	assert(t, Result{Type: Null}.Less(Result{Type: JSON}, true))
	assert(t, Result{Type: Null}.Less(Result{Type: Number}, true))
	assert(t, Result{Type: Null}.Less(Result{Type: String}, true))
	assert(t, !Result{Type: False}.Less(Result{Type: Null}, true))
	assert(t, Result{Type: False}.Less(Result{Type: True}, true))
	assert(t, Result{Type: String, Str: "abc"}.Less(Result{Type: String, Str: "bcd"}, true))
	assert(t, Result{Type: String, Str: "ABC"}.Less(Result{Type: String, Str: "abc"}, true))
	assert(t, !Result{Type: String, Str: "ABC"}.Less(Result{Type: String, Str: "abc"}, false))
	assert(t, Result{Type: Number, Num: 123}.Less(Result{Type: Number, Num: 456}, true))
	assert(t, !Result{Type: Number, Num: 456}.Less(Result{Type: Number, Num: 123}, true))
	assert(t, !Result{Type: Number, Num: 456}.Less(Result{Type: Number, Num: 456}, true))
	assert(t, stringLessInsensitive("abcde", "BBCDE"))
	assert(t, stringLessInsensitive("abcde", "bBCDE"))
	assert(t, stringLessInsensitive("Abcde", "BBCDE"))
	assert(t, stringLessInsensitive("Abcde", "bBCDE"))
	assert(t, !stringLessInsensitive("bbcde", "aBCDE"))
	assert(t, !stringLessInsensitive("bbcde", "ABCDE"))
	assert(t, !stringLessInsensitive("Bbcde", "aBCDE"))
	assert(t, !stringLessInsensitive("Bbcde", "ABCDE"))
	assert(t, !stringLessInsensitive("abcde", "ABCDE"))
	assert(t, !stringLessInsensitive("Abcde", "ABCDE"))
	assert(t, !stringLessInsensitive("abcde", "ABCDE"))
	assert(t, !stringLessInsensitive("ABCDE", "ABCDE"))
	assert(t, !stringLessInsensitive("abcde", "abcde"))
	assert(t, !stringLessInsensitive("123abcde", "123Abcde"))
	assert(t, !stringLessInsensitive("123Abcde", "123Abcde"))
	assert(t, !stringLessInsensitive("123Abcde", "123abcde"))
	assert(t, !stringLessInsensitive("123abcde", "123abcde"))
	assert(t, !stringLessInsensitive("124abcde", "123abcde"))
	assert(t, !stringLessInsensitive("124Abcde", "123Abcde"))
	assert(t, !stringLessInsensitive("124Abcde", "123abcde"))
	assert(t, !stringLessInsensitive("124abcde", "123abcde"))
	assert(t, stringLessInsensitive("124abcde", "125abcde"))
	assert(t, stringLessInsensitive("124Abcde", "125Abcde"))
	assert(t, stringLessInsensitive("124Abcde", "125abcde"))
	assert(t, stringLessInsensitive("124abcde", "125abcde"))
}

func TestIssue6(t *testing.T) {
	data := `{
      "code": 0,
      "msg": "",
      "data": {
        "sz002024": {
          "qfqday": [
            [
              "2014-01-02",
              "8.93",
              "9.03",
              "9.17",
              "8.88",
              "621143.00"
            ],
            [
              "2014-01-03",
              "9.03",
              "9.30",
              "9.47",
              "8.98",
              "1624438.00"
            ]
          ]
        }
      }
    }`

	var num []string
	for _, v := range Get(data, "data.sz002024.qfqday.0").Array() {
		num = append(num, v.String())
	}
	if fmt.Sprintf("%v", num) != "[2014-01-02 8.93 9.03 9.17 8.88 621143.00]" {
		t.Fatalf("invalid result")
	}
}

var exampleJSON = `{
	"widget": {
		"debug": "on",
		"window": {
			"title": "Sample Konfabulator Widget",
			"name": "main_window",
			"width": 500,
			"height": 500
		},
		"image": {
			"src": "Images/Sun.png",
			"hOffset": 250,
			"vOffset": 250,
			"alignment": "center"
		},
		"text": {
			"data": "Click Here",
			"size": 36,
			"style": "bold",
			"vOffset": 100,
			"alignment": "center",
			"onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"
		}
	}
}`

func TestNewParse(t *testing.T) {
	//fmt.Printf("%v\n", parse2(exampleJSON, "widget").String())
}

func TestUnmarshalMap(t *testing.T) {
	var m1 = Parse(exampleJSON).Value().(map[string]interface{})
	var m2 map[string]interface{}
	if err := json.Unmarshal([]byte(exampleJSON), &m2); err != nil {
		t.Fatal(err)
	}
	b1, err := json.Marshal(m1)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := json.Marshal(m2)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Compare(b1, b2) != 0 {
		t.Fatal("b1 != b2")
	}
}

func TestSingleArrayValue(t *testing.T) {
	var json = `{"key": "value","key2":[1,2,3,4,"A"]}`
	var result = Get(json, "key")
	var array = result.Array()
	if len(array) != 1 {
		t.Fatal("array is empty")
	}
	if array[0].String() != "value" {
		t.Fatalf("got %s, should be %s", array[0].String(), "value")
	}

	array = Get(json, "key2.#").Array()
	if len(array) != 1 {
		t.Fatalf("got '%v', expected '%v'", len(array), 1)
	}

	array = Get(json, "key3").Array()
	if len(array) != 0 {
		t.Fatalf("got '%v', expected '%v'", len(array), 0)
	}

}

var manyJSON = `  {
	"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{
	"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{
	"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{
	"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{
	"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{
	"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{
	"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"a":{"hello":"world"
	}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}
	"position":{"type":"Point","coordinates":[-115.24,33.09]},
	"loves":["world peace"],
	"name":{"last":"Anderson","first":"Nancy"},
	"age":31
	"":{"a":"emptya","b":"emptyb"},
	"name.last":"Yellow",
	"name.first":"Cat",
}`

func combine(results []Result) string {
	return fmt.Sprintf("%v", results)
}
func TestManyBasic(t *testing.T) {
	testWatchForFallback = true
	defer func() {
		testWatchForFallback = false
	}()
	testMany := func(shouldFallback bool, expect string, paths ...string) {
		results := GetManyBytes(
			[]byte(manyJSON),
			paths...,
		)
		if len(results) != len(paths) {
			t.Fatalf("expected %v, got %v", len(paths), len(results))
		}
		if fmt.Sprintf("%v", results) != expect {
			fmt.Printf("%v\n", paths)
			t.Fatalf("expected %v, got %v", expect, results)
		}
		//if testLastWasFallback != shouldFallback {
		//	t.Fatalf("expected %v, got %v", shouldFallback, testLastWasFallback)
		//}
	}
	testMany(false, "[Point]", "position.type")
	testMany(false, `[emptya ["world peace"] 31]`, ".a", "loves", "age")
	testMany(false, `[["world peace"]]`, "loves")
	testMany(false, `[{"last":"Anderson","first":"Nancy"} Nancy]`, "name", "name.first")
	testMany(true, `[]`, strings.Repeat("a.", 40)+"hello")
	res := Get(manyJSON, strings.Repeat("a.", 48)+"a")
	testMany(true, `[`+res.String()+`]`, strings.Repeat("a.", 48)+"a")
	// these should fallback
	testMany(true, `[Cat Nancy]`, "name\\.first", "name.first")
	testMany(true, `[world]`, strings.Repeat("a.", 70)+"hello")
}
func testMany(t *testing.T, json string, paths, expected []string) {
	testManyAny(t, json, paths, expected, true)
	testManyAny(t, json, paths, expected, false)
}
func testManyAny(t *testing.T, json string, paths, expected []string, bytes bool) {
	var result []Result
	for i := 0; i < 2; i++ {
		var which string
		if i == 0 {
			which = "Get"
			result = nil
			for j := 0; j < len(expected); j++ {
				if bytes {
					result = append(result, GetBytes([]byte(json), paths[j]))
				} else {
					result = append(result, Get(json, paths[j]))
				}
			}
		} else if i == 1 {
			which = "GetMany"
			if bytes {
				result = GetManyBytes([]byte(json), paths...)
			} else {
				result = GetMany(json, paths...)
			}
		}
		for j := 0; j < len(expected); j++ {
			if result[j].String() != expected[j] {
				t.Fatalf("Using key '%s' for '%s'\nexpected '%v', got '%v'", paths[j], which, expected[j], result[j].String())
			}
		}
	}
}
func TestIssue20(t *testing.T) {
	json := `{ "name": "FirstName", "name1": "FirstName1", "address": "address1", "addressDetails": "address2", }`
	paths := []string{"name", "name1", "address", "addressDetails"}
	expected := []string{"FirstName", "FirstName1", "address1", "address2"}
	t.Run("SingleMany", func(t *testing.T) { testMany(t, json, paths, expected) })
}

func TestIssue21(t *testing.T) {
	json := `{ "Level1Field1":3, 
	           "Level1Field4":4, 
			   "Level1Field2":{ "Level2Field1":[ "value1", "value2" ], 
			   "Level2Field2":{ "Level3Field1":[ { "key1":"value1" } ] } } }`
	paths := []string{"Level1Field1", "Level1Field2.Level2Field1", "Level1Field2.Level2Field2.Level3Field1", "Level1Field4"}
	expected := []string{"3", `[ "value1", "value2" ]`, `[ { "key1":"value1" } ]`, "4"}
	t.Run("SingleMany", func(t *testing.T) { testMany(t, json, paths, expected) })
}

func TestRandomMany(t *testing.T) {
	var lstr string
	defer func() {
		if v := recover(); v != nil {
			println("'" + hex.EncodeToString([]byte(lstr)) + "'")
			println("'" + lstr + "'")
			panic(v)
		}
	}()
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 512)
	for i := 0; i < 50000; i++ {
		n, err := rand.Read(b[:rand.Int()%len(b)])
		if err != nil {
			t.Fatal(err)
		}
		lstr = string(b[:n])
		paths := make([]string, rand.Int()%64)
		for i := range paths {
			var b []byte
			n := rand.Int() % 5
			for j := 0; j < n; j++ {
				if j > 0 {
					b = append(b, '.')
				}
				nn := rand.Int() % 10
				for k := 0; k < nn; k++ {
					b = append(b, 'a'+byte(rand.Int()%26))
				}
			}
			paths[i] = string(b)
		}
		GetMany(lstr, paths...)
	}
}

type ComplicatedType struct {
	unsettable int
	Tagged     string `json:"tagged"`
	NotTagged  bool
	Nested     struct {
		Yellow string `json:"yellow"`
	}
	NestedTagged struct {
		Green string
		Map   map[string]interface{}
		Ints  struct {
			Int   int `json:"int"`
			Int8  int8
			Int16 int16
			Int32 int32
			Int64 int64 `json:"int64"`
		}
		Uints struct {
			Uint   uint
			Uint8  uint8
			Uint16 uint16
			Uint32 uint32
			Uint64 uint64
		}
		Floats struct {
			Float64 float64
			Float32 float32
		}
		Byte byte
		Bool bool
	} `json:"nestedTagged"`
	LeftOut      string `json:"-"`
	SelfPtr      *ComplicatedType
	SelfSlice    []ComplicatedType
	SelfSlicePtr []*ComplicatedType
	SelfPtrSlice *[]ComplicatedType
	Interface    interface{} `json:"interface"`
	Array        [3]int
	Time         time.Time `json:"time"`
	Binary       []byte
	NonBinary    []byte
}

var complicatedJSON = `
{
	"tagged": "OK",
	"Tagged": "KO",
	"NotTagged": true,
	"unsettable": 101,
	"Nested": {
		"Yellow": "Green",
		"yellow": "yellow"
	},
	"nestedTagged": {
		"Green": "Green",
		"Map": {
			"this": "that", 
			"and": "the other thing"
		},
		"Ints": {
			"Uint": 99,
			"Uint16": 16,
			"Uint32": 32,
			"Uint64": 65
		},
		"Uints": {
			"int": -99,
			"Int": -98,
			"Int16": -16,
			"Int32": -32,
			"int64": -64,
			"Int64": -65
		},
		"Uints": {
			"Float32": 32.32,
			"Float64": 64.64
		},
		"Byte": 254,
		"Bool": true
	},
	"LeftOut": "you shouldn't be here",
	"SelfPtr": {"tagged":"OK","nestedTagged":{"Ints":{"Uint32":32}}},
	"SelfSlice": [{"tagged":"OK","nestedTagged":{"Ints":{"Uint32":32}}}],
	"SelfSlicePtr": [{"tagged":"OK","nestedTagged":{"Ints":{"Uint32":32}}}],
	"SelfPtrSlice": [{"tagged":"OK","nestedTagged":{"Ints":{"Uint32":32}}}],
	"interface": "Tile38 Rocks!",
	"Interface": "Please Download",
	"Array": [0,2,3,4,5],
	"time": "2017-05-07T13:24:43-07:00",
	"Binary": "R0lGODlhPQBEAPeo",
	"NonBinary": [9,3,100,115]
}
`

func TestUnmarshal(t *testing.T) {
	var s1 ComplicatedType
	var s2 ComplicatedType
	if err := json.Unmarshal([]byte(complicatedJSON), &s1); err != nil {
		t.Fatal(err)
	}
	if err := Unmarshal([]byte(complicatedJSON), &s2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(&s1, &s2) {
		t.Fatal("not equal")
	}
	var str string
	if err := json.Unmarshal([]byte(Get(complicatedJSON, "LeftOut").Raw), &str); err != nil {
		t.Fatal(err)
	}
	assert(t, str == Get(complicatedJSON, "LeftOut").String())
}

func testvalid(json string, expect bool) {
	_, ok := validpayload([]byte(json), 0)
	if ok != expect {
		panic("mismatch")
	}
}

func TestValidBasic(t *testing.T) {
	testvalid("0", true)
	testvalid("00", false)
	testvalid("-00", false)
	testvalid("-.", false)
	testvalid("0.0", true)
	testvalid("10.0", true)
	testvalid("10e1", true)
	testvalid("10EE", false)
	testvalid("10E-", false)
	testvalid("10E+", false)
	testvalid("10E123", true)
	testvalid("10E-123", true)
	testvalid("10E-0123", true)
	testvalid("", false)
	testvalid(" ", false)
	testvalid("{}", true)
	testvalid("{", false)
	testvalid("-", false)
	testvalid("-1", true)
	testvalid("-1.", false)
	testvalid("-1.0", true)
	testvalid(" -1.0", true)
	testvalid(" -1.0 ", true)
	testvalid("-1.0 ", true)
	testvalid("-1.0 i", false)
	testvalid("-1.0 i", false)
	testvalid("true", true)
	testvalid(" true", true)
	testvalid(" true ", true)
	testvalid(" True ", false)
	testvalid(" tru", false)
	testvalid("false", true)
	testvalid(" false", true)
	testvalid(" false ", true)
	testvalid(" False ", false)
	testvalid(" fals", false)
	testvalid("null", true)
	testvalid(" null", true)
	testvalid(" null ", true)
	testvalid(" Null ", false)
	testvalid(" nul", false)
	testvalid(" []", true)
	testvalid(" [true]", true)
	testvalid(" [ true, null ]", true)
	testvalid(" [ true,]", false)
	testvalid(`{"hello":"world"}`, true)
	testvalid(`{ "hello": "world" }`, true)
	testvalid(`{ "hello": "world", }`, false)
	testvalid(`{"a":"b",}`, false)
	testvalid(`{"a":"b","a"}`, false)
	testvalid(`{"a":"b","a":}`, false)
	testvalid(`{"a":"b","a":1}`, true)
	testvalid(`{"a":"b","a": 1, "c":{"hi":"there"} }`, true)
	testvalid(`{"a":"b","a": 1, "c":{"hi":"there", "easy":["going",{"mixed":"bag"}]} }`, true)
	testvalid(`""`, true)
	testvalid(`"`, false)
	testvalid(`"\n"`, true)
	testvalid(`"\"`, false)
	testvalid(`"\\"`, true)
	testvalid(`"a\\b"`, true)
	testvalid(`"a\\b\\\"a"`, true)
	testvalid(`"a\\b\\\uFFAAa"`, true)
	testvalid(`"a\\b\\\uFFAZa"`, false)
	testvalid(`"a\\b\\\uFFA"`, false)
	testvalid(string(complicatedJSON), true)
	testvalid(string(exampleJSON), true)
}

var jsonchars = []string{"{", "[", ",", ":", "}", "]", "1", "0", "true", "false", "null", `""`, `"\""`, `"a"`}

func makeRandomJSONChars(b []byte) {
	var bb []byte
	for len(bb) < len(b) {
		bb = append(bb, jsonchars[rand.Int()%len(jsonchars)]...)
	}
	copy(b, bb[:len(b)])
}
func TestValidRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 100000)
	start := time.Now()
	for time.Since(start) < time.Second*3 {
		n := rand.Int() % len(b)
		rand.Read(b[:n])
		validpayload(b[:n], 0)
	}

	start = time.Now()
	for time.Since(start) < time.Second*3 {
		n := rand.Int() % len(b)
		makeRandomJSONChars(b[:n])
		validpayload(b[:n], 0)
	}
}

type BenchStruct struct {
	Widget struct {
		Window struct {
			Name string `json:"name"`
		} `json:"window"`
		Image struct {
			HOffset int `json:"hOffset"`
		} `json:"image"`
		Text struct {
			OnMouseUp string `json:"onMouseUp"`
		} `json:"text"`
	} `json:"widget"`
}

var benchPaths = []string{
	"widget.window.name",
	"widget.image.hOffset",
	"widget.text.onMouseUp",
}

var benchManyPaths = []string{
	"widget.window.name",
	"widget.image.hOffset",
	"widget.text.onMouseUp",
	"widget.window.title",
	"widget.image.alignment",
	"widget.text.style",
	"widget.window.height",
	"widget.image.src",
	"widget.text.data",
	"widget.text.size",
}

func BenchmarkGJSONGet(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j := 0; j < len(benchPaths); j++ {
			if Get(exampleJSON, benchPaths[j]).Type == Null {
				t.Fatal("did not find the value")
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}
func BenchmarkGJSONGetMany4Paths(t *testing.B) {
	benchmarkGJSONGetManyN(t, 4)
}
func BenchmarkGJSONGetMany8Paths(t *testing.B) {
	benchmarkGJSONGetManyN(t, 8)
}
func BenchmarkGJSONGetMany16Paths(t *testing.B) {
	benchmarkGJSONGetManyN(t, 16)
}
func BenchmarkGJSONGetMany32Paths(t *testing.B) {
	benchmarkGJSONGetManyN(t, 32)
}
func BenchmarkGJSONGetMany64Paths(t *testing.B) {
	benchmarkGJSONGetManyN(t, 64)
}
func BenchmarkGJSONGetMany128Paths(t *testing.B) {
	benchmarkGJSONGetManyN(t, 128)
}
func benchmarkGJSONGetManyN(t *testing.B, n int) {
	var paths []string
	for len(paths) < n {
		paths = append(paths, benchManyPaths...)
	}
	paths = paths[:n]
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		results := GetMany(exampleJSON, paths...)
		if len(results) == 0 {
			t.Fatal("did not find the value")
		}
		for j := 0; j < len(results); j++ {
			if results[j].Type == Null {
				t.Fatal("did not find the value")
			}
		}
	}
	t.N *= len(paths) // because we are running against 3 paths
}

func BenchmarkGJSONUnmarshalMap(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j := 0; j < len(benchPaths); j++ {
			parts := strings.Split(benchPaths[j], ".")
			m, _ := Parse(exampleJSON).Value().(map[string]interface{})
			var v interface{}
			for len(parts) > 0 {
				part := parts[0]
				if len(parts) > 1 {
					m = m[part].(map[string]interface{})
					if m == nil {
						t.Fatal("did not find the value")
					}
				} else {
					v = m[part]
					if v == nil {
						t.Fatal("did not find the value")
					}
				}
				parts = parts[1:]
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}

func BenchmarkGJSONUnmarshalStruct(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j := 0; j < len(benchPaths); j++ {
			var s BenchStruct
			if err := Unmarshal([]byte(exampleJSON), &s); err != nil {
				t.Fatal(err)
			}
			switch benchPaths[j] {
			case "widget.window.name":
				if s.Widget.Window.Name == "" {
					t.Fatal("did not find the value")
				}
			case "widget.image.hOffset":
				if s.Widget.Image.HOffset == 0 {
					t.Fatal("did not find the value")
				}
			case "widget.text.onMouseUp":
				if s.Widget.Text.OnMouseUp == "" {
					t.Fatal("did not find the value")
				}
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}

func BenchmarkJSONUnmarshalMap(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j := 0; j < len(benchPaths); j++ {
			parts := strings.Split(benchPaths[j], ".")
			var m map[string]interface{}
			if err := json.Unmarshal([]byte(exampleJSON), &m); err != nil {
				t.Fatal(err)
			}
			var v interface{}
			for len(parts) > 0 {
				part := parts[0]
				if len(parts) > 1 {
					m = m[part].(map[string]interface{})
					if m == nil {
						t.Fatal("did not find the value")
					}
				} else {
					v = m[part]
					if v == nil {
						t.Fatal("did not find the value")
					}
				}
				parts = parts[1:]
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}

func BenchmarkJSONUnmarshalStruct(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j := 0; j < len(benchPaths); j++ {
			var s BenchStruct
			if err := json.Unmarshal([]byte(exampleJSON), &s); err != nil {
				t.Fatal(err)
			}
			switch benchPaths[j] {
			case "widget.window.name":
				if s.Widget.Window.Name == "" {
					t.Fatal("did not find the value")
				}
			case "widget.image.hOffset":
				if s.Widget.Image.HOffset == 0 {
					t.Fatal("did not find the value")
				}
			case "widget.text.onMouseUp":
				if s.Widget.Text.OnMouseUp == "" {
					t.Fatal("did not find the value")
				}
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}

func BenchmarkJSONDecoder(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j := 0; j < len(benchPaths); j++ {
			dec := json.NewDecoder(bytes.NewBuffer([]byte(exampleJSON)))
			var found bool
		outer:
			for {
				tok, err := dec.Token()
				if err != nil {
					if err == io.EOF {
						break
					}
					t.Fatal(err)
				}
				switch v := tok.(type) {
				case string:
					if found {
						// break out once we find the value.
						break outer
					}
					switch benchPaths[j] {
					case "widget.window.name":
						if v == "name" {
							found = true
						}
					case "widget.image.hOffset":
						if v == "hOffset" {
							found = true
						}
					case "widget.text.onMouseUp":
						if v == "onMouseUp" {
							found = true
						}
					}
				}
			}
			if !found {
				t.Fatal("field not found")
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}

func BenchmarkFFJSONLexer(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j := 0; j < len(benchPaths); j++ {
			l := fflib.NewFFLexer([]byte(exampleJSON))
			var found bool
		outer:
			for {
				t := l.Scan()
				if t == fflib.FFTok_eof {
					break
				}
				if t == fflib.FFTok_string {
					b, _ := l.CaptureField(t)
					v := string(b)
					if found {
						// break out once we find the value.
						break outer
					}
					switch benchPaths[j] {
					case "widget.window.name":
						if v == "\"name\"" {
							found = true
						}
					case "widget.image.hOffset":
						if v == "\"hOffset\"" {
							found = true
						}
					case "widget.text.onMouseUp":
						if v == "\"onMouseUp\"" {
							found = true
						}
					}
				}
			}
			if !found {
				t.Fatal("field not found")
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}

func skipCC(l *jlexer.Lexer, n int) {
	for i := 0; i < n; i++ {
		l.Skip()
		l.WantColon()
		l.Skip()
		l.WantComma()
	}
}
func skipGroup(l *jlexer.Lexer, n int) {
	l.WantColon()
	l.Delim('{')
	skipCC(l, n)
	l.Delim('}')
	l.WantComma()
}
func easyJSONWindowName(t *testing.B, l *jlexer.Lexer) {
	if l.String() == "window" {
		l.WantColon()
		l.Delim('{')
		skipCC(l, 1)
		if l.String() == "name" {
			l.WantColon()
			if l.String() == "" {
				t.Fatal("did not find the value")
			}
		}
	}
}
func easyJSONImageHOffset(t *testing.B, l *jlexer.Lexer) {
	if l.String() == "image" {
		l.WantColon()
		l.Delim('{')
		skipCC(l, 1)
		if l.String() == "hOffset" {
			l.WantColon()
			if l.Int() == 0 {
				t.Fatal("did not find the value")
			}
		}
	}
}
func easyJSONTextOnMouseUp(t *testing.B, l *jlexer.Lexer) {
	if l.String() == "text" {
		l.WantColon()
		l.Delim('{')
		skipCC(l, 5)
		if l.String() == "onMouseUp" {
			l.WantColon()
			if l.String() == "" {
				t.Fatal("did not find the value")
			}
		}
	}
}
func easyJSONWidget(t *testing.B, l *jlexer.Lexer, j int) {
	l.WantColon()
	l.Delim('{')
	switch benchPaths[j] {
	case "widget.window.name":
		skipCC(l, 1)
		easyJSONWindowName(t, l)
	case "widget.image.hOffset":
		skipCC(l, 1)
		if l.String() == "window" {
			skipGroup(l, 4)
		}
		easyJSONImageHOffset(t, l)
	case "widget.text.onMouseUp":
		skipCC(l, 1)
		if l.String() == "window" {
			skipGroup(l, 4)
		}
		if l.String() == "image" {
			skipGroup(l, 4)
		}
		easyJSONTextOnMouseUp(t, l)
	}
}
func BenchmarkEasyJSONLexer(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j := 0; j < len(benchPaths); j++ {
			l := &jlexer.Lexer{Data: []byte(exampleJSON)}
			l.Delim('{')
			if l.String() == "widget" {
				easyJSONWidget(t, l, j)
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}

func BenchmarkJSONParserGet(t *testing.B) {
	data := []byte(exampleJSON)
	keys := make([][]string, 0, len(benchPaths))
	for i := 0; i < len(benchPaths); i++ {
		keys = append(keys, strings.Split(benchPaths[i], "."))
	}
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j, k := range keys {
			if j == 1 {
				// "widget.image.hOffset" is a number
				v, _ := jsonparser.GetInt(data, k...)
				if v == 0 {
					t.Fatal("did not find the value")
				}
			} else {
				// "widget.window.name",
				// "widget.text.onMouseUp",
				v, _ := jsonparser.GetString(data, k...)
				if v == "" {
					t.Fatal("did not find the value")
				}
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}
func jsoniterWindowName(t *testing.B, iter *jsoniter.Iterator) {
	var v string
	for {
		key := iter.ReadObject()
		if key != "window" {
			iter.Skip()
			continue
		}
		for {
			key := iter.ReadObject()
			if key != "name" {
				iter.Skip()
				continue
			}
			v = iter.ReadString()
			break
		}
		break
	}
	if v == "" {
		t.Fatal("did not find the value")
	}
}

func jsoniterTextOnMouseUp(t *testing.B, iter *jsoniter.Iterator) {
	var v string
	for {
		key := iter.ReadObject()
		if key != "text" {
			iter.Skip()
			continue
		}
		for {
			key := iter.ReadObject()
			if key != "onMouseUp" {
				iter.Skip()
				continue
			}
			v = iter.ReadString()
			break
		}
		break
	}
	if v == "" {
		t.Fatal("did not find the value")
	}
}
func jsoniterImageOffset(t *testing.B, iter *jsoniter.Iterator) {
	var v int
	for {
		key := iter.ReadObject()
		if key != "image" {
			iter.Skip()
			continue
		}
		for {
			key := iter.ReadObject()
			if key != "hOffset" {
				iter.Skip()
				continue
			}
			v = iter.ReadInt()
			break
		}
		break
	}
	if v == 0 {
		t.Fatal("did not find the value")
	}
}
func jsoniterWidget(t *testing.B, iter *jsoniter.Iterator, j int) {
	for {
		key := iter.ReadObject()
		if key != "widget" {
			iter.Skip()
			continue
		}
		switch benchPaths[j] {
		case "widget.window.name":
			jsoniterWindowName(t, iter)
		case "widget.image.hOffset":
			jsoniterImageOffset(t, iter)
		case "widget.text.onMouseUp":
			jsoniterTextOnMouseUp(t, iter)
		}
		break
	}
}

// Commented by Aleksandr Rassanov, 'cause bad apicall of jsoniter.ParseString
//func BenchmarkJSONIterator(t *testing.B) {
//	t.ReportAllocs()
//	t.ResetTimer()
//	for i := 0; i < t.N; i++ {
//		for j := 0; j < len(benchPaths); j++ {
//			iter := jsoniter.ParseString(exampleJSON)
//			jsoniterWidget(t, iter, j)
//		}
//	}
//	t.N *= len(benchPaths) // because we are running against 3 paths
//}

var massiveJSON = func() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i := 0; i < 100; i++ {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(exampleJSON)
	}
	buf.WriteString("]")
	return buf.String()
}()

func BenchmarkConvertNone(t *testing.B) {
	json := massiveJSON
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		Get(json, "50.widget.text.onMouseUp")
	}
}
func BenchmarkConvertGet(t *testing.B) {
	data := []byte(massiveJSON)
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		Get(string(data), "50.widget.text.onMouseUp")
	}
}
func BenchmarkConvertGetBytes(t *testing.B) {
	data := []byte(massiveJSON)
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		GetBytes(data, "50.widget.text.onMouseUp")
	}
}
func BenchmarkParseUintNumParser(t *testing.B) {
	var s = "634866135153775564"
	for i := 0; i < t.N; i++ {
		parseUint(s)
	}
}
func BenchmarkStdlibParseUintNumParser(t *testing.B) {
	var s = "634866135153775564"
	for i := 0; i < t.N; i++ {
		strconv.ParseUint(s, 10, 64)
	}
}
func BenchmarkParseIntNumParser(t *testing.B) {
	var s = "-634866135153775564"
	for i := 0; i < t.N; i++ {
		parseInt(s)
	}
}
func BenchmarkStdlibParseIntNumParser(t *testing.B) {
	var s = "-634866135153775564"
	for i := 0; i < t.N; i++ {
		strconv.ParseInt(s, 10, 64)
	}
}
