package main

import (
	"fmt"
	"testing"
)

func Test_getUserFromGit(t *testing.T) {
	type args struct {
		userName string
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{"validgituser", args{userName: "rajatparida86"}, nil},
		{"emptygituser", args{userName: ""}, fmt.Errorf("Failed to fetch from Github. Github returned with Status code: 404 and Message: Not Found")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getUserFromGit(tt.args.userName)
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("getUserFromGit() ERRORs:  \nGot = %v \nExpected = %v", err, tt.err)
				}
			} else {
				if err != tt.err {
					t.Errorf("getUserFromGit() ERRORs:  \nGot = %v \nExpected = %v", err, tt.err)
				}
			}
		})
	}
}
