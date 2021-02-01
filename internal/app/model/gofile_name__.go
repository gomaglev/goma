package model

import (
	"context"

	"app_module__/internal/app/dto"

	"app_module__/pkg/proto/common"
	gopackage_name__ "app_module__/pkg/proto/gopackage_name__"
)

// IMessageTypeName__ interface for MessageTypeName__ model
type IMessageTypeName__ interface {
	// List   message
	List(ctx context.Context, params *dto.ListMessageTypeNamePlural__Param, opts ...*common.QueryOptions) (*gopackage_name__.MessageTypeNamePlural__, error)
	// Get    message
	Get(ctx context.Context, params *dto.GetMessageTypeName__Param) (*gopackage_name__.MessageTypeName__, error)
	// Create message
	Create(ctx context.Context, message *gopackage_name__.MessageTypeName__) (*common.IDResult, error)
	// Update message
	Update(ctx context.Context, params *dto.UpdateMessageTypeName__Param, message *gopackage_name__.MessageTypeName__) (*int64, error)
	// Update message by columns
	UpdateColumns(ctx context.Context, params *dto.UpdateMessageTypeName__Param, values map[string]interface{}) (*int64, error)
	// Delete message
	Delete(ctx context.Context, params *dto.DeleteMessageTypeName__Param) (*int64, error)
}
