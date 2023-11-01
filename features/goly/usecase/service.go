package usecase

import (
	"bytes"
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"shortlink/features/goly"
	"shortlink/features/goly/dtos"
	"shortlink/helpers"
	"time"

	"github.com/jung-kurt/gofpdf"
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

func (svc *service) ExportIPToPDfAndSave() (string, error) {
	// Create a new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Define font and size
	pdf.SetFont("Arial", "B", 16)

	// Add a title
	pdf.Cell(40, 10, "IP Addresses")

	// Fetch the IP address data from your database
	ipAddresses := svc.model.PaginateIP(1, 10) // Adjust page and size as needed

	// Loop through the IP addresses and add them to the PDF
	for _, ip := range ipAddresses {
		pdf.Ln(10)
		pdf.Cell(0, 10, ip.Address)
	}

	// Generate the PDF content
	pdfBuffer := new(bytes.Buffer)
	err := pdf.Output(pdfBuffer)
	if err != nil {
		return "", err
	}

	// Save the PDF in the user's "Downloads" directory
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	downloadsPath := filepath.Join(usr.HomeDir, "Downloads")
	filePath := filepath.Join(downloadsPath, "ip_addresses.pdf")

	pdfFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer pdfFile.Close()

	_, err = pdfBuffer.WriteTo(pdfFile)
	if err != nil {
		return "", err
	}

	return filePath, nil
}



func (svc *service) FindAllIP(page, size int) []goly.IPAdresses{
	var ip []goly.IPAdresses

	ipEnt := svc.model.PaginateIP(page, size)

	for _, ips := range ipEnt {
		var data goly.IPAdresses

		if err := smapping.FillStruct(&data, smapping.MapFields(ips)); err != nil {
			log.Error(err.Error())
		}
		ip = append(ip, data)
	}
	return ip
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
	goly, err := svc.model.FindByGolyUrl(url)
	if err != nil {
		return goly, err
	}
	
	if time.Now().After(goly.ExpiryDate) {
		return goly, errors.New("Link is expired")
	}
	return svc.model.FindByGolyUrl(url)
}

func(svc *service) SearchGoly(short string) ([]goly.Goly, error) {
    return svc.model.SearchingGoly(short)
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
	goly.ExpiryDate = time.Now().AddDate(0,0,newGoly.ExpiryInDays)
	if goly.Random {
		goly.Custom = helpers.RandomURL(8)
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

func (svc *service) StoreIPAddress(goly goly.Goly, ip string) error {
	return svc.model.StoreIPForGoly(goly.ID, ip)
}