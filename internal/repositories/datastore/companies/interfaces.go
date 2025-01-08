package companies

import "context"

type CompanyRepository interface {
	CreateNewRecord(ctx context.Context, input CreateNewRecordInput) (output CreateNewRecordOutput, err error)
	GetAll(ctx context.Context, input GetAllInput) (output GetAllOutput, err error)
}
