package env

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
	"testing"
	"time"
)

func TestEnvBool(t *testing.T) {
	os.Setenv("PROP", "true")

	config := struct {
		Prop bool `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, true, config.Prop)
}

func TestEnvIntegers(t *testing.T) {
	os.Setenv("PROP", "123")

	config := struct {
		PropInt   int   `env:"PROP"`
		PropInt8  int8  `env:"PROP"`
		PropInt16 int16 `env:"PROP"`
		PropInt32 int32 `env:"PROP"`
		PropInt64 int64 `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, int(123), config.PropInt)
	Equals(t, int8(123), config.PropInt8)
	Equals(t, int16(123), config.PropInt16)
	Equals(t, int32(123), config.PropInt32)
	Equals(t, int64(123), config.PropInt64)
}

func TestIntegerRanges(t *testing.T) {
	testCases := []struct {
		Prop8  int8
		Prop16 int16
		Prop32 int32
		Prop64 int64
	}{
		{Prop8: math.MinInt8, Prop16: math.MinInt16, Prop32: math.MinInt32, Prop64: math.MinInt64},
		{Prop8: math.MaxInt8, Prop16: math.MaxInt16, Prop32: math.MaxInt32, Prop64: math.MaxInt64},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase), func(t *testing.T) {
			os.Setenv("PROP8", fmt.Sprintf("%d", testCase.Prop8))
			os.Setenv("PROP16", fmt.Sprintf("%d", testCase.Prop16))
			os.Setenv("PROP32", fmt.Sprintf("%d", testCase.Prop32))
			os.Setenv("PROP64", fmt.Sprintf("%d", testCase.Prop64))

			config := struct {
				Prop8  int8  `env:"PROP8"`
				Prop16 int16 `env:"PROP16"`
				Prop32 int32 `env:"PROP32"`
				Prop64 int64 `env:"PROP64"`
			}{}

			ErrorNil(t, Set(&config))
			Equals(t, testCase.Prop8, config.Prop8)
			Equals(t, testCase.Prop16, config.Prop16)
			Equals(t, testCase.Prop32, config.Prop32)
			Equals(t, testCase.Prop64, config.Prop64)
		})
	}
}

func TestUnsignedIntegerRanges(t *testing.T) {
	testCases := []struct {
		Prop8  uint8
		Prop16 uint16
		Prop32 uint32
		Prop64 uint64
	}{
		{Prop8: 0, Prop16: 0, Prop32: 0, Prop64: 0},
		{Prop8: math.MaxUint8, Prop16: math.MaxUint16, Prop32: math.MaxUint32, Prop64: math.MaxUint64},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase), func(t *testing.T) {
			os.Setenv("PROP8", fmt.Sprintf("%d", testCase.Prop8))
			os.Setenv("PROP16", fmt.Sprintf("%d", testCase.Prop16))
			os.Setenv("PROP32", fmt.Sprintf("%d", testCase.Prop32))
			os.Setenv("PROP64", fmt.Sprintf("%d", testCase.Prop64))

			config := struct {
				Prop8  uint8  `env:"PROP8"`
				Prop16 uint16 `env:"PROP16"`
				Prop32 uint32 `env:"PROP32"`
				Prop64 uint64 `env:"PROP64"`
			}{}

			ErrorNil(t, Set(&config))
			Equals(t, testCase.Prop8, config.Prop8)
			Equals(t, testCase.Prop16, config.Prop16)
			Equals(t, testCase.Prop32, config.Prop32)
			Equals(t, testCase.Prop64, config.Prop64)
		})
	}
}

func TestUnsignedFloatRanges(t *testing.T) {
	testCases := []struct {
		Prop32 float32
		Prop64 float64
	}{
		{Prop32: math.SmallestNonzeroFloat32, Prop64: math.SmallestNonzeroFloat64},
		{Prop32: math.MaxFloat32, Prop64: math.MaxFloat64},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase), func(t *testing.T) {
			os.Setenv("PROP32", fmt.Sprintf("%g", testCase.Prop32))
			os.Setenv("PROP64", fmt.Sprintf("%g", testCase.Prop64))

			config := struct {
				Prop32 float32 `env:"PROP32"`
				Prop64 float64 `env:"PROP64"`
			}{}

			ErrorNil(t, Set(&config))
			Equals(t, testCase.Prop32, config.Prop32)
			Equals(t, testCase.Prop64, config.Prop64)
		})
	}
}

