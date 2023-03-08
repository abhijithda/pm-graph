// Copyright (c) 2021 Veritas Technologies LLC. All rights reserved. IP63-2828-7171-04-15-9
package pg

import (
	"os"
	"testing"

	"github.com/VeritasOS/plugin-manager/pluginmanager"
)

func Test_getStatusColor(t *testing.T) {
	if os.Getenv("INTEGRATION_TEST") == "RUNNING" {
		t.Skip("Not applicable while running integration tests.")
		return
	}

	type args struct {
		status string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Start",
			args: args{status: pluginmanager.DStatusStart},
			want: "blue",
		},
		{
			name: "Ok/Pass",
			args: args{status: pluginmanager.DStatusOk},
			want: "green",
		},
		{
			name: "Fail",
			args: args{status: pluginmanager.DStatusFail},
			want: "red",
		},
		{
			name: "Skip",
			args: args{status: pluginmanager.DStatusSkip},
			want: "yellow",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStatusColor(tt.args.status); got != tt.want {
				t.Errorf("getStatusColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
