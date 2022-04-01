/**
 * @author zhaiyuanji
 * @date 2022年02月18日 10:09 下午
 */
package doctor

import (
	"OpenSchedule/src/database"
	"testing"
)

var dao = NewDoctorDao(database.GetElasticSearchEngine(), database.GetMySqlEngine())
func TestDao_SearchDoctor(t *testing.T) {
	dao.SyncInsurance()
}
