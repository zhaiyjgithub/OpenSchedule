/**
 * @author zhaiyuanji
 * @date 2021年11月18日 4:11 下午
 */
package job

import (
	"OpenSchedule/service/doctorService"
	"OpenSchedule/service/scheduleService"
	"testing"
)

func TestSyncScheduleJob_SyncDefaultScheduleSettingsToAllDoctor(t *testing.T) {
	doctorService := doctorService.NewService()
	scheduleService := scheduleService.NewService()
	j := NewJob()
	j.RegisterService(doctorService, scheduleService)
	j.SyncDefaultScheduleSettingsToAllDoctor()
}
