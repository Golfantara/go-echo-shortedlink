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

func (svc *service) FindAllGoly(page, size int) []dtos.GolyResponse {
	var goly []dtos.GolyResponse

	golyEnt := svc.model.Paginate(page, size)

	for _, golys := range golyEnt {
		var data dtos.GolyResponse

		if err := smapping.FillStruct(&data, smapping.MapFields(golys)); err != nil {
			log.Error(err.Error())
		}

		goly = append(goly, data)
	}
	return goly
}

func (svc *service) FindGolyByID(golyID int) *dtos.GolyResponse {
	res := dtos.GolyResponse{}
	goly := svc.model.SelectByID(golyID)

	if goly == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(goly))
	if err != nil {
		log.Error(err)
		return nil
	}
	return &res
}

func(svc *service) GetGolyByUrl(url string) (goly.Goly, error) {
	return svc.model.FindByGolyUrl(url)
}

func (svc *service) IncreaseClickAndRedirect(goly goly.Goly) error {
    goly.Clicked += 1
    return svc.model.UpdateButton(goly)
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


func (svc *service) Modify(golyData dtos.CreateGolyInput, golyID int) bool {
	newGoly := goly.Goly{}

	err := smapping.FillStruct(&newGoly, smapping.MapFields(golyData))
	if err != nil {
		log.Error(err)
		return false
	}

	newGoly.ID = uint64(golyID)
	rowsAffected := svc.model.Update(newGoly)

	if rowsAffected <= 0 {
		log.Error("there is no goly updated!")
		return false
	}
	return true
}

func (svc *service) Remove(golyID int) bool {
	rowsAffected := svc.model.DeleteByID(golyID)

	if rowsAffected <= 0 {
        log.Error("there is no user removed!")
        return false
    }
	return true
}