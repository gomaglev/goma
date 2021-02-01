package dto

import (
	"app_module__/internal/app/model/gorm/entity"
	"app_module__/pkg/util"

	"app_module__/pkg/proto/common"
	proto "app_module__/pkg/proto/gopackage_name__"
)

// GetMessageTypeName__Param
type GetMessageTypeName__Param struct {
	Id string
}

// ListMessageTypeNamePlural__Param for bll inner usage
type ListMessageTypeNamePlural__Param struct {
	Pagination *common.PaginationParam
	Id         string
	Ids        []string
	//DtoProtoParentIDs__
}

// ListMessageTypeNamePlural__Options
type ListMessageTypeNamePlural__Options struct {
	OrderByFields []*common.OrderByField
	SelectFields  []string
}

// UpdateMessageTypeName__Param
type UpdateMessageTypeName__Param struct {
	Id string
}

// DeleteMessageTypeName__Param
type DeleteMessageTypeName__Param struct {
	Id  string
	Ids []string
}

// ProtoMessageTypeNamePlural__
type ProtoMessageTypeNamePlural__ proto.MessageTypeNamePlural__

// ToIDs
func (a *ProtoMessageTypeNamePlural__) ToIDs() []string {
	list := make([]string, len(a.List))
	for i, item := range a.List {
		list[i] = item.Id
	}
	return list
}

// ToMap
func (a *ProtoMessageTypeNamePlural__) ToMap() map[string]*proto.MessageTypeName__ {
	protoMap := make(map[string]*proto.MessageTypeName__)

	for _, item := range a.List {
		protoMap[item.Id] = item
	}

	return protoMap
}

// ProtoMessageTypeName__ proto type
type ProtoMessageTypeName__ struct {
	MessageTypeName__ *proto.MessageTypeName__
}

// ToEntity converts to entity
func (p *ProtoMessageTypeName__) ToEntity() *entity.MessageTypeName__ {
	item := new(entity.MessageTypeName__)
	_ = util.StructMapToStruct(p, item)
	// conversion
	return item
}
