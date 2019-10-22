package mockapi

import (
	"timesheet/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) GetSummary(year, month int) ([]model.TransactionTimesheet, error) {
	argument := mock.Called(year, month)
	return argument.Get(0).([]model.TransactionTimesheet), argument.Error(1)
}

func (mock *MockRepository) CreateIncome(year, month int, memberID string, income model.Incomes) error {
	argument := mock.Called(year, month, memberID, income)
	return argument.Error(0)
}

func (mock *MockRepository) GetIncomes(memberID string, year, month int) ([]model.Incomes, error) {
	argument := mock.Called(memberID, year, month)
	return argument.Get(0).([]model.Incomes), argument.Error(1)
}

func (mock *MockRepository) GetMemberByID(memberID string) ([]model.Member, error) {
	argument := mock.Called(memberID)
	return argument.Get(0).([]model.Member), argument.Error(1)
}

func (mock *MockRepository) VerifyTransactionTimsheet(transactionTimesheet []model.TransactionTimesheet) error {
	argument := mock.Called(transactionTimesheet)
	return argument.Error(0)
}

func (mock *MockRepository) VerifyTimesheet(payment model.Timesheet, memberID string, year int, month int) error {
	argument := mock.Called(payment, memberID, year, month)
	return argument.Error(0)
}

type MockTimesheet struct {
	mock.Mock
}

func (mock *MockTimesheet) CalculatePayment(incomes []model.Incomes) model.Timesheet {
	argument := mock.Called(incomes)
	return argument.Get(0).(model.Timesheet)
}

func (mock *MockTimesheet) CalculatePaymentSummary(member []model.Member, incomes []model.Incomes, year, month int) []model.TransactionTimesheet {
	argument := mock.Called(member, incomes, year, month)
	return argument.Get(0).([]model.TransactionTimesheet)
}

func (mock *MockTimesheet) GetSummaryByID(memberID string, year, month int) (model.SummaryTimesheet, error) {
	argument := mock.Called(memberID, year, month)
	return argument.Get(0).(model.SummaryTimesheet), argument.Error(1)
}
