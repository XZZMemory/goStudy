package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 1.0 计划数据
	mpUrl1 := "sslocal://microapp?start_page=pages%2Fcalculator%2Findex%3FlandId%3D2313%26source%3D1100726054%26projectid%3D__PROJECT_ID__%26promotionid%3D__PROMOTION_ID__%26creativetype%3D__CTYPE__%26clickid%3D__CLICKID__&__ad_app_type__=2&version_type=current&app_id=tt396d4fd0a65382ed01&scene=0&version=v2&bdpsum=2852497"
	u1, _ := url.Parse(mpUrl1)
	q1 := u1.Query()

	// 2.0 计划数据
	mpUrl2 := "sslocal://microapp?start_page=pages%2Fcalculator%2Findex%3FlandId%3D2313%26mid1%3D0%26mid2%3D7197725821368221735%26mid3%3D7239313247488278566%26promotion_name%3D小赢2.0小程序兜底链接-抖音-手动-投放2.0-下载小程序-已申请-口播-还款-0602-36%%26promotionid%3D7239925482619666465%26requestid%3D20230609172333B2F0DA020EDF3DF959E0%26source%3D1100726055%26union_site%3D__UNION_SITE__%26vid%3D__VID__%26projectid%3D7239924572828155962%26promotionid%3D7239925482619666465%26creativetype%3D15%26clickid%3DB.4FeqvvAlOdfyrfLLyZVFMR6DbZWTmzy2Bsp8ZeUUfLepWIpnDBgE6f2RvvOsBbNNjn777qeQtFlWUGmnuTQaFpjIlyvIhcJWP5Ozsmf2k47eC&version_type=current&scene=0&app_id=tt396d4fd0a65382ed01&version=v2&bdpsum=2c763ec&bdp_log=%7B%22launch_from%22%3A%22ad%22%7D"
	u2, _ := url.Parse(mpUrl2)
	q2 := u2.Query()
	//

	mpUrl21 := "sslocal://microapp?start_page=pages%2Fcalculator%2Findex%3FlandId%3D2313%26mid1%3D0%26mid2%3D7197725821368221735%26mid3%3D7239313247488278566%26promotion_name%3D2.0%E5%B0%8F%E7%A8%8B%E5%BA%8F%E5%85%9C%E5%BA%95%E9%93%BE%E6%8E%A5-%E6%8A%96%E9%9F%B3-%E6%89%8B%E5%8A%A8-%E6%8A%95%E6%94%BE2.0-%E4%B8%8B%E8%BD%BD%E5%B0%8F%E7%A8%8B%E5%BA%8F-%E5%B7%B2%E7%94%B3%E8%AF%B7-%E5%8F%A3%E6%92%AD-%E8%BF%98%E6%AC%BE-0602-36%2526promotionid%3D7239925482619666465%26requestid%3D20230609172333B2F0DA020EDF3DF959E0%26source%3D1100726055%26union_site%3D__UNION_SITE__%26vid%3D__VID__%26projectid%3D7239924572828155962%26promotionid%3D7239925482619666465%26creativetype%3D15%26clickid%3DB.4FeqvvAlOdfyrfLLyZVFMR6DbZWTmzy2Bsp8ZeUUfLepWIpnDBgE6f2RvvOsBbNNjn777qeQtFlWUGmnuTQaFpjIlyvIhcJWP5Ozsmf2k47eC&version_type=current&scene=0&app_id=tt396d4fd0a65382ed01&version=v2&bdpsum=2c763ec&bdp_log=%7B%22launch_from%22%3A%22ad%22%7D"
	u21, _ := url.Parse(mpUrl21)
	q21 := u21.Query()
	// 2.0 decode后数据
	//mpUrl3 := "sslocal://microapp?start_page=pages%2Fcalculator%2Findex%3FlandId%3D2313%26mid1%3D__MID1__%26mid2%3D__MID2__%26mid3%3D__MID3__%26promotion_name%3D__PROMOTION_NAME__%26promotionid%3D__PROMOTION_ID__%26requestid%3D__REQUESTID__%26source%3D1100726055%26union_site%3D__UNION_SITE__%26vid%3D__VID__%26projectid%3D__PROJECT_ID__%26promotionid%3D__PROMOTION_ID__%26creativetype%3D__CTYPE__%26clickid%3D__CLICKID__&version_type=current&scene=0&app_id=tt396d4fd0a65382ed01&version=v2&bdpsum=2c763ec&bdp_log=%7B%22launch_from%22%3A%22ad%22%7D"
	//
	//u3, _ := url.Parse(mpUrl3)
	//q3 := u3.Query()
	fmt.Println(q1)
	fmt.Println(q2)
	fmt.Println(q21)

}
