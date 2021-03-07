package domain

import (
	"github.com/yiranzai/github-starred/domain/model"
	"github.com/yiranzai/github-starred/util"
)

// NewNumbersFromStringSlice create new Numbers with numbers from string slice
func NewNumbersFromStringSlice(strNumbers []string) (model.Numbers, error) {
	rawNumbers, err := util.ConvertStringSliceToIntSlice(strNumbers)
	if err != nil {
		return nil, err
	}
	return model.NewNumbers(rawNumbers), nil
}
