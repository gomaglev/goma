package gopackage_name__

import (
	"context"

	"app_module__/internal/app/dto"
	"app_module__/internal/app/model"
	iutil "app_module__/internal/pkg/util"

	"github.com/google/wire"
)

// make sure PbName__ implements PbName__ServiceServer
var _ PbName__ServiceServer = (*PbName__Service)(nil)

// PbName__Set injection
var PbName__Set = wire.NewSet(wire.Struct(new(PbName__Service), "*"), wire.Bind(new(PbName__ServiceServer), new(*PbName__Service)))

// PbName__ implements PbName__ServiceServer
type PbName__Service struct {
	PbName__Model model.IMessageTypeName__
}

// Get pb_name__
func (a *PbName__Service) Get(ctx context.Context, req *GetPbName__Request) (*GetPbName__Response, error) {
	params := &dto.GetMessageTypeName__Param{
		Id: req.Id,
	}
	pb_name__, err := a.PbName__Model.Get(ctx, params)
	res := &GetPbName__Response{
		PbName__: pb_name__,
	}
	return res, err
}

// List pb_name_plural__
func (a *PbName__Service) List(ctx context.Context, req *ListPbNamePlural__Request) (*ListPbNamePlural__Response, error) {
	params := &dto.ListMessageTypeNamePlural__Param{
		Ids: req.Ids,
	}
	pb_name_plural__, err := a.PbName__Model.List(ctx, params)
	res := &ListPbNamePlural__Response{
		PbNamePlural__: pb_name_plural__,
	}
	return res, err
}

// Create pb_name__
func (a *PbName__Service) Create(ctx context.Context, req *CreatPbName__Request) (*CreatPbName__Response, error) {
	req.PbName__.Id = iutil.NewID()
	result, err := a.PbName__Model.Create(ctx, req.PbName__)
	res := &CreatPbName__Response{
		Id: result.Id,
	}
	return res, err
}

// Update pb_name__
func (a *PbName__Service) Update(ctx context.Context, req *UpdatePbName__Request) (*UpdatePbName__Response, error) {
	params := &dto.UpdateMessageTypeName__Param{
		Id: req.Id,
	}
	rows, err := a.PbName__Model.Update(ctx, params, req.PbName__)
	res := &UpdatePbName__Response{
		Updated: *rows,
	}
	return res, err
}

// Delete pb_name__
func (a *PbName__Service) Delete(ctx context.Context, req *DeletePbName__Request) (*DeletePbName__Response, error) {
	params := &dto.DeleteMessageTypeName__Param{
		Id:  req.Id,
		Ids: req.Ids,
	}
	rows, err := a.PbName__Model.Delete(ctx, params)
	res := &DeletePbName__Response{
		Deleted: *rows,
	}
	return res, err
}
