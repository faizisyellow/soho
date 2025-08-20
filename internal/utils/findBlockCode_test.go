package utils

import (
	"strings"
	"testing"
)

func TestFindBlockCode(t *testing.T) {

	testsCases := []struct {
		name     string
		data     string
		term     string
		expected int
	}{
		{
			name: "should return the last interfaces of close block line",
			data: `
	type Service struct {
		UsersService interface {
				RegisterAccount(context.Context, RegisterRequest) (*RegisterResponse, error)
				ActivateAccount(context.Context, string) error
				Login(context.Context, LoginRequest) (*repository.UserModel, error)
				DeleteAccount(context.Context, int) error
				FindUserById(context.Context, int) (*repository.UserModel, error)
		}	
		UsersService interface {
		        RegisterAccount(context.Context, RegisterRequest) (*RegisterResponse, error)
				ActivateAccount(context.Context, string) error
				Login(context.Context, LoginRequest) (*repository.UserModel, error)
				DeleteAccount(context.Context, int) error
				FindUserById(context.Context, int) (*repository.UserModel, error)
		}
	}
			`,
			term:     "interface",
			expected: 16,
		},
		{
			name: "should return inner block code line",
			data: `
			type Service struct {
				UsersService interface {
					RegisterAccount(context.Context, RegisterRequest) (*RegisterResponse, error)
					ActivateAccount(context.Context, string) error
					Login(context.Context, LoginRequest) (*repository.UserModel, error)
					DeleteAccount(context.Context, int) error
					FindUserById(context.Context, int) (*repository.UserModel, error)
				}
			}
				
			func New(store repository.Repository, txfnc db.TransFnc, Db *sql.DB) *Service {
					return &Service{
						UsersService: &UsersServices{Repository: store, TransFnc: txfnc, Db: Db},
					}
			}
			`,
			term:     "return &Service",
			expected: 15,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			dat := strings.NewReader(tc.data)
			res, err := FindBlockCode(dat, tc.term)
			if err != nil {
				t.Error(err)
			}

			if res != tc.expected {
				t.Errorf("expected %v but got %v", tc.expected, res)
			}
		})
	}
}
