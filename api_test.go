package etherscan

import (
	"fmt"
	"testing"
)

func TestGetBalance(t *testing.T) {
	//s, e := GetBlocksMined("QK93QKZ9R3P9SW3D3J6IRGACWQG2UT2H7B", "0x9dd134d14d1e65f84b706d6f205cd5b1cd03a46b", "", "", "")
	
	s, e := GetContractCreator( "QK93QKZ9R3P9SW3D3J6IRGACWQG2UT2H7B","0x3267c5b73cc15e253b1a90c01366b17d560bc6fb")
	fmt.Println(s,e)
}

func TestParseUrl(t *testing.T) {
	type args struct {
		url    string
		params []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseUrl(tt.args.url, tt.args.params...); got != tt.want {
				t.Errorf("ParseUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_get(t *testing.T) {
	type args struct {
		url string
		v   any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := get(tt.args.url, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
