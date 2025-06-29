package services

import (
	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/handlers/dtos"
	r "github.com/alifrahmadian/alif-bookcabin-coding-test/internal/repositories"
	"github.com/alifrahmadian/alif-bookcabin-coding-test/pkg/constants"
	e "github.com/alifrahmadian/alif-bookcabin-coding-test/pkg/errors"
)

type SeatMapService interface {
	GetSeatMapBySeatsItineraryPartID(id int64) (*dtos.SeatMapResponse, error)
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

func (s *seatMapService) GetSeatMapBySeatsItineraryPartID(id int64) (*dtos.SeatMapResponse, error) {
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

	var segmentDTO dtos.Segment
	var passengerDTO dtos.Passenger
	var seatMapDTO dtos.SeatMap

	var seatsItineraryPartsDTO []dtos.SeatsItineraryPart
	var segmentSeatMapsDTO []dtos.SegmentSeatMap
	var passengerSeatMapsDTO []dtos.PassengerSeatMap
	var frequentFliersDTO []*dtos.FrequentFlyer
	var cabinsDTO []dtos.Cabin
	var seatRowsDTO []dtos.SeatRow
	var seatsDTO []dtos.Seat

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

			frequentFliers, err := s.FrequentFlyerRepo.GetFrequentFliersByPassengerID(passenger.ID)
			if err != nil {
				return nil, err
			}

			for _, frequentFlyer := range frequentFliers {
				frequentFliersDTO = append(frequentFliersDTO, &dtos.FrequentFlyer{
					Airline:    frequentFlyer.Airline,
					Number:     frequentFlyer.Number,
					TierNumber: uint(frequentFlyer.TierNumber),
				})
			}

			seatMap, err := s.SeatMapRepo.GetSeatMapByID(passengerSeatMap.SeatMapID)
			if err != nil {
				return nil, err
			}

			cabins, err := s.CabinRepo.GetCabinsByAircraft(seatMap.Aircraft)
			if err != nil {
				return nil, err
			}

			for _, cabin := range cabins {
				seatRows, err := s.SeatRowRepo.GetSeatRowsByCabinID(cabin.ID)
				if err != nil {
					return nil, err
				}

				for _, seatRow := range seatRows {
					seats, err := s.SeatRepo.GetSeatsBySeatRowIDAndSegmentID(seatRow.ID, segment.ID)
					if err != nil {
						return nil, err
					}

					for _, seat := range seats {
						if seat.StorefrontSlotCode != constants.CONST_SEAT_STOREFRONT_SLOT_CODE_SEAT {
							seatsDTO = append(seatsDTO, dtos.Seat{
								SlotCharacteristics: seat.SlotCharacteristics,
								StorefrontSlotCode:  seat.StorefrontSlotCode,
								Available:           seat.Available,
								Entitled:            seat.Entitled,
								FeeWaived:           seat.FeeWaived,
								FreeOfCharge:        seat.FreeOfCharge,
								OriginallySelected:  seat.OriginallySelected,
							})
						} else {
							seatsDTO = append(seatsDTO, dtos.Seat{
								StorefrontSlotCode:  seat.StorefrontSlotCode,
								Available:           seat.Available,
								Code:                seat.Code,
								Designations:        seat.Designations,
								Entitled:            seat.Entitled,
								FeeWaived:           seat.FeeWaived,
								EntitledRuleId:      seat.EntitledRuleID,
								FeeWaivedRuleId:     seat.FeeWaiveRuleID,
								SeatCharacteristics: seat.SeatCharacteristics,
								Limitations:         seat.Limitations,
								RefundIndicator:     seat.RefundIndicator,
								FreeOfCharge:        seat.FreeOfCharge,
								Prices: dtos.Prices{
									Alternatives: [][]dtos.Alternative{
										{
											{
												Amount:   seat.PriceAmount,
												Currency: seat.PriceCurrency,
											},
										},
									},
								},
								Taxes: dtos.Taxes{
									Alternatives: [][]dtos.Alternative{
										{
											{
												Amount:   seat.TaxAmount,
												Currency: seat.TaxCurrency,
											},
										},
									},
								},
								Total: dtos.Total{
									Alternatives: [][]dtos.Alternative{
										{
											{
												Amount:   seat.TotalAmount,
												Currency: seat.TotalCurrency,
											},
										},
									},
								},
								OriginallySelected:     seat.OriginallySelected,
								RawSeatCharacteristics: seat.RawSeatCharacteristics,
							})
						}
					}

					seatRowsDTO = append(seatRowsDTO, dtos.SeatRow{
						RowNumber: seatRow.RowNumber,
						SeatCodes: seatRow.SeatCodes,
						Seats:     seatsDTO,
					})
				}

				cabinsDTO = append(cabinsDTO, dtos.Cabin{
					Deck:        cabin.Deck,
					SeatColumns: cabin.SeatColumns,
					SeatRows:    seatRowsDTO,
					FirstRow:    cabin.FirstRow,
					LastRow:     cabin.LastRow,
				})
			}

			seatMapDTO = dtos.SeatMap{
				RowsDisabledCauses: seatMap.RowsDisabledCauses,
				Aircraft:           seatMap.Aircraft,
				Cabins:             cabinsDTO,
			}

			passengerDTO = dtos.Passenger{
				PassengerIndex:      passenger.ID,
				PassengerNameNumber: passenger.PassengerNameNumber,
				PassengerDetails: dtos.PassengerDetails{
					FirstName: passenger.PassengerDetails.FirstName,
					LastName:  passenger.PassengerDetails.LastName,
				},
				PassengerInfo: dtos.PassengerInfo{
					DateOfBirth: passenger.PassengerInfo.DateOfBirth.Format("2006-01-02"),
					Gender:      passenger.PassengerInfo.Gender,
					Type:        passenger.PassengerInfo.Type,
					Emails:      passenger.PassengerInfo.Emails,
					Phones:      passenger.PassengerInfo.Phones,
					Address: dtos.Address{
						Street1:     passenger.Address.Street1,
						Street2:     passenger.Address.Street2,
						Postcode:    passenger.Address.Postcode,
						State:       passenger.Address.State,
						City:        passenger.Address.City,
						Country:     passenger.Address.Country,
						AddressType: passenger.Address.AddressType,
					},
				},
				Preferences: dtos.Preference{
					SpecialPreferences: dtos.SpecialPreference{
						MealPreference:               passenger.Preferences.MealPreference,
						SeatPreference:               passenger.Preferences.SeatPreference,
						SpecialRequests:              passenger.Preferences.SpecialRequests,
						SpecialServiceRequestRemarks: passenger.Preferences.SpecialServiceRequestRemarks,
					},
					FrequentFlyer: frequentFliersDTO,
				},
				DocumentInfo: dtos.DocumentInfo{
					IssuingCountry: passenger.DocumentInfo.IssuingCountry,
					CountryOfBirth: passenger.DocumentInfo.CountryOfBirth,
					DocumentType:   passenger.DocumentInfo.DocumentType,
					Nationality:    passenger.DocumentInfo.Nationality,
				},
			}

			passengerSeatMapsDTO = append(passengerSeatMapsDTO, dtos.PassengerSeatMap{
				SeatSelectionEnabledForPax: passengerSeatMap.SeatSelectionEnabledForPax,
				SeatMap:                    seatMapDTO,
				Passenger:                  passengerDTO,
			})

		}

		segmentDTO = dtos.Segment{
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

		segmentSeatMapsDTO = append(segmentSeatMapsDTO, dtos.SegmentSeatMap{
			PassengerSeatMaps: passengerSeatMapsDTO,
			Segment:           segmentDTO,
		})
	}

	seatsItineraryPartsDTO = append(seatsItineraryPartsDTO, dtos.SeatsItineraryPart{
		SegmentSeatMaps: segmentSeatMapsDTO,
	})

	seatMapResponse := &dtos.SeatMapResponse{
		SeatsItineraryParts: seatsItineraryPartsDTO,
		SelectedSeats:       nil,
	}

	return seatMapResponse, nil
}
