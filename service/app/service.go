package app

import (
	"context"
	"gitee.com/langya_platform/langya/platform/app"
	"gitee.com/langya_platform/langya/platform/common"
)

type LangyaPlatformService struct {
	app.UnimplementedLangYaPlatformServer
}

func (*LangyaPlatformService) ServiceContractBook(
	ctx context.Context, req *app.ContractBookReq,
) (
	*app.ContractBook, error,
) {
	return &app.ContractBook{
		Persons: []*app.Person{
			&app.Person{
				Id:     1,
				Name:   "demo",
				Gender: common.GenderType_MALE,
				Number: "abc123",
			},
		},
	}, nil
}
