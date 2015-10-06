package valid

import (
	"errors"
	"gopkg.in/validator.v2"
)

const (
	IsErr   string = "IS_ERR" //エラーが有るかどうか
	Nonzero string = "NZ"     //空白じゃないかどうか
	Min2    string = "MI2"    //最低２文字
	Min4    string = "MI4"    //最低４文字
	Max2    string = "MA2"    //最高２文字
	Max12   string = "MA12"   //最高１２文字
	Max18   string = "MA18"   //最高18文字
	Ran     string = "RAN"    //数字とアルファベット
	Ra      string = "RA"     //アルファベット
	Rn      string = "RN"     //数字
	Bc      string = "BC"     //0 1　どちらか
)

//nonzero  NZ
type ch_nonzero struct {
	T string `validate:"nonzero"`
}

//min 4  more than 4     MI4
type ch_min4 struct {
	T string `validate:"min=4"`
}

type ch_min2 struct {
	T string `validate:"min=2"`
}

//max 12 less than 12   MA12
type ch_max12 struct {
	T string `validate:"max=12"`
}

type ch_max2 struct {
	T string `validate:"max=2"`
}

type ch_max18 struct {
	T string `validate:"max=18"`
}

//regexp allow alphanumeric and _  RAN
type ch_regexpAlphanumeric_ struct {
	T string `validate:"regexp=^[a-zA-Z0-9_]+$"`
}

//regexp alphabet     RA
type ch_regexpAlphabets struct {
	T string `validate:"regexp=^[a-zA-Z]+$"`
}

//regexp number       RN
type ch_regexpNumber struct {
	T string `validate:"regexp=^[0-9]+$"`
}

//Binary check 0 or 1 BC
type ch_binary struct {
	T string `validate:"^[01]$"`
}

//return final check data
func SetError(checked string, strings ...string) string {
	var finalErr error = nil
	for _, s := range strings {
		switch s {
		case Nonzero:
			nur := ch_nonzero{T: checked}
			v := validator.Validate(nur)
			if v != nil {
				finalErr = v
			}
		case Min2:
			nur := ch_min2{T: checked}
			v := validator.Validate(nur)
			if v != nil {
				finalErr = v
			}
		case Min4:
			nur := ch_min4{T: checked}
			v := validator.Validate(nur)
			if v != nil {
				finalErr = v
			}
		case Max2:
			nur := ch_max2{T: checked}
			v := validator.Validate(nur)
			if v != nil {
				finalErr = v
			}
		case Max12:
			nur := ch_max12{T: checked}
			v := validator.Validate(nur)
			if v != nil {
				finalErr = v
			}
		case Max18:
			nur := ch_max18{T: checked}
			v := validator.Validate(nur)
			if v != nil {
				finalErr = v
			}
		case Ran:
			nur := ch_regexpAlphanumeric_{T: checked}
			v := validator.Validate(nur)
			if v != nil {
				finalErr = v
			}
		case Ra:
			nur := ch_regexpAlphabets{T: checked}
			v := validator.Validate(nur)
			if v != nil {
				finalErr = v
			}
		case Rn:
			nur := ch_regexpNumber{T: checked}
			v := validator.Validate(nur)
			if v != nil {
				finalErr = v
			}
		case Bc:
			nur := ch_binary{T: checked}
			v := validator.Validate(nur)
			if v != nil {
				finalErr = errors.New("Out of range")
			}
		}
	}
	if finalErr != nil {
		return finalErr.Error()
	} else {
		return ""
	}
}

//set every error data
func SetErrors(checked string, err map[string]error, strings ...string) error {
	var finalErr error = nil
	for _, s := range strings {
		switch s {
		case Nonzero:
			nur := ch_nonzero{T: checked}
			v := validator.Validate(nur)
			err[Nonzero] = v
			finalErr = v
		case Min2:
			nur := ch_min2{T: checked}
			v := validator.Validate(nur)
			err[Min2] = v
			finalErr = v
		case Min4:
			nur := ch_min4{T: checked}
			v := validator.Validate(nur)
			err[Min4] = v
			finalErr = v
		case Max2:
			nur := ch_max2{T: checked}
			v := validator.Validate(nur)
			err[Max2] = v
			finalErr = v
		case Max12:
			nur := ch_max12{T: checked}
			v := validator.Validate(nur)
			err[Max12] = v
			finalErr = v
		case Max18:
			nur := ch_max18{T: checked}
			v := validator.Validate(nur)
			err[Max18] = v
			finalErr = v
		case Ran:
			nur := ch_regexpAlphanumeric_{T: checked}
			v := validator.Validate(nur)
			err[Ran] = v
			finalErr = v
		case Ra:
			nur := ch_regexpAlphabets{T: checked}
			v := validator.Validate(nur)
			err[Ra] = v
			finalErr = v
		case Rn:
			nur := ch_regexpNumber{T: checked}
			v := validator.Validate(nur)
			err[Rn] = v
			finalErr = v
		case Bc:
			nur := ch_binary{T: checked}
			v := validator.Validate(nur)
			err[Bc] = v
			finalErr = v
		}
		err[IsErr] = finalErr
	}
	if finalErr != nil {
		return finalErr
	} else {
		return nil
	}
}
