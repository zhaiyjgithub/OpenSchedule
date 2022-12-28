package scheduleService

import (
	"fmt"
	"testing"
	"time"
)

var testNpi = int64(1902809254)

func TestService_GetOneDayTimeSlotByNpi(t *testing.T) {
	s := NewService()
	ttime := time.Now().UTC()
	l, err := s.GetOneDayTimeSlotByNpi(testNpi, ttime)
	if err != nil {
		t.Errorf("test failed: %v\r\n", err.Error())
	} else {
		fmt.Println("Test success")
		for _, st := range l {
			fmt.Printf("%v\r\n", st)
		}
	}
}

func TestService_CheckTimeSlotIsAvailable(t *testing.T) {
	s := NewService()
	ttime := time.Now().Add(time.Hour * 6).UTC()
	ok, err := s.CheckTimeSlotIsAvailable(testNpi, ttime)
	if err != nil {
		t.Errorf("test failed: %v\r\n", err.Error())
	} else {
		fmt.Printf("Is available: %v\r\n", ok)
	}
}

func TestService_GetAppointment(t *testing.T) {
	s := NewService()
	appts, err := s.GetAppointmentInfo(2, 1, 10)
	if err != nil {
		t.Errorf("test failed: %v\r\n", err.Error())
	} else {
		fmt.Printf("appts: %v\r\n", appts)
	}
}
