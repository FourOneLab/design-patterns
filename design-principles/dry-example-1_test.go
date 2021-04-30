package design_principles

import "testing"

func TestUserAuthenticator_authenticate(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"1",
			args{
				username: "dsad09.",
				password: "djsad.0",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserAuthenticator{}
			if err := u.Authenticate(tt.args.username, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("authenticate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
