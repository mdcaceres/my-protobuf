package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUser() {
	a := basic.Address{
		Street:  "gothan",
		City:    "Gothan",
		Country: "USA",
	}

	u := basic.User{
		Id:       99,
		Username: "batman",
		IsActive: true,
		Password: []byte("baticave"),
		Emails:   []string{"batman@gmail.com", "bruceway@gmail.com"},
		Gender:   1,
		Address:  &a,
	}

	jsonBytes, _ := protojson.Marshal(&u)

	log.Println(string(jsonBytes))
}

func ProtoUserToJson() {
	u := basic.User{
		Id:       99,
		Username: "batman",
		IsActive: true,
		Password: []byte("baticave"),
		Emails:   []string{"batman@gmail.com", "bruceway@gmail.com"},
		Gender:   1,
	}

	jsonBytes, _ := protojson.Marshal(&u)

	log.Println(string(jsonBytes))
}
