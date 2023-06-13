package main

//
//import (
//	"encoding/base64"
//	"fmt"
//	"github.com/golang/protobuf/proto"
//)
//
//func main() {
//	msg := "Hello, 世界"
//
//	data, _ := base64.RawURLEncoding.DecodeString("EJiosffn7ZADGKC48IiM9YkBIKjKgICzjeYCMAw4wbgCQiIyMDIzMDQxNDE0MTkxOTQ0NDBFQUJBQ0NBRjMzOEE4MEYzSMG4ApABAA")
//
//	creativeStrDecoded, err := base64.StdEncoding.DecodeString(creativeStr)
//	bytes, err := base64.StdEncoding.DecodeString(creativeStr)
//
//	pbItem := CreativeItems{}
//	err = proto.Unmarshal(bytes, &pbItem)
//	fmt.Println("creativeStrDecoded res", string(creativeStrDecoded))
//
//	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
//	fmt.Println("encode res", encoded)
//	decoded, err := base64.StdEncoding.DecodeString(encoded)
//	if err != nil {
//		fmt.Println("decode error:", err)
//		return
//	}
//	fmt.Println("encode res", string(decoded))
//
//}
//
//type CreativeItems struct {
//	state         protoimpl.MessageState
//	sizeCache     protoimpl.SizeCache
//	unknownFields protoimpl.UnknownFields
//
//	Materials            []*MaterialItem      `protobuf:"bytes,1,rep,name=materials,proto3" json:"materials,omitempty"`                                                             // 通用字段
//	TitleId              int64                `protobuf:"varint,5,opt,name=title_id,json=titleId,proto3" json:"title_id,omitempty"`                                                 // 标题id，头条lite标题改写在用
//	Tagv3Id              []int64              `protobuf:"varint,6,rep,packed,name=tagv3_id,json=tagv3Id,proto3" json:"tagv3_id,omitempty"`                                          // 标签id，抖音激励在用
//	MarketPhraseId       int64                `protobuf:"varint,7,opt,name=market_phrase_id,json=marketPhraseId,proto3" json:"market_phrase_id,omitempty"`                          // 营销文案id，抖音结构化button在用
//	XiguaMarketPhraseId  int64                `protobuf:"varint,8,opt,name=xigua_market_phrase_id,json=xiguaMarketPhraseId,proto3" json:"xigua_market_phrase_id,omitempty"`         // 西瓜信息流，副标题,
//	ToutiaoExplainTags   []int64              `protobuf:"varint,9,rep,packed,name=toutiao_explain_tags,json=toutiaoExplainTags,proto3" json:"toutiao_explain_tags,omitempty"`       // 头条合规样式，标签物料
//	ToutiaoExplainPhrase int64                `protobuf:"varint,10,opt,name=toutiao_explain_phrase,json=toutiaoExplainPhrase,proto3" json:"toutiao_explain_phrase,omitempty"`       // 头条合规样式，短句物料
//	FanqieNovelTag       int64                `protobuf:"varint,11,opt,name=fanqie_novel_tag,json=fanqieNovelTag,proto3" json:"fanqie_novel_tag,omitempty"`                         // 番茄小说使用的标签
//	ItemId               int64                `protobuf:"varint,12,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`                                                   // 标题视频联合优选，返回的视频
//	LubanProductTagIds   []int64              `protobuf:"varint,13,rep,packed,name=luban_product_tag_ids,json=lubanProductTagIds,proto3" json:"luban_product_tag_ids,omitempty"`    // 番茄小说使用的鲁班电商多个标签
//	LubanProductNameId   int64                `protobuf:"varint,14,opt,name=luban_product_name_id,json=lubanProductNameId,proto3" json:"luban_product_name_id,omitempty"`           // 番茄小说使用的鲁班电商的商品名
//	UnionLongPhraseIds   []int64              `protobuf:"varint,15,rep,packed,name=union_long_phrase_ids,json=unionLongPhraseIds,proto3" json:"union_long_phrase_ids,omitempty"`    // 穿山甲联盟使用的营销长句
//	UnionShortPhraseIds  []int64              `protobuf:"varint,16,rep,packed,name=union_short_phrase_ids,json=unionShortPhraseIds,proto3" json:"union_short_phrase_ids,omitempty"` // 穿山甲联盟使用的营销短句
//	UnionCreativeTagIds  []int64              `protobuf:"varint,17,rep,packed,name=union_creative_tag_ids,json=unionCreativeTagIds,proto3" json:"union_creative_tag_ids,omitempty"` // 穿山甲联盟使用的创意标签
//	UnionRewriteTitle    *UnionRewriteTitle   `protobuf:"bytes,18,opt,name=union_rewrite_title,json=unionRewriteTitle,proto3" json:"union_rewrite_title,omitempty"`                 // 穿山甲联盟标题改写
//	UnionVideoInfo       *UnionVideoInfo      `protobuf:"bytes,19,opt,name=union_video_info,json=unionVideoInfo,proto3" json:"union_video_info,omitempty"`                          // 给穿山甲的视频信息
//	VariationId          int64                `protobuf:"varint,20,opt,name=variation_id,json=variationId,proto3" json:"variation_id,omitempty"`                                    // 穿山甲落地页id
//	PlayableVarIds       []int64              `protobuf:"varint,21,rep,packed,name=playable_var_ids,json=playableVarIds,proto3" json:"playable_var_ids,omitempty"`                  //playable dco组件ids
//	UnionCtaTextIds      []int64              `protobuf:"varint,22,rep,packed,name=union_cta_text_ids,json=unionCtaTextIds,proto3" json:"union_cta_text_ids,omitempty"`             // 穿山甲cta文案id
//	UnionRewriteVideo    *UnionRewriteVideo   `protobuf:"bytes,23,opt,name=union_rewrite_video,json=unionRewriteVideo,proto3" json:"union_rewrite_video,omitempty"`                 // 穿山甲改写视频信息
//	UnionRewriteLuQuery  *UnionRewriteLuQuery `protobuf:"bytes,24,opt,name=union_rewrite_lu_query,json=unionRewriteLuQuery,proto3" json:"union_rewrite_lu_query,omitempty"`         // 穿山甲 搜索LU广告的query优选
//}
