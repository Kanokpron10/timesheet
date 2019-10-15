package timesheet_test

import (
	"database/sql"
	"testing"
	"timesheet/internal/timesheet"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
)

func Test_GetSummary_Input_Year_2018_Month_17_Should_Be_TransactionTimesheet(t *testing.T) {
	expected := []timesheet.TransactionTimesheet{
		{
			ID:                     1,
			MemberID:               "001",
			MemberNameTH:           "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  12,
			Year:                   2017,
			Company:                "siam_chamnankit",
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
	}
	dBConnection, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer dBConnection.Close()
	year := 2017
	month := 12
	repository := timesheet.TimesheetRepository{
		DBConnection: dBConnection,
	}

	actual, err := repository.GetSummary(year, month)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_UpdateIncomeByID_Input_Year_2018_Month_12_MemberID_001_Income_Should_Be_True(t *testing.T) {
	expected := true

	dBConnection, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer dBConnection.Close()

	year := 2018
	month := 12
	memberID := "001"
	incomes := []timesheet.Incomes{
		{
			Day:                      28,
			StartTimeAM:              "09:00:00",
			EndTimeAM:                "12:00:00",
			StartTimePM:              "13:00:00",
			EndTimePM:                "18:00:00",
			Overtime:                 0,
			TotalHours:               8,
			CoachingCustomerCharging: 15000,
			CoachingPaymentRate:      10000,
			TrainingWage:             0,
			OtherWage:                0,
			Company:                  "siam_chamnankit",
			Description:              "[KBTG] 2 Days Agile Project Management",
		},
	}

	repository := timesheet.TimesheetRepository{
		DBConnection: dBConnection,
	}

	actual := repository.UpdateIncomeByID(year, month, memberID, incomes)

	assert.Equal(t, expected, actual)
}
