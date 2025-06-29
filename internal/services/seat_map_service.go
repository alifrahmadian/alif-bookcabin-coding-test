package services

import (
	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/handlers/dtos"
	r "github.com/alifrahmadian/alif-bookcabin-coding-test/internal/repositories"
	e "github.com/alifrahmadian/alif-bookcabin-coding-test/pkg/errors"
)

type SeatMapService interface {
	GetSeatMapBySeatsItineraryPartID(id int64) ([]*dtos.SeatMapResponse, error)
}

type seatMapService struct {
	CabinRepo              r.CabinRepository
	FrequentFlyerRepo      r.FrequentFlyerRepository
	PassengerRepo          r.PassengerRepository
	PassengerSeatMapRepo   r.PassengerSeatMapRepository
	SeatMapRepo            r.SeatMapRepository
	SeatRepo               r.SeatRepository
	SeatRowRepo            r.SeatRowRepository
	SeatsItineraryPartRepo r.SeatsItineraryPartRepository
	SegmentRepo            r.SegmentRepository
	SegmentSeatMapRepo     r.SegmentSeatMapRepository
}

func NewSeatMapService(
	cabinRepo r.CabinRepository,
	frequentFlyerRepo r.FrequentFlyerRepository,
	passengerRepo r.PassengerRepository,
	passengerSeatMapRepo r.PassengerSeatMapRepository,
	seatMapRepo r.SeatMapRepository,
	seatRepo r.SeatRepository,
	seatRowRepo r.SeatRowRepository,
	seatsItineraryPartRepo r.SeatsItineraryPartRepository,
	segmentRepo r.SegmentRepository,
	segmentSeatMapRepo r.SegmentSeatMapRepository,
) SeatMapService {
	return &seatMapService{
		CabinRepo:              cabinRepo,
		FrequentFlyerRepo:      frequentFlyerRepo,
		PassengerRepo:          passengerRepo,
		PassengerSeatMapRepo:   passengerSeatMapRepo,
		SeatMapRepo:            seatMapRepo,
		SeatRepo:               seatRepo,
		SeatRowRepo:            seatRowRepo,
		SeatsItineraryPartRepo: seatsItineraryPartRepo,
		SegmentRepo:            segmentRepo,
		SegmentSeatMapRepo:     segmentSeatMapRepo,
	}
}

func (s *seatMapService) GetSeatMapBySeatsItineraryPartID(id int64) ([]*dtos.SeatMapResponse, error) {
	isSeatsItineraryPartIDExist, err := s.SeatsItineraryPartRepo.CheckIfExist(id)
	if err != nil {
		return nil, err
	}

	if !isSeatsItineraryPartIDExist {
		return nil, e.ErrSeatMapSeatsItineraryPartNotExist
	}

	segmentSeatMaps, err := s.SegmentSeatMapRepo.GetSegmentSeatMapsBySeatsItineraryPartID(id)
	if err != nil {
		return nil, err
	}

	var segmentDTO *dtos.Segment
	var passengerDTO dtos.Passenger
	var passengerSeatMapsDTO []*dtos.PassengerSeatMap

	// populate segmentSeatMaps response
	for _, segmentSeatMap := range segmentSeatMaps {
		segment, err := s.SegmentRepo.GetSegmentAndFlightBySegmentID(segmentSeatMap.ID)
		if err != nil {
			return nil, err
		}

		passengerSeatMaps, err := s.PassengerSeatMapRepo.GetPassengerSeatMapsBySegmentSeatMapID(segmentSeatMap.SegmentID)
		if err != nil {
			return nil, err
		}

		for _, passengerSeatMap := range passengerSeatMaps {
			passenger, err := s.PassengerRepo.GetPassengerByID(passengerSeatMap.PassengerID)
			if err != nil {
				return nil, err
			}

			passengerDTO = &dtos.Passenger{}

		}

		segmentDTO = &dtos.Segment{
			Type: segment.Type,
			SegmentOfferInformation: dtos.SegmentOfferInformation{
				FlightMiles: segment.SegmentOfferInformation.FlightMiles,
				AwardFare:   segment.SegmentOfferInformation.AwardFlight,
			},
			Duration:   segment.Duration,
			CabinClass: segment.CabinClass,
			Equipment:  segment.Equipment,
			Flight: dtos.Flight{
				FlightNumber:          segment.Flight.FlightNumber,
				OperatingFlightNumber: segment.Flight.OperatingFlightNumber,
				AirlineCode:           segment.Flight.AirlineCode,
				OperatingAirlineCode:  segment.Flight.OperatingAirlineCode,
				StopAirports:          segment.Flight.StopAirports,
				DepartureTerminal:     segment.Flight.DepartureTerminal,
				ArrivalTerminal:       segment.Flight.ArrivalTerminal,
			},
			Origin:                      segment.Origin,
			Destination:                 segment.Destination,
			Departure:                   segment.Departure.Format("2006-01-02T15:04:05"),
			Arrival:                     segment.Arrival.Format("2006-01-02T15:04:05"),
			BookingClass:                segment.BookingClass,
			LayoverDuration:             segment.LayoverDuration,
			FareBasis:                   segment.FareBasis,
			SubjectToGovernmentApproval: segment.SubjectToGovernmentApproval,
			SegmentRef:                  segment.SegmentRef,
		}

	}

	return nil, nil
}
