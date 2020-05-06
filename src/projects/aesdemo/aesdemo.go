package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gopractice/common"
)

func main() {
	order := &CreateReqData{
		ChannelId:         3,
		SenderCell:        "1234567891",
		SenderFirstName:   "b2c sender",
		OriDetailedAddr:   "hotel,lagos,ng",
		OriCityDistrict:   "Lagos,Lagos_test",
		ReceiverCell:      "9876543210",
		ReceiverFirstName: "b2c receiver 【",
		DestDetailedAddr:  "hotel,lagos,ng",
		DestCityDistrict:  "Lagos,Lagos_test",
		DestLat:           7.4001399,
		DestLng:           3.7677819,
		ProductCategory:   1,
		Price:             15,
		Weight:            10,
		Volume:            "2*2*2",
		PaymentMethod:     1,
	}

	encode, e := common.DataAesEncrypt(order, "aCqonzvpalooWtVT", "oexpressgogo!@#$")

	if e != nil {
		fmt.Printf("error: %s \n", e)
		return
	}

	fmt.Printf("encode string: %s \n", encode)

	decode, e := common.DataAesDecrypt(encode, "aCqonzvpalooWtVT", "oexpressgogo!@#$")

	if e != nil {
		fmt.Printf("error: %s \n", e)
		return
	}

	fmt.Printf("decode string: %s \n", decode)

	decode, e = common.DataAesDecrypt("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", "aCqonzvpalooWtVT", "oexpressgogo!@#$")

	if e != nil {
		fmt.Printf("error: %s \n", e)
		return
	}

	fmt.Printf("decode string: %s \n", decode)

	decodeData, e := base64.StdEncoding.DecodeString("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")

	if e != nil {
		return
	}

	decResult, e := common.AesDecrypt(decodeData, []byte("aCqonzvpalooWtVT"), []byte("oexpressgogo!@#$"))

	if e != nil {
		return
	}

	rsp := &DetailRsp{}

	e = json.Unmarshal(decResult, rsp)

	if e != nil{
		return
	}

	enStr, _ := common.DataAesEncrypt(rsp, "aCqonzvpalooWtVT", "oexpressgogo!@#$")

	fmt.Printf("encreypt\n%s\n\n", enStr)

	common.DataAesDecrypt(enStr, "aCqonzvpalooWtVT", "oexpressgogo!@#$")

}

// HTTP POST request: /third_part/order/create
type CreateReqData struct {
	ChannelId int64 `protobuf:"varint,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id"`
	// 寄件用户电话
	SenderCell string `protobuf:"bytes,2,opt,name=sender_cell,json=senderCell,proto3" json:"sender_cell"`
	// 寄件人名
	SenderFirstName string `protobuf:"bytes,3,opt,name=sender_first_name,json=senderFirstName,proto3" json:"sender_first_name"`
	// 下单人地址
	OriDetailedAddr string `protobuf:"bytes,4,opt,name=ori_detailed_addr,json=oriDetailedAddr,proto3" json:"ori_detailed_addr"`
	// 寄件人二级区域
	OriCityDistrict string `protobuf:"bytes,5,opt,name=ori_city_district,json=oriCityDistrict,proto3" json:"ori_city_district"`
	// 收件人电话
	ReceiverCell string `protobuf:"bytes,6,opt,name=receiver_cell,json=receiverCell,proto3" json:"receiver_cell"`
	// 收件人名
	ReceiverFirstName string `protobuf:"bytes,7,opt,name=receiver_first_name,json=receiverFirstName,proto3" json:"receiver_first_name"`
	// 收件人地址
	DestDetailedAddr string `protobuf:"bytes,8,opt,name=dest_detailed_addr,json=destDetailedAddr,proto3" json:"dest_detailed_addr"`
	// 收件人二级区域
	DestCityDistrict string `protobuf:"bytes,9,opt,name=dest_city_district,json=destCityDistrict,proto3" json:"dest_city_district"`
	// 收件人维度
	DestLat float64 `protobuf:"fixed64,10,opt,name=dest_lat,json=destLat,proto3" json:"dest_lat"`
	// 收件人经度
	DestLng float64 `protobuf:"fixed64,11,opt,name=dest_lng,json=destLng,proto3" json:"dest_lng"`
	// 货品种类
	ProductCategory int32 `protobuf:"varint,12,opt,name=product_category,json=productCategory,proto3" json:"product_category"`
	// 货品价值
	Price float64 `protobuf:"fixed64,13,opt,name=price,proto3" json:"price"`
	// 备注
	Comment string `protobuf:"bytes,14,opt,name=comment,proto3" json:"comment"`
	// 货品重量
	Weight float64 `protobuf:"fixed64,15,opt,name=weight,proto3" json:"weight"`
	// 货品体积 Deprecated
	Volume string `protobuf:"bytes,16,opt,name=volume,proto3" json:"volume"`
	// 支付方式（1-现金支付 2-线上支付 3.货到付款 4.周结 5.月结）
	PaymentMethod int32 `protobuf:"varint,17,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method"`
}

