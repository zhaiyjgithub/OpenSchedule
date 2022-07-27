/**
 * @author zhaiyuanji
 * @date 2021年11月17日 2:50 下午
 */
package job

import (
	"OpenSchedule/service/doctorService"
	"OpenSchedule/service/scheduleService"
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

const Len = 100

type SyncScheduleJob struct {
	doctorService   doctorService.Service
	scheduleService scheduleService.Service
}

func NewJob() *SyncScheduleJob {
	return &SyncScheduleJob{}
}

func (j *SyncScheduleJob) RegisterService(doctorService doctorService.Service, scheduleService scheduleService.Service) {
	j.doctorService = doctorService
	j.scheduleService = scheduleService
}

func (j *SyncScheduleJob) Test() {
	s := gocron.NewScheduler(time.UTC)

	counter := 0
	s.Every(1).Seconds().Do(func() {
		counter = counter + 1
		fmt.Println("time: ", counter)
	})
	s.StartAsync()
}

func (j *SyncScheduleJob) StartToSyncDoctorNextAvailableDateAsync() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Days().At("16:00").Do(j.SyncDoctorsNextAvailableDate)
	s.StartAsync()
	fmt.Println("Set the scheduler to sync the next available date for all doctors. ")
}

func (j *SyncScheduleJob) SyncDoctorsNextAvailableDate() {
	page := 1
	pageSize := 100

	startTime := time.Now().UTC()
	fmt.Println("Start to Sync Doctors Next Available Date", startTime.Format(time.RFC3339))
	for {
		doctors := j.doctorService.GetDoctorByPage(page, pageSize)
		err := j.scheduleService.SyncMultiDoctorsScheduleNextAvailableDateToES(doctors)
		if err != nil {
			fmt.Println("bulk update error: ", err)
		}
		fmt.Println("page : ", page)
		page = page + 1
		if len(doctors) < pageSize {
			fmt.Println("The syn task is completed.")
			break
		}
	}
	endTime := time.Now().UTC()
	fmt.Println("End the sync process", endTime.Format(time.RFC3339))

	duration := endTime.Sub(startTime)
	fmt.Println("Duration: ", duration)
}

func (j *SyncScheduleJob) SyncDefaultScheduleSettingsToAllDoctor() {
	testNpi := int64(1902809254)
	defaultSettings := j.scheduleService.GetScheduleSettings(testNpi)
	page := 1
	pageSize := 100

	fmt.Println("begin to sync", time.Now().UTC().Format(time.RFC3339))
	for {
		doctors := j.doctorService.GetDoctorByPage(page, pageSize)
		for _, doc := range doctors {
			defaultSettings.ID = 0
			defaultSettings.Npi = doc.Npi
			err := j.scheduleService.SetScheduleSettings(defaultSettings)
			if err != nil {
				fmt.Println("save schedule settings err: ", err.Error())
			}
		}
		fmt.Println("page : ", page)
		page = page + 1
		if len(doctors) < pageSize {
			fmt.Println("The sync task is completed.")
			break
		}
	}
	fmt.Println("end the sync", time.Now().UTC().Format(time.RFC3339), "page: ", page)
}
