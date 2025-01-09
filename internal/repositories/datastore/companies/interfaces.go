package companies

import "context"

type CompanyRepository interface {
	CreateNewRecord(ctx context.Context, input CreateNewRecordInput) (output CreateNewRecordOutput, err error)
	GetAll(ctx context.Context, input GetAllInput) (output GetAllOutput, err error)
	GetAllCache(ctx context.Context, input GetAllCacheInput) (output GetAllCacheOutput, err error)
	SetCache(ctx context.Context, input SetCacheInput) (err error)
	ClearCache(ctx context.Context) (err error)
}
