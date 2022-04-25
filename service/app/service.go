package app

import (
	"context"
	"gitee.com/langya_platform/langya/platform/app"
	"gitee.com/langya_platform/langya/platform/common"
	"gitee.com/langya_platform/pkg/xzap"
	"go.uber.org/zap"
)

type LangyaPlatformService struct {
	app.UnimplementedLangYaPlatformServer
}

func (*LangyaPlatformService) ServiceContractBook(
	ctx context.Context, req *app.ContractBookReq,
) (
	*app.ContractBook, error,
) {
	xzap.Info("contractBook", zap.Any("book", 123))
	panic(any("出错了"))
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
