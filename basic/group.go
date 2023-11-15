package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicGroup() {
	batman := basic.User{
		Id:       99,
		Username: "batman",
		IsActive: true,
		Password: []byte("baticave"),
		Emails:   []string{"batman@gmail.com", "bruceway@gmail.com"},
		Gender:   1,
	}

	nightwing := basic.User{
		Id:       100,
		Username: "nightwing",
		IsActive: true,
		Password: []byte("baticave"),
		Emails:   []string{"nightwing@gmail.com", "nightwingy@gmail.com"},
		Gender:   1,
	}

	robin := basic.User{
		Id:       101,
		Username: "Robin",
		IsActive: true,
		Password: []byte("baticave"),
		Emails:   []string{"robin@gmail.com", "robin@gmail.com"},
		Gender:   1,
	}

	batFamily := basic.Group{
		Id:          1,
		Name:        "Bat family",
		Users:       []*basic.User{&batman, &nightwing, &robin},
		Description: "The clasic bat family",
	}

	jsonBytes, _ := protojson.Marshal(&batFamily)

	log.Println(string(jsonBytes))

}
