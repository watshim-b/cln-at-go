package usecase

import (
	form "github.com/watshim-b/cln-at-go/form/web"
)

type LoginUsecase interface {

	// ログイン機能を提供します
	Login(f form.LoginForm) error
	
}