func TestEnvUnsignedIntegers(t *testing.T) {
	os.Setenv("PROP", "123")

	config := struct {
		PropUint   uint   `env:"PROP"`
		PropUint8  uint8  `env:"PROP"`
		PropUint16 uint16 `env:"PROP"`
		PropUint32 uint32 `env:"PROP"`
		PropUint64 uint64 `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, uint(123), config.PropUint)
	Equals(t, uint8(123), config.PropUint8)
	Equals(t, uint16(123), config.PropUint16)
	Equals(t, uint32(123), config.PropUint32)
	Equals(t, uint64(123), config.PropUint64)
}

func TestEnvFloats(t *testing.T) {
	os.Setenv("PROP", "1.23")

	config := struct {
		PropFloat32 float32 `env:"PROP"`
		PropFloat64 float64 `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, float32(1.23), config.PropFloat32)
	Equals(t, float64(1.23), config.PropFloat64)
}

func TestEnvString(t *testing.T) {
	os.Setenv("PROP", "}D-Z2P£T!E*#zE=.gc@")

	config := struct {
		Prop string `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, "}D-Z2P£T!E*#zE=.gc@", config.Prop)
}

func TestEnvSetUnexportedProperty(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		prop string `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Assert(t, strings.Contains(err.Error(), "field cannot be set"))
}

func TestInvalidValueForRequiredTag(t *testing.T) {
	os.Unsetenv("PROP")

	config := struct {
		Prop int `env:"PROP" required:"invalid"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Assert(t, strings.HasPrefix(err.Error(), "invalid required tag 'invalid'"))
}

func TestEnvNoEnvTag(t *testing.T) {
	config := struct {
		Prop string
	}{}

	ErrorNil(t, Set(&config))
}

func TestEnvRequiredWhenProvided(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop string `env:"PROP" required:"true"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, "hello", config.Prop)
}

func TestEnvRequiredWhenMissing(t *testing.T) {
	config := struct {
		Prop string `env:"MISSING_PROP" required:"true"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
}

func TestEnvNotRequiredImplicitWhenMissing(t *testing.T) {
	os.Unsetenv("PROP")

	config := struct {
		Prop string `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNil(t, err)
}

func TestEnvNotRequiredExplicitWhenMissing(t *testing.T) {
	os.Unsetenv("PROP")

	config := struct {
		Prop string `env:"PROP" required:"false"`
	}{}

	err := Set(&config)
	ErrorNil(t, err)
}

func TestInvalidConfigurationForBoolType(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop bool `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Assert(t, strings.HasPrefix(err.Error(), "error setting Prop"))
}

func TestInvalidConfigurationForIntType(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop int `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Assert(t, strings.HasPrefix(err.Error(), "error setting Prop"))
}

func TestInvalidConfigurationForUintType(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop uint `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Assert(t, strings.HasPrefix(err.Error(), "error setting Prop"))
}

func TestInvalidConfigurationForFloatType(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop float32 `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Assert(t, strings.HasPrefix(err.Error(), "error setting Prop"))
}

func TestEnvNonPointer(t *testing.T) {
	config := struct {
		Prop float32 `env:"PROP"`
	}{}

	err := Set(config)
	ErrorNotNil(t, err)
	Equals(t, err.Error(), "struct is not a pointer")
}

func TestEnvCustomTypeAliasedPrimativeWithoutSetter(t *testing.T) {
	os.Setenv("PROP", "1234")

	config := struct {
		Prop myInt `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, myInt(1234), config.Prop)
}

type myInt int16

func TestEnvCustomTypeAliasedPrimativeWithSetter(t *testing.T) {
	t.SkipNow()

	os.Setenv("PROP", "1234")

	config := struct {
		Prop stringSlice `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	fmt.Println(config.Prop)
	Equals(t, stringSlice{"a", "b", "c"}, config.Prop)
}

type stringSlice []string

func (s stringSlice) Set(value string) (err error) {
	slice := strings.Split(value, ",")
	s = stringSlice(slice)
	return
}

func TestEnvCustomTypeStruct(t *testing.T) {
	os.Setenv("PROP", "3h2m1s")

	config := struct {
		Timeout *configDuration `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, time.Hour*3+time.Minute*2+time.Second*1, config.Timeout.Duration)
}

func TestEnvCustomTypeStructWithError(t *testing.T) {
	os.Setenv("PROP", "3h2m1s")

	config := struct {
		Timeout *configDurationError `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Assert(t, strings.HasPrefix(err.Error(), "error in custom setter"))
	Assert(t, strings.Contains(err.Error(), errConfigDurationError.Error()))
}

type configDuration struct {
	Duration time.Duration
}

func (d *configDuration) Set(config string) (err error) {
	d.Duration, err = time.ParseDuration(config)
	return
}

type configDurationError struct {
	Duration time.Duration
}

var errConfigDurationError = errors.New("example error from custom Set code")

func (d *configDurationError) Set(config string) (err error) {
	return errConfigDurationError
}