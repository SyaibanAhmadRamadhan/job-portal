package company

import "context"

type CompanyService interface {
	GetListCompany(ctx context.Context, input GetListCompanyInput) (output GetListCompanyOutput, err error)
}
