/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 14:17:19
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-13 10:25:17
 */
package enums

const (
	UnSubscribed = iota
	Subscribed
)

var subscribeStatusFlags = map[int]string{
	Subscribed:   "Subscribed",
	UnSubscribed: "UnSubscribed",
}

func GetSubscribeStatusDesc(value int) string {
	desc, ok := subscribeStatusFlags[value]
	if ok {
		return desc
	}

	return subscribeStatusFlags[Fail]
}
