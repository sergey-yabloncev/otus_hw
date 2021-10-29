package hw10programoptimization

import (
	"bufio"
	"errors"
	"io"
	"strings"

	"github.com/valyala/fastjson"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	if domain == "" {
		return nil, errors.New("domain cannot be empty")
	}
	scanner := bufio.NewScanner(r)
	result := make(DomainStat)

	for scanner.Scan() {
		email := fastjson.GetString(scanner.Bytes(), "Email")

		if strings.HasSuffix(email, "."+domain) {
			result[strings.ToLower(strings.SplitN(email, "@", 2)[1])]++
		}
	}
	return result, nil
}
