package connection

import (
	"io"
	"reflect"
	_ "strings"
	"testing"
	_ "time"

	"golang.org/x/crypto/ssh"
)

func TestConnection_PublicKeyFile(t *testing.T) {
	type fields struct {
		Host     string
		Password string
		Command  string
		Quiet    bool
		Timeout  int
	}
	type args struct {
		file string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ssh.AuthMethod
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				Host:     tt.fields.Host,
				Password: tt.fields.Password,
				Command:  tt.fields.Command,
				Quiet:    tt.fields.Quiet,
				Timeout:  tt.fields.Timeout,
			}
			if got := c.PublicKeyFile(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connection.PublicKeyFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnection_MakeSSHConfig(t *testing.T) {
	type fields struct {
		Host     string
		Password string
		Command  string
		Quiet    bool
		Timeout  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *ssh.ClientConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				Host:     tt.fields.Host,
				Password: tt.fields.Password,
				Command:  tt.fields.Command,
				Quiet:    tt.fields.Quiet,
				Timeout:  tt.fields.Timeout,
			}
			if got := c.MakeSSHConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connection.MakeSSHConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnection_Connect(t *testing.T) {
	type fields struct {
		Host     string
		Password string
		Command  string
		Quiet    bool
		Timeout  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *ssh.Session
		want1  io.Reader
		want2  io.Reader
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				Host:     tt.fields.Host,
				Password: tt.fields.Password,
				Command:  tt.fields.Command,
				Quiet:    tt.fields.Quiet,
				Timeout:  tt.fields.Timeout,
			}
			got, got1, got2 := c.Connect()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connection.Connect() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Connection.Connect() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Connection.Connect() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestConnection_ExecuteCmd(t *testing.T) {
	type fields struct {
		Host     string
		Password string
		Command  string
		Quiet    bool
		Timeout  int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				Host:     tt.fields.Host,
				Password: tt.fields.Password,
				Command:  tt.fields.Command,
				Quiet:    tt.fields.Quiet,
				Timeout:  tt.fields.Timeout,
			}
			if got := c.ExecuteCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connection.ExecuteCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
