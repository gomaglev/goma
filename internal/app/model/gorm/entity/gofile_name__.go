package entity

import (
	"context"

	"app_module__/pkg/util"

	proto "app_module__/pkg/proto/gopackage_name__"

	"gorm.io/gorm"
)

// MessageTypeName__
type MessageTypeName__ struct {
	Id string
	//EntityProtoParentIDs__
}

// GetMessageTypeName__DB
func GetMessageTypeName__DB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, (*gorm.DB)(defDB), new(MessageTypeName__))
}

// MessageTypeName__Proto
type MessageTypeName__Proto proto.MessageTypeName__

// TableName
func (e MessageTypeName__) TableName() string {
	return "table_name__"
}

// ToProto
func (e *MessageTypeName__) ToProto() *proto.MessageTypeName__ {
	item := new(proto.MessageTypeName__)
	_ = util.StructMapToStruct(e, item)

	return item
}

// MessageTypeNamePlural__
type MessageTypeNamePlural__ []*MessageTypeName__

// ToProtos
func (e MessageTypeNamePlural__) ToProtos() []*proto.MessageTypeName__ {
	list := make([]*proto.MessageTypeName__, len(e))
	for i, v := range e {
		list[i] = v.ToProto()
	}
	return list
}
