package middlewaresUsecases

import _pkgMiddlewaresMiddlewaresRepositories "github.com/MarkTBSS/061_Sign_Up/modules/middlewares/middlewaresRepositories"

type IMiddlewaresUsecase interface {
}

type middlewaresUsecase struct {
	middlewaresRepository _pkgMiddlewaresMiddlewaresRepositories.IMiddlewaresRepository
}

func MiddlewaresUsecase(middlewaresRepository _pkgMiddlewaresMiddlewaresRepositories.IMiddlewaresRepository) IMiddlewaresUsecase {
	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}
