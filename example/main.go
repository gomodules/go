package main

import (
	"encoding/json"
	"fmt"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

type Instance struct {
	Label string `json:"label,omitempty" tf:"Label"`
	Alert *Alert `json:"alert,omitempty" tf:"Alert"`
}

type Alert struct {
	Name string `json:"name,omitempty" tf:"Name"`
}

type instanceCodec struct {
}

func (instanceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*Alert)(ptr) == nil
}

func (ic instanceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	alert := (*Alert)(ptr)
	var alerts []Alert
	if alert != nil {
		alerts = []Alert{*alert}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
	}.Froze()

	byt, _ := jsonit.Marshal(alerts)

	stream.Write(byt)
}

func (ic instanceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*Alert)(ptr) = Alert{}
		return
	case jsoniter.ArrayValue:
		alertsByte := iter.SkipAndReturnBytes()
		if len(alertsByte) > 0 {
			var alerts []Alert

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
			}.Froze()
			jsonit.Unmarshal(alertsByte, &alerts)

			if len(alerts) > 0 {
				*(*Alert)(ptr) = alerts[0]
			} else {
				*(*Alert)(ptr) = Alert{}
			}
		} else {
			*(*Alert)(ptr) = Alert{}
		}
	default:
		iter.ReportError("decode Alert", "unexpected JSON type")
	}
}

func main() {
	kind, err := jsoniter.GetKind(reflect2.TypeOf(Alert{}).Type1())
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeDecoders: map[string]jsoniter.ValDecoder{
			kind: instanceCodec{},
		},
		TypeEncoders: map[string]jsoniter.ValEncoder{
			kind: instanceCodec{},
		},
	}.Froze()

	instance := Instance{
		Label: "test",
		Alert: &Alert{
			Name: "test2",
		},
	}

	fmt.Print("JSON marshal output:\t\t\t")
	jsonByt, err := json.Marshal(instance)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonByt))

	fmt.Print("JSON Iterator marshal output:\t")
	jsonItByt, err := jsonit.Marshal(instance)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonItByt))

	fmt.Print("JSON unmarshal output:\t\t\t")
	jsonInstance := Instance{}
	err = json.Unmarshal(jsonByt, &jsonInstance)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", jsonInstance)

	fmt.Print("JSON Iterator unmarshal output:\t")
	jsonItInstance := Instance{}
	err = jsonit.Unmarshal(jsonItByt, &jsonItInstance)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", jsonItInstance)
}
