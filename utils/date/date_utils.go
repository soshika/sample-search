package date

import (
	"fmt"
	"github.com/leekchan/timeutil"
	"github.com/soshika/sample-search/logger"
	"strconv"
	"strings"
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

var (
	DateTimeService dateTimeServiceInterface = &dateTimesService{}
)

type dateTimesService struct{}

type dateTimeServiceInterface interface {
	GetNow() time.Time
	GetNowString() string
	GetNowDBFormat() string
	ConvertToDate(string) time.Time
	Date(int, int, int) time.Time
	DeltaTime(string, string) (*int, error)
	DateIsPassed(string) error
	IsExpired(string) error
	IsNotExpired(string) error
}

func (s *dateTimesService) GetNow() time.Time {
	return time.Now().UTC()
}

func (s *dateTimesService) GetNowString() string {
	return s.GetNow().Format(apiDateLayout)
}

func (s *dateTimesService) Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func (s *dateTimesService) DeltaTime(checkInDate string, checkOutDate string) (*int, error) {

	checkIn := strings.Split(checkInDate, "-")

	checkInDay, _ := strconv.ParseInt(checkIn[2], 10, 64)
	checkInMonth, _ := strconv.ParseInt(checkIn[1], 10, 64)
	checkInYear, _ := strconv.ParseInt(checkIn[0], 10, 64)

	checkOut := strings.Split(checkOutDate, "-")

	checkOutDay, _ := strconv.ParseInt(checkOut[2], 10, 64)
	checkOutMonth, _ := strconv.ParseInt(checkOut[1], 10, 64)
	checkOutYear, _ := strconv.ParseInt(checkOut[0], 10, 64)

	t1 := s.Date(int(checkInYear), int(checkInMonth), int(checkInDay))
	t2 := s.Date(int(checkOutYear), int(checkOutMonth), int(checkOutDay))
	days := int(t2.Sub(t1).Hours() / 24)

	return &days, nil
}

func (s *dateTimesService) DateIsPassed(date string) error {
	base, err := time.Parse("2006-01-02", date)
	if err != nil {
		logger.Error("error when trying to parse date", err)
		return err
	}

	oneMinutesLater := timeutil.Timedelta{
		Days:         1,
		Seconds:      0,
		Microseconds: 0,
		Milliseconds: 0,
		Minutes:      0,
		Hours:        0,
		Weeks:        0,
	}

	base = base.Add(oneMinutesLater.Duration())
	current := time.Now().UTC()
	result := current.Sub(base).Minutes()

	if result >= 1 {
		logger.Error("date is passed", err)
		return err
	}
	return nil
}

func (s *dateTimesService) IsExpired(dateVerified string) error {
	logger.Info("Enter to Is-Expired function")

	dateVerified = strings.Replace(dateVerified, " ", "T", -1)
	dateVerified += "Z"

	base, err := time.Parse(apiDateLayout, dateVerified)
	if err != nil {
		logger.Error("error when trying to parse date", err)
		return err
	}

	oneMinutesLater := timeutil.Timedelta{
		Days:         0,
		Seconds:      0,
		Microseconds: 0,
		Milliseconds: 0,
		Minutes:      1,
		Hours:        0,
		Weeks:        0,
	}

	base = base.Add(oneMinutesLater.Duration())
	current := time.Now().UTC()
	result := current.Sub(base).Minutes()

	if result >= 1 {
		logger.Error("verification code is expired", err)
		return err
	}
	logger.Info("Close from Is-Expired function successfully!")
	return nil
}

func (s *dateTimesService) IsNotExpired(dateVerified string) error {
	logger.Info("Enter to Is-Not-Expired function")

	dateVerified = strings.Replace(dateVerified, " ", "T", -1)
	dateVerified += "Z"

	base, err := time.Parse(apiDateLayout, dateVerified)
	if err != nil {
		logger.Error("error when trying to parse date ", err)
		return err
	}

	oneMinutesLater := timeutil.Timedelta{
		Days:         0,
		Seconds:      0,
		Microseconds: 0,
		Milliseconds: 0,
		Minutes:      1,
		Hours:        0,
		Weeks:        0,
	}

	base = base.Add(oneMinutesLater.Duration())
	current := time.Now().UTC()
	result := current.Sub(base).Minutes()

	if result < 1 {
		logger.Error("verification code is not expired", err)
		return err
	}
	logger.Info("Close from Is-Not-Expired function successfully")
	return nil
}

func (s *dateTimesService) GetNowDBFormat() string {
	ret := s.GetNow().Format(apiDbLayout)
	return ret
}

func (s *dateTimesService) ConvertToDate(dateTime string) time.Time {
	t, err := time.Parse(apiDateLayout, dateTime)

	if err != nil {
		fmt.Println(err)
	}
	return t
}
