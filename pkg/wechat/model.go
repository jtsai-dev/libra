/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 16:16:45
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-15 11:29:50
 */
package wechat

type Client struct {
	AppId     string
	AppSecret string
}

type ErrorBase struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type WxToken struct {
	ErrorBase
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
}

type WxUserInfo struct {
	ErrorBase
	OpenId         string   `json:"openid"`
	Nickname       string   `json:"nickname"`
	Sex            int      `json:"sex"`
	Province       string   `json:"province"`
	City           string   `json:"city"`
	Country        string   `json:"country"`
	HeadImgurl     string   `json:"headimgurl"`
	Privilege      []string `json:"privilege"`
	UnionId        string   `json:"unionid"`
	Subscribe      int      `json:"subscribe"`
	Language       string   `json:"language"`
	SubscribeTime  int64    `json:"subscribe_time"`
	Remark         string   `json:"remark"`
	SubscribeScene string   `json:"subscribe_scene"`
	QrScene        int      `json:"qr_scene"`
	QrSceneStr     string   `json:"qr_scene_str"`
}

// ---- XML ----
/*
<xml>
    <URL><![CDATA[http://yyeiei.eicp.net/api/wechat]]></URL>
    <ToUserName><![CDATA[gh_1a17b789776d]]></ToUserName>
    <FromUserName><![CDATA[og9sYuMGhdgJ_G8Yki_MZzubcm6w]]></FromUserName>
    <CreateTime>1565603897</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[subscribe]]></Event>
    <Latitude></Latitude>
    <Longitude></Longitude>
    <Precision></Precision>
    <MsgId>0</MsgId>
</xml>
*/
type WxEvent struct {
	Url          string `xml:"URL"`
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	// EventKey     string  `xml:"EventKey"` // 事件KEY值，是一个32位无符号整数，即创建二维码时的二维码scene_id
	Latitude  float64 `xml:"Latitude"`
	Longitude float64 `xml:"Longitude"`
	Precision float64 `xml:"Precision"`
	MsgId     int64   `xml:"MsgId"`
}
