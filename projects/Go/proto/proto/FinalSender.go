// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// Version: 1.1.0.0

package proto

import "errors"
import "../fbe"

// Workaround for Go unused imports issue
var _ = fbe.Version

// Fast Binary Encoding proto final sender
type FinalSender struct {
    fbe.Sender
    orderModel *OrderFinalModel
    balanceModel *BalanceFinalModel
    accountModel *AccountFinalModel
}

// Create a new proto final sender
func NewFinalSender(buffer *fbe.Buffer) *FinalSender {
    return &FinalSender{
        *fbe.NewSender(buffer, false, false),
        NewOrderFinalModel(buffer),
        NewBalanceFinalModel(buffer),
        NewAccountFinalModel(buffer),
    }
}

// Sender models accessors

func (s *FinalSender) OrderModel() *OrderFinalModel { return s.orderModel }
func (s *FinalSender) BalanceModel() *BalanceFinalModel { return s.balanceModel }
func (s *FinalSender) AccountModel() *AccountFinalModel { return s.accountModel }

// Send methods

func (s *FinalSender) Send(value interface{}) (int, error) {
    switch value := value.(type) {
    case *Order:
        return s.SendOrder(value)
    case *Balance:
        return s.SendBalance(value)
    case *Account:
        return s.SendAccount(value)
    }
    return 0, nil
}

func (s *FinalSender) SendOrder(value *Order) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.orderModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("proto.Order serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.orderModel.Verify() {
        return 0, errors.New("proto.Order validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        if err := s.OnSendLogHandler.OnSendLog(message); err != nil {
            return 0, err
        }
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *FinalSender) SendBalance(value *Balance) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.balanceModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("proto.Balance serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.balanceModel.Verify() {
        return 0, errors.New("proto.Balance validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        if err := s.OnSendLogHandler.OnSendLog(message); err != nil {
            return 0, err
        }
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *FinalSender) SendAccount(value *Account) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.accountModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("proto.Account serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.accountModel.Verify() {
        return 0, errors.New("proto.Account validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        if err := s.OnSendLogHandler.OnSendLog(message); err != nil {
            return 0, err
        }
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}