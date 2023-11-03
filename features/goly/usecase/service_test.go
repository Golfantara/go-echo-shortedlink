package usecase

import (
	"errors"
	"shortlink/features/goly"
	"shortlink/features/goly/dtos"
	"shortlink/features/goly/mocks"
	mockHelpers "shortlink/helpers/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindAll(t *testing.T) {
	randomURL := mockHelpers.NewRandomURLInterface(t)
	repo := mocks.NewRepository(t)
	service := New(repo, randomURL)

	golyEntities := []goly.Goly{
		{ID: 1, UserID: "1", Redirect: "someURL", Custom: "dida"},
		{ID: 2, UserID: "2", Redirect: "someURL", Custom: "dejan"},
	}
	repo.On("Paginate", 1, 10).Return(golyEntities).Once()
	goly := service.FindAllGoly(1, 10)

	expected := []dtos.GolyResponse{
		{ID: 1, UserID: "1", Redirect: "someURL", Custom: "dida"},
		{ID: 2, UserID: "2", Redirect: "someURL", Custom: "dejan"},
	}
	assert.Equal(t, expected, goly)
	repo.AssertExpectations(t)
}


func TestFindAllIP(t *testing.T) {
	randomURL := mockHelpers.NewRandomURLInterface(t)
	repo := mocks.NewRepository(t)
	service := New(repo, randomURL)

	ipEntities := []goly.IPAdresses{
		{ID: 1, GolyID: 1, Address: "someIPAddress"},
		{ID: 2, GolyID: 2, Address: "someIPAddress"},
	}
	repo.On("PaginateIP", 1, 10).Return(ipEntities).Once()
	ip := service.FindAllIP(1, 10)

	expected := []goly.IPAdresses{
		{ID: 1,GolyID: 1, Address: "someIPAddress"},
		{ID: 2,GolyID: 2, Address: "someIPAddress"},
	}
	assert.Equal(t, expected, ip)
	repo.AssertExpectations(t)
}

func TestFindGolyByID(t *testing.T){
	randomURL := mockHelpers.NewRandomURLInterface(t)
	repo := mocks.NewRepository(t)
	service := New(repo, randomURL)
	
	golyID := 1

	golyEntities := &goly.Goly{
		ID: 1,
		UserID: "1",
		Redirect: "someURL",
		Custom: "dida",
		Clicked: 2,
	}
	t.Run("success find by id", func(t *testing.T) {
		repo.On("SelectByID", golyID).Return(golyEntities).Once()

		result := service.FindGolyByID(golyID)
		
		expected := &dtos.GolyResponse{
			ID: golyEntities.ID,
			UserID: golyEntities.UserID,
			Redirect: golyEntities.Redirect,
			Custom: golyEntities.Custom,
			Clicked: golyEntities.Clicked,
		}

		assert.Equal(t, expected, result)

		repo.AssertExpectations(t)
	})
}

func TestRemove(t *testing.T){
	randomURL := mockHelpers.NewRandomURLInterface(t)
	repo := mocks.NewRepository(t)
	service := New(repo, randomURL)

	golyID := 1
	repo.On("DeleteByID", golyID).Return(int64(1)).Once()
	result := service.Remove(golyID)

	assert.True(t, result)
	repo.AssertExpectations(t)
}

func TestModify(t *testing.T){
	randomURL := mockHelpers.NewRandomURLInterface(t)
	repo := mocks.NewRepository(t)
	service := New(repo, randomURL)

	var updateData = dtos.CreateGolyInput{
		UserID: "1",
		Custom: "dida",
		Redirect: "someURL",
		Random: false,
		ExpiryInDays: 1,
	}
	var goly = goly.Goly{
		ID: 1,
		UserID: "1",
		Redirect: "someURL",
		Custom: "dejan",
		Clicked: 1,
		Random: false,
	}
	t.Run("succes update data", func(t *testing.T) {
		repo.On("Update", mock.Anything).Return(int64(1)).Once()

		result := service.Modify(updateData, int(goly.ID))
		assert.True(t, result)
		repo.AssertExpectations(t)
	})
}

func TestStoreIPAddress(t *testing.T){
	randomURL := mockHelpers.NewRandomURLInterface(t)
	repo := mocks.NewRepository(t)
	service := New(repo, randomURL)

	golyID := uint64(1)
	ipAddress := "192.168.1.1"

	repo.On("StoreIPForGoly", golyID, ipAddress).Return(nil)

	err := service.StoreIPAddress(goly.Goly{ID: golyID}, ipAddress)

	repo.AssertExpectations(t)

	assert.NoError(t, err)
}

func TestSearchGoly(t *testing.T){
	randomURL := mockHelpers.NewRandomURLInterface(t)
	repo := mocks.NewRepository(t)
	service := New(repo, randomURL)

	searchTerm := "example"

	mockResponse := []goly.Goly{
		{ID: 1, UserID: "1", Redirect: "someURL", Custom: "example1"},
		{ID: 2, UserID: "2", Redirect: "someURL", Custom: "example1"},
	}

	repo.On("SearchingGoly", searchTerm).Return(mockResponse, nil)

	result, err := service.SearchGoly(searchTerm)
	repo.AssertExpectations(t)

    // Check if there are no errors
    assert.NoError(t, err)
    assert.Equal(t, mockResponse, result)
}

func TestIncreaseClickAndRedirect(t *testing.T){
	randomURL := mockHelpers.NewRandomURLInterface(t)
	repo := mocks.NewRepository(t)
	service := New(repo, randomURL)

	golyToIncrease := goly.Goly{ID: 1, Clicked: 5}

   
    repo.On("UpdateButton", mock.MatchedBy(func(goly goly.Goly) bool {
        return goly.ID == golyToIncrease.ID && goly.Clicked == golyToIncrease.Clicked+1
    })).Return(nil)


    err := service.IncreaseClickAndRedirect(golyToIncrease)


    repo.AssertExpectations(t)


    assert.NoError(t, err)
}

func TestExportIPToPDF(t *testing.T){
	randomURL := mockHelpers.NewRandomURLInterface(t)
	repo := mocks.NewRepository(t)
	service := New(repo, randomURL)

    mockIPAddresses := []goly.IPAdresses{
        {ID: 1, Address: "ip1"},
        {ID: 2, Address: "ip2"},
    }

    repo.On("PaginateIP", 1, 10).Return(mockIPAddresses)

    filePath, err := service.ExportIPToPDfAndSave()

    repo.AssertExpectations(t)

    assert.NoError(t, err)
    assert.NotEmpty(t, filePath)
}

func TestGetGolyByUrl(t *testing.T) {
	randomURL := mockHelpers.NewRandomURLInterface(t)
	repo := mocks.NewRepository(t)
	service := New(repo, randomURL)

	mockGoly := goly.Goly{
        ID: 1,
        ExpiryDate: time.Now().AddDate(0, 0, -1), // Expiry date 1 day ago (expired)
    }

    repo.On("FindByGolyUrl", "test_url").Return(mockGoly, nil)

    resultGoly, err := service.GetGolyByUrl("test_url")
    repo.AssertExpectations(t)

    assert.Error(t, err)
    assert.Equal(t, errors.New("Link is expired"), err)

    assert.Equal(t, mockGoly, resultGoly)
}