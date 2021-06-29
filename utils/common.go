package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func ValidatePasswordPolices(password string) error {
	if len(password) < 8 {
		return errors.New("La contraseña debe contener como minimo 8 caracteres")
	}
	needAlphabetic := true
	needNumber := true
	needSymbol := true
	for _, i := range password {
		if unicode.IsLetter(i) {
			needAlphabetic = false
		} else if unicode.IsDigit(i) {
			needNumber = false
		} else {
			needSymbol = false
		}
	}
	if needAlphabetic {
		return errors.New("La contraseña debe contener como minimo una letra")
	}
	if needNumber {
		return errors.New("La contraseña debe contener como minimo un numero")
	}
	if needSymbol {
		return errors.New("La contraseña debe contener como minimo un simbolo")
	}

	return nil

}

func CleanEmail(email string) string {
	email = strings.Replace(email, " ", "", -1)
	email = strings.ToLower(email)
	return email
}

func CleanPhoneNumber(phoneNumber string) string {
	phoneNumber = strings.Replace(phoneNumber, " ", "", -1)
	phoneNumber = strings.Replace(phoneNumber, "-", "", -1)
	return phoneNumber
}

func GetPrevPage(page int) int {
	if page > 1 {
		return page - 1
	}
	return 1
}

func GetNextPage(page int, totalPage int) int {
	if page > totalPage {
		return totalPage
	}
	if page != totalPage {
		return page + 1
	}
	return page
}

func SqlNullStringToString(val sql.NullString) string {
	response, err := val.Value()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s", response)
}

func ConvertUnit8ToArray(i *uint8) []uint8 {
	var response []uint8
	if i != nil {
		response = append(response, *i)
	}
	return response
}

func ConvertArraytoUint8(i []uint8) *uint8 {
	var response uint8
	if len(i) > 0 {
		response = uint8(i[0])
	}
	return &response
}
