package usecase

import (
	"shortlink/features/goly"
	"shortlink/features/goly/dtos"
	"shortlink/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model goly.Repository
}

func New(model goly.Repository) goly.UseCase {
	return &service {
        model: model,
    }
}

func (svc *service) Create(newGoly dtos.CreateGolyInput) *dtos.GolyResponse {
	goly := goly.Goly{}
	err := smapping.FillStruct(&goly, smapping.MapFields(newGoly))

	if err != nil {
		log.Error(err)
		return nil
	}
	if goly.Random {
		goly.Goly = helpers.RandomURL(8)
	}
	golyID := svc.model.Insert(&goly)
	if golyID == nil {
		return nil
	}
	resGoly := dtos.GolyResponse{}
	errRes := smapping.FillStruct(&resGoly, smapping.MapFields(goly))
	if errRes != nil {
		log.Error(errRes)
		return nil
	}
	return &resGoly
}