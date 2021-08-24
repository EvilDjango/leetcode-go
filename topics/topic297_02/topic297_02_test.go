package topic297_02

import (
	"leetcode-go/tree"
	"reflect"
	"testing"
)

/*
Oops, forgot to write comments. Good luck, bro.

Copyright (c) @2021 deerhunter0837@gmail.com All Rights Reserved.
@author xuejunc
@createTime   8/23/21 7:52 PM
*/

func TestCodec_serialize_and_deserialize(t *testing.T) {
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"",
			args{tree.New(1, 2, 3, 4, tree.None, 5, 6)},
			"1|2|4||||3|5|||6||",
		},
		{
			"",
			args{tree.New(1, 2, 3)},
			"1|2|||3||",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.serialize(tt.args.root); got != tt.want {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
			if got := this.deserialize(tt.want); !reflect.DeepEqual(got, tt.args.root) {
				t.Errorf("deserialize() = %v, want %v", got, tt.args.root)
			}
		})
	}
}
