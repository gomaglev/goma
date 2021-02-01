package model

import (
	"context"

	"app_module__/internal/app/dto"
	"app_module__/internal/app/model"
	"app_module__/internal/app/model/gorm/entity"

	"app_module__/pkg/proto/common"

	"gorm.io/gorm"

	gopackage_name__ "app_module__/pkg/proto/gopackage_name__"

	"github.com/google/wire"
	"github.com/pkg/errors"
)

var _ model.IMessageTypeName__ = (*MessageTypeName__)(nil)

// MessageTypeName__Set
var MessageTypeName__Set = wire.NewSet(wire.Struct(new(MessageTypeName__), "*"), wire.Bind(new(model.IMessageTypeName__), new(*MessageTypeName__)))

// MessageTypeName__
type MessageTypeName__ struct {
	DB *gorm.DB
}

// List
func (a *MessageTypeName__) List(ctx context.Context, params *dto.ListMessageTypeNamePlural__Param, opts ...*common.QueryOptions) (*gopackage_name__.MessageTypeNamePlural__, error) {
	opt := GetQueryOption(opts...)
	db := entity.GetMessageTypeName__DB(ctx, a.DB)

	if v := params.Id; v != "" {
		db = db.Where("id = ?", v)
	}

	if v := params.Ids; len(v) > 0 {
		db = db.Where("id in ?", v)
	}

	var list entity.MessageTypeNamePlural__
	pr, err := WrapPageQuery(ctx, db, params.Pagination, opt, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &gopackage_name__.MessageTypeNamePlural__{
		Pagination: pr,
		List:       list.ToProtos(),
	}
	return qr, nil
}

// Get
func (a *MessageTypeName__) Get(ctx context.Context, params *dto.GetMessageTypeName__Param) (*gopackage_name__.MessageTypeName__, error) {
	var message entity.MessageTypeName__
	ok, err := FindOne(ctx, entity.GetMessageTypeName__DB(ctx, a.DB).Where(`id=?`, params.Id), &message)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return message.ToProto(), nil
}

// Create
func (a *MessageTypeName__) Create(ctx context.Context, item *gopackage_name__.MessageTypeName__) (*common.IDResult, error) {
	pitem := dto.ProtoMessageTypeName__{MessageTypeName__: item}
	eitem := pitem.ToEntity()
	result := entity.GetMessageTypeName__DB(ctx, a.DB).Create(eitem)
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &common.IDResult{
		Id: item.Id,
	}, nil
}

// Update
func (a *MessageTypeName__) Update(ctx context.Context, params *dto.UpdateMessageTypeName__Param, item *gopackage_name__.MessageTypeName__) (*int64, error) {
	pitem := dto.ProtoMessageTypeName__{MessageTypeName__: item}
	eitem := pitem.ToEntity()
	result := entity.GetMessageTypeName__DB(ctx, a.DB).Where("id=?", params.Id).Updates(eitem)
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &result.RowsAffected, nil
}

// UpdateColumns updates order columns
func (a *MessageTypeName__) UpdateColumns(ctx context.Context, params *dto.UpdateMessageTypeName__Param, values map[string]interface{}) (*int64, error) {
	db := entity.GetMessageTypeName__DB(ctx, a.DB)

	db = db.Where("id = ?", params.Id)

	result := db.UpdateColumns(values)
	if err := result.Error; err != nil {
		return &result.RowsAffected, errors.WithStack(err)
	}
	return &result.RowsAffected, nil
}

// Delete
func (a *MessageTypeName__) Delete(ctx context.Context, params *dto.DeleteMessageTypeName__Param) (*int64, error) {
	result := entity.GetMessageTypeName__DB(ctx, a.DB).Where("id=?", params.Id).Delete(entity.MessageTypeName__{})
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &result.RowsAffected, nil
}
