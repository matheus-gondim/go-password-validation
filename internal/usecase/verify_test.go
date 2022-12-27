package usecase

import (
	"testing"

	"github.com/matheus-gondim/go-password-validation/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestVerify_Exec(t *testing.T) {
	type args struct {
		verifyPassword entities.VerifyPassword
	}
	tests := []struct {
		name string
		args args
		want entities.Verified
	}{
		{
			name: "should return true when the password is valid ",
			args: args{
				verifyPassword: entities.VerifyPassword{
					Password: "TesteSenhaForte!1234&",
					Rules: []entities.Rule{
						{Rule: "minSize", Value: 8},
						{Rule: "minSpecialChars", Value: 2},
						{Rule: "noRepeted", Value: 0},
						{Rule: "minDigit", Value: 4},
						{Rule: "minUppercase", Value: 3},
						{Rule: "minLowercase", Value: 5},
					},
				},
			},
			want: entities.Verified{
				Verify:  true,
				NoMatch: []string{},
			},
		},
		{
			name: "should return false when the password is not valid ",
			args: args{
				verifyPassword: entities.VerifyPassword{
					Password: "esteSS!14",
					Rules: []entities.Rule{
						{Rule: "minSize", Value: 10},
						{Rule: "minSpecialChars", Value: 2},
						{Rule: "noRepeted", Value: 0},
						{Rule: "minDigit", Value: 4},
						{Rule: "minUppercase", Value: 3},
						{Rule: "minLowercase", Value: 5},
					},
				},
			},
			want: entities.Verified{
				Verify:  false,
				NoMatch: []string{"minSize", "minSpecialChars", "noRepeted", "minDigit", "minUppercase", "minLowercase"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Verify{}
			resp, err := v.Exec(tt.args.verifyPassword)
			assert.Empty(t, err)
			assert.Equal(t, tt.want.NoMatch, resp.NoMatch)
			assert.Equal(t, tt.want.Verify, resp.Verify)
		})
	}
}
