package dto

import (
	"GOMA/internal/app/model/gorm/entity"
	"GOMA/pkg/util"

	"GOMA/pkg/proto/common"
	proto "GOMA/pkg/proto/order/item/message"
)

// GetOrderItemMessageParam
type GetOrderItemMessageParam struct {
	Id string
}

// ListOrderItemMessagesParam for bll inner usage
type ListOrderItemMessagesParam struct {
	Pagination *common.PaginationParam
	Id         string
	OrderId    string
	ItemId     string
	Ids        []string
}

// ListOrderItemMessagesOptions
type ListOrderItemMessagesOptions struct {
	OrderByFields []*common.OrderByField
	SelectFields  []string
}

// UpdateOrderItemMessageParam
type UpdateOrderItemMessageParam struct {
	Id string
}

// DeleteOrderItemMessageParam
type DeleteOrderItemMessageParam struct {
	Id  string
	Ids string
}

// ProtoOrderItemMessages
type ProtoOrderItemMessages proto.OrderItemMessages

// ToIDs
func (a *ProtoOrderItemMessages) ToIDs() []string {
	list := make([]string, len(a.List))
	for i, item := range a.List {
		list[i] = item.Id
	}
	return list
}

// ToMap
func (a *ProtoOrderItemMessages) ToMap() map[string]*proto.OrderItemMessage {
	protoMap := make(map[string]*proto.OrderItemMessage)

	for _, item := range a.List {
		protoMap[item.Id] = item
	}

	return protoMap
}

// ProtoOrderItemMessage proto type
type ProtoOrderItemMessage struct {
	OrderItemMessage *proto.OrderItemMessage
}

// ToEntity converts to entity
func (p *ProtoOrderItemMessage) ToEntity() *entity.OrderItemMessage {
	item := new(entity.OrderItemMessage)
	_ = util.StructMapToStruct(p, item)
	// conversion
	return item
}