type Order struct {
	// 唯一订单号 取件运单号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 创建订单用户id
	CreateUserId int64 `protobuf:"varint,2,opt,name=create_user_id,json=createUserId,proto3" json:"create_user_id"`
	// 订单来源 0:货主下单\n1:揽件员下单 2:后台 3.b2b 4.批量创建 5.b2c
	OrderSource int32 `protobuf:"varint,3,opt,name=order_source,json=orderSource,proto3" json:"order_source"`
	// 寄件用户电话
	SenderCell string `protobuf:"bytes,4,opt,name=sender_cell,json=senderCell,proto3" json:"sender_cell"`
	// 寄件人名
	SenderFirstName string `protobuf:"bytes,5,opt,name=sender_first_name,json=senderFirstName,proto3" json:"sender_first_name"`
	// 寄件人姓
	SenderLastName string `protobuf:"bytes,6,opt,name=sender_last_name,json=senderLastName,proto3" json:"sender_last_name"`
	// 0:需要揽件 1:自送
	WithoutCollect int32 `protobuf:"varint,7,opt,name=without_collect,json=withoutCollect,proto3" json:"without_collect"`
	// 自送or揽件时划分的hub
	OriHubId int64 `protobuf:"varint,8,opt,name=ori_hub_id,json=oriHubId,proto3" json:"ori_hub_id"`
	// 寄出地址纬度
	OriLat float64 `protobuf:"fixed64,9,opt,name=ori_lat,json=oriLat,proto3" json:"ori_lat"`
	// 寄出地址经度
	OriLng float64 `protobuf:"fixed64,10,opt,name=ori_lng,json=oriLng,proto3" json:"ori_lng"`
	// 寄出地址
	OriAddr string `protobuf:"bytes,11,opt,name=ori_addr,json=oriAddr,proto3" json:"ori_addr"`
	// 寄出地址细节
	OriDetailedAddr string `protobuf:"bytes,12,opt,name=ori_detailed_addr,json=oriDetailedAddr,proto3" json:"ori_detailed_addr"`
	// 收件人电话
	ReceiverCell string `protobuf:"bytes,13,opt,name=receiver_cell,json=receiverCell,proto3" json:"receiver_cell"`
	// 收件人名
	ReceiverFirstName string `protobuf:"bytes,14,opt,name=receiver_first_name,json=receiverFirstName,proto3" json:"receiver_first_name"`
	// 收件人姓
	ReceiverLastName string `protobuf:"bytes,15,opt,name=receiver_last_name,json=receiverLastName,proto3" json:"receiver_last_name"`
	// 收件人维度
	DestLat float64 `protobuf:"fixed64,16,opt,name=dest_lat,json=destLat,proto3" json:"dest_lat"`
	// 收件人经度
	DestLng float64 `protobuf:"fixed64,17,opt,name=dest_lng,json=destLng,proto3" json:"dest_lng"`
	// 收件地址
	DestAddr string `protobuf:"bytes,18,opt,name=dest_addr,json=destAddr,proto3" json:"dest_addr"`
	// 收件地址详情
	DestDetailedAddr string `protobuf:"bytes,19,opt,name=dest_detailed_addr,json=destDetailedAddr,proto3" json:"dest_detailed_addr"`
	// 当前运输单id
	CurrentTransportId int64 `protobuf:"varint,20,opt,name=current_transport_id,json=currentTransportId,proto3" json:"current_transport_id"`
	// 当前挂起对应的异常记录id,  0: 未挂起，状态正常
	CurrentHoldRecordId int64 `protobuf:"varint,21,opt,name=current_hold_record_id,json=currentHoldRecordId,proto3" json:"current_hold_record_id"`
	// 订单状态 0:未定义,\n1:已创,\n2:未支,\n3:已揽,\n4:仓储,\n5:转运,\n6:派送,\n7:已签,\n8:已取,\n9:异常关闭
	Status int32 `protobuf:"varint,22,opt,name=status,proto3" json:"status"`
	// 揽收时间
	CollectedTime int64 `protobuf:"varint,23,opt,name=collected_time,json=collectedTime,proto3" json:"collected_time"`
	// 签收时间
	FinishTime int64 `protobuf:"varint,24,opt,name=finish_time,json=finishTime,proto3" json:"finish_time"`
	// 取消时间
	CancelTime int64 `protobuf:"varint,25,opt,name=cancel_time,json=cancelTime,proto3" json:"cancel_time"`
	// 0: 货主取消\n1: 揽件员取消\n2: 系统取消
	CancelRole int32 `protobuf:"varint,26,opt,name=cancel_role,json=cancelRole,proto3" json:"cancel_role"`
	// 货品种类
	ProductCategory int32 `protobuf:"varint,27,opt,name=product_category,json=productCategory,proto3" json:"product_category"`
	// 货品价值
	Price float64 `protobuf:"fixed64,28,opt,name=price,proto3" json:"price"`
	// 备注
	Comment string `protobuf:"bytes,29,opt,name=comment,proto3" json:"comment"`
	// 收件码
	DeliveryCode int32 `protobuf:"varint,30,opt,name=delivery_code,json=deliveryCode,proto3" json:"delivery_code"`
	// 订单创建时间戳
	CreateTime int64 `protobuf:"varint,31,opt,name=create_time,json=createTime,proto3" json:"create_time"`
	// 更新时间
	UpdateTime int64 `protobuf:"varint,32,opt,name=update_time,json=updateTime,proto3" json:"update_time"`
	// 货品重量
	Weight float64 `protobuf:"fixed64,33,opt,name=weight,proto3" json:"weight"`
	// 货品体积
	Volume string `protobuf:"bytes,34,opt,name=volume,proto3" json:"volume"`
	// 收件人划分的hub
	DestHubId int64 `protobuf:"varint,35,opt,name=dest_hub_id,json=destHubId,proto3" json:"dest_hub_id"`
	// 处理时间
	CloseTime int64 `protobuf:"varint,36,opt,name=close_time,json=closeTime,proto3" json:"close_time"`
	// 运输费用
	DeliverFee float64 `protobuf:"fixed64,37,opt,name=deliver_fee,json=deliverFee,proto3" json:"deliver_fee"`
	// 货品种类名称
	ProductCategoryName string `protobuf:"bytes,38,opt,name=product_category_name,json=productCategoryName,proto3" json:"product_category_name"`
	// basic fee
	BasicFee float64 `protobuf:"fixed64,39,opt,name=basic_fee,json=basicFee,proto3" json:"basic_fee"`
	// weight fee
	WeightFee float64 `protobuf:"fixed64,40,opt,name=weight_fee,json=weightFee,proto3" json:"weight_fee"`
	// insurance fee
	InsuranceFee float64 `protobuf:"fixed64,41,opt,name=insurance_fee,json=insuranceFee,proto3" json:"insurance_fee"`
	// pickup fee
	PickupFee float64 `protobuf:"fixed64,42,opt,name=pickup_fee,json=pickupFee,proto3" json:"pickup_fee"`
	// tax fee
	TaxFee float64 `protobuf:"fixed64,43,opt,name=tax_fee,json=taxFee,proto3" json:"tax_fee"`
	// 支付方式: 1-现金支付 2-线上支付  3-到付
	PaymentMethod int32 `protobuf:"varint,44,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method"`
	// 取消备注
	CancelComment string `protobuf:"bytes,45,opt,name=cancel_comment,json=cancelComment,proto3" json:"cancel_comment"`
	// 揽件图片地址
	PickupPicUrlList string `protobuf:"bytes,46,opt,name=pickup_pic_url_list,json=pickupPicUrlList,proto3" json:"pickup_pic_url_list"`
	// 揽件员确认订单时间
	ConfirmTime int64 `protobuf:"varint,47,opt,name=confirm_time,json=confirmTime,proto3" json:"confirm_time"`
	// 城市id
	CityId int64 `protobuf:"varint,48,opt,name=city_id,json=cityId,proto3" json:"city_id"`
	// 送达图片地址
	DeliveredPicUrlList string `protobuf:"bytes,49,opt,name=delivered_pic_url_list,json=deliveredPicUrlList,proto3" json:"delivered_pic_url_list"`
	// convert id to string
	IdStr string `protobuf:"bytes,50,opt,name=id_str,json=idStr,proto3" json:"id_str"`
	// 订单后四位，用于模糊查询
	ItemCode string `protobuf:"bytes,51,opt,name=item_code,json=itemCode,proto3" json:"item_code"`
	// 是否回款
	CashReceived int32 `protobuf:"varint,52,opt,name=cash_received,json=cashReceived,proto3" json:"cash_received"`
	// 是否使用万能码
	UseUniversalCode int32 `protobuf:"varint,53,opt,name=use_universal_code,json=useUniversalCode,proto3" json:"use_universal_code"`
	// 回款时间戳
	CashReceivedTime int64 `protobuf:"varint,54,opt,name=cash_received_time,json=cashReceivedTime,proto3" json:"cash_received_time"`
	// 是否打印面单
	PrintStatus int32 `protobuf:"varint,55,opt,name=print_status,json=printStatus,proto3" json:"print_status"`
	// 特殊价格模式： 0 normal默认值, 1 special
	SpecialPriceMode int32 `protobuf:"varint,56,opt,name=special_price_mode,json=specialPriceMode,proto3" json:"special_price_mode"`
	// 寄件人城市Id
	SenderCityId int64 `protobuf:"varint,57,opt,name=sender_city_id,json=senderCityId,proto3" json:"sender_city_id"`
	// 寄件人区Id
	SenderDistrictId int64 `protobuf:"varint,58,opt,name=sender_district_id,json=senderDistrictId,proto3" json:"sender_district_id"`
	// 收件人城市Id
	ReceiverCityId int64 `protobuf:"varint,59,opt,name=receiver_city_id,json=receiverCityId,proto3" json:"receiver_city_id"`
	// 收件人区Id
	ReceiverDistrictId int64 `protobuf:"varint,60,opt,name=receiver_district_id,json=receiverDistrictId,proto3" json:"receiver_district_id"`
	// 揽件图片地址数组
	PickupPicUrlListArr []string `protobuf:"bytes,100000,rep,name=pickup_pic_url_list_arr,json=pickupPicUrlListArr,proto3" json:"pickup_pic_url_list_arr"`
	// 送达图片地址数组
	DeliveredPicUrlListArr []string `protobuf:"bytes,100001,rep,name=delivered_pic_url_list_arr,json=deliveredPicUrlListArr,proto3" json:"delivered_pic_url_list_arr"`
	// 收费明细
	CostDetail []*CostDetail `protobuf:"bytes,100002,rep,name=cost_detail,json=costDetail,proto3" json:"cost_detail"`
}

type CostDetail struct {
	// 费用名称
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key"`
	// 费用金额
	Val string `protobuf:"bytes,2,opt,name=val,proto3" json:"val"`
}

type DetailRsp struct {
	// 订单信息
	Order          *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order"`
	AppOrderStatus int32  `protobuf:"varint,2,opt,name=app_order_status,json=appOrderStatus,proto3" json:"app_order_status"`
	// transport detail item
	DetailItem []*TransportDetailItem `protobuf:"bytes,3,rep,name=detail_item,json=detailItem,proto3" json:"detail_item"`
}

type TransportDetailItem struct {
	// status
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status"`
	// timestamp
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp"`
	// information
	Information string `protobuf:"bytes,3,opt,name=information,proto3" json:"information"`
	// 揽件员\派件员 姓名
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	// 揽件员\派件员 电话
	Phone string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone"`
}
