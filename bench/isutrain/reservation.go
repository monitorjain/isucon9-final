package isutrain

import (
	"errors"
	"time"
)

var (
	ErrNoSeats = errors.New("座席がありません")
)

// TODO: Statusのconst

// PaymentMethod は決済方法です
type PaymentMethod string

const (
	// CreditCard はクレカ決済
	CreditCard PaymentMethod = "creditcard"
	// TicketVendingMachine は券売機決済
	TicketVendingMachine PaymentMethod = "ticket-vending-machine"
)

// ReservationStatus は予約状態です
type ReservationStatus string

const (
	// Pending は予約が確定していない状態です
	Pending ReservationStatus = "pending"
	// Ok は予約が確定した状態です
	Ok ReservationStatus = "ok"
)

// SeatReservation は座席予約です
type SeatReservation struct {
	// ID は採番された予約IDです
	ID int `json:"id"`
	// PaymentMethod は決済種別です
	PaymentMethod string `json:"payment_method"`
	// Status は予約の決済状況です
	Status string `json:"status"`
	// ReserveAt は予約時刻です
	ReserveAt time.Time `json:"reserve_at"`
}

// SeatReservations は座席一覧です
type SeatReservations []*SeatReservation

type ReservationRequest struct {
	// Train構造体
	TrainClass string `json:"train_class"`
	TrainName  string `json:"train_name"`
	// TrainSeat構造体
	SeatClass string     `json:"seat_class"`
	Seats     TrainSeats `json:"seats"`
	// ログインで得る情報
	UserId int `json:"user_id"`
	// それ以外
	//// 区間
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	//// 未調査
	Date    string `json:"date"`
	CarNum  int    `json:"car_num"`
	Payment string `json:"payment"`
	Child   int    `json:"child"`
	Adult   int    `json:"adult"`
	Type    string `json:"type"`
}

func NewReservationRequest(trains *TrainSearchResponse, seats TrainSeats) (*ReservationRequest, error) {
	req := &ReservationRequest{}

	// Train構造体
	req.TrainClass = trains.Class
	req.TrainName = trains.Name

	// TrainSeat構造体
	if len(seats) == 0 {
		return nil, ErrNoSeats
	}
	req.SeatClass = seats[0].Class
	req.Seats = seats

	return req, nil
}

type ReservationResponse struct {
	ReservationID string `json:"reservation_id"`
	IsOk          bool   `json:"is_ok"`
}