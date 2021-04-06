package _1gob_demo

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

type s struct {
	data map[string]interface{}
}


func jsonDemo() {
	var s1 = s{
		data: make(map[string]interface{}, 8),
	}
	s1.data["count"] = 1 // int
	ret, err := json.Marshal(s1.data)
	if err != nil {
		fmt.Println("marshal failed", err)
	}
	fmt.Printf("%#v\n", string(ret))

	var s2 = s{
		data: make(map[string]interface{}, 8),
	}
	err = json.Unmarshal(ret, &s2.data)
	if err != nil {
		fmt.Println("unmarshal failed", err)
	}
	fmt.Println(s2)
	for _, v := range s2.data {
		fmt.Printf("value:%v, type:%T\n", v, v)
	}
}

func gobDemo() {
	var s1 = s{
		data: make(map[string]interface{}, 8),
	}
	s1.data["count"] = 1 // int
	// encode
	buf := new(bytes.Buffer) // 指针
	enc := gob.NewEncoder(buf) // 造一个编码器对象
	err := enc.Encode(s1.data)
	if err != nil {
		fmt.Println("gob encode failed, err:", err)
		return
	}
	b := buf.Bytes() // 拿到编码之后的字节数据
	fmt.Println(b)


	var s2 = s{
		data: make(map[string]interface{}, 8),
	}
	// decode
	dec := gob.NewDecoder(bytes.NewBuffer(b)) // 造一个新的解码器对象
	err = dec.Decode(&s2.data) // 把二进制数据解码到s2.data
	if err != nil {
		fmt.Println("gob decode failed, err", err)
		return
	}
	fmt.Println(s2.data)
	for _, v := range s2.data {
		fmt.Printf("value:%v, type:%T\n", v, v)
	}
}

func main(){
	//jsonDemo()
	gobDemo()
}

