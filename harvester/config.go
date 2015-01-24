package main

import (
	"encoding/json"
	"os"
	"time"
)

type MgoDialInfo struct {
	// Addrs holds the addresses for the seed servers.
	Addrs []string `json:"addrs"`

	// Timeout is the amount of time to wait for a server to respond when
	// first connecting and on follow up operations in the session. If
	// timeout is zero, the call may block forever waiting for a connection
	// to be established.
	Timeout time.Duration `json:"timeout"`

	// Database is the default database name used when the Session.DB method
	// is called with an empty name, and is also used during the intial
	// authenticatoin if Source is unset.
	Database string `json:"database"`

	// Username and Password inform the credentials for the initial authentication
	// done on the database defined by the Source field.
	Username string `json:"username"`
	Password string `json:"password"`

	// collection name for chillingeffects.Notices
	NoticeCollectionName string `json:"notice_collection_name"`
}

type Config struct {
	MongoDB *MgoDialInfo `json:"mongodb"`
	RunMode string       `json:"runmode"`
	IDRange *IDRange     `json:"id_range"`
}

type IDRange struct {
	Low  int `json:"low"`
	High int `json:"high"`
}

func NewConfig(data string) (*Config, error) {
	config := &Config{}
	err := json.Unmarshal([]byte(data), config)
	return config, err
}

func LoadConfig(fileName string) (*Config, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	return config, err
}
