/**
 * @author zhaiyuanji
 * @date 2022年02月22日 3:08 下午
 */
package doctorModel

type TimeSlot struct {
	Offset               int `json:"offset"`
	AvailableSlotsNumber int `json:"availableSlotsNumber"`
}
