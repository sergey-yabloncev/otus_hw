package memorystorage_test

import (
	"context"
	"sync"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/golang/mock/gomock"
	"github.com/sergey-yabloncev/otus_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/sergey-yabloncev/otus_hw/hw12_13_14_15_calendar/internal/storage"
	memory "github.com/sergey-yabloncev/otus_hw/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/stretchr/testify/suite"
)

type AppSuite struct {
	suite.Suite
	logger memory.Logger
}

func (s *AppSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()
	s.logger = logger.NewMockContract(ctrl)
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(AppSuite))
}

func (s *AppSuite) TestCreateEvent() {
	memoryStorage := memory.New(s.logger)
	event := storage.Event{}
	err := faker.FakeData(&event)
	s.Require().NoError(err)
	err = memoryStorage.CreateEvent(context.Background(), event)
	s.Require().NoError(err)
	events, err := memoryStorage.GetEvents(context.Background())
	s.Require().NoError(err)
	s.Require().Equal(1, len(events))
	s.Require().Equal(event, events[0])
}

func (s *AppSuite) TestGetEvent() {
	memoryStorage := memory.New(s.logger)
	event := storage.Event{}
	err := faker.FakeData(&event)
	s.Require().NoError(err)
	err = memoryStorage.CreateEvent(context.Background(), event)
	s.Require().NoError(err)

	events, err := memoryStorage.GetEvents(context.Background())
	s.Require().NoError(err)
	s.Require().Equal(event, events[0])

	newEvent := storage.Event{}
	faker.FakeData(&newEvent)
	newEvent.ID = event.ID
	err = memoryStorage.UpdateEvent(context.Background(), newEvent)
	s.Require().NoError(err)
	s.Require().NotEqual(event, newEvent)

	events, err = memoryStorage.GetEvents(context.Background())
	s.Require().NoError(err)
	s.Require().Equal(newEvent, events[0])
}

func (s *AppSuite) TestDeleteEvent() {
	memoryStorage := memory.New(s.logger)
	event := storage.Event{}
	faker.FakeData(&event)
	memoryStorage.CreateEvent(context.Background(), event)
	events, _ := memoryStorage.GetEvents(context.Background())
	s.Require().Equal(1, len(events))
	memoryStorage.DeleteEvent(context.Background(), event)
	events, _ = memoryStorage.GetEvents(context.Background())
	s.Require().Equal(0, len(events))
}

func (s *AppSuite) TestMultithreading() {
	memoryStorage := memory.New(s.logger)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1_000; i++ {
			event := storage.Event{}
			faker.FakeData(&event)
			err := memoryStorage.CreateEvent(context.Background(), event)
			s.Require().NoError(err)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000; i++ {
			event := storage.Event{}
			faker.FakeData(&event)
			err := memoryStorage.CreateEvent(context.Background(), event)
			s.Require().NoError(err)
		}
	}()
	wg.Wait()

	events, err := memoryStorage.GetEvents(context.Background())
	s.Require().NoError(err)
	s.Require().Equal(2_000, len(events))
}
