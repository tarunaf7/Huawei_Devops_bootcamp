

package money

import (
	"errors"

	pb "github.com/GoogleCloudPlatform/microservices-demo/src/frontend/genproto"
)

const (
	nanosMin = -999999999
	nanosMax = +999999999
	nanosMod = 1000000000
)

var (
	ErrInvalidValue        = errors.New("one of the specified money values is invalid")
	ErrMismatchingCurrency = errors.New("mismatching currency codes")
)


func IsValid(m pb.Money) bool {
	return signMatches(m) && validNanos(m.GetNanos())
}

func signMatches(m pb.Money) bool {
	return m.GetNanos() == 0 || m.GetUnits() == 0 || (m.GetNanos() < 0) == (m.GetUnits() < 0)
}

func validNanos(nanos int32) bool { return nanosMin <= nanos && nanos <= nanosMax }


func IsZero(m pb.Money) bool { return m.GetUnits() == 0 && m.GetNanos() == 0 }



func IsPositive(m pb.Money) bool {
	return IsValid(m) && m.GetUnits() > 0 || (m.GetUnits() == 0 && m.GetNanos() > 0)
}



func IsNegative(m pb.Money) bool {
	return IsValid(m) && m.GetUnits() < 0 || (m.GetUnits() == 0 && m.GetNanos() < 0)
}



func AreSameCurrency(l, r pb.Money) bool {
	return l.GetCurrencyCode() == r.GetCurrencyCode() && l.GetCurrencyCode() != ""
}



func AreEquals(l, r pb.Money) bool {
	return l.GetCurrencyCode() == r.GetCurrencyCode() &&
		l.GetUnits() == r.GetUnits() && l.GetNanos() == r.GetNanos()
}


func Negate(m pb.Money) pb.Money {
	return pb.Money{
		Units:        -m.GetUnits(),
		Nanos:        -m.GetNanos(),
		CurrencyCode: m.GetCurrencyCode()}
}



func Must(v pb.Money, err error) pb.Money {
	if err != nil {
		panic(err)
	}
	return v
}




func Sum(l, r pb.Money) (pb.Money, error) {
	if !IsValid(l) || !IsValid(r) {
		return pb.Money{}, ErrInvalidValue
	} else if l.GetCurrencyCode() != r.GetCurrencyCode() {
		return pb.Money{}, ErrMismatchingCurrency
	}
	units := l.GetUnits() + r.GetUnits()
	nanos := l.GetNanos() + r.GetNanos()

	if (units == 0 && nanos == 0) || (units > 0 && nanos >= 0) || (units < 0 && nanos <= 0) {
		// same sign <units, nanos>
		units += int64(nanos / nanosMod)
		nanos = nanos % nanosMod
	} else {
		// different sign. nanos guaranteed to not to go over the limit
		if units > 0 {
			units--
			nanos += nanosMod
		} else {
			units++
			nanos -= nanosMod
		}
	}

	return pb.Money{
		Units:        units,
		Nanos:        nanos,
		CurrencyCode: l.GetCurrencyCode()}, nil
}



func MultiplySlow(m pb.Money, n uint32) pb.Money {
	out := m
	for n > 1 {
		out = Must(Sum(out, m))
		n--
	}
	return out
}
