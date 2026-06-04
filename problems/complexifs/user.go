// Package complexifs ports the "deeply nested if" smell. User is a data object with many
// fields; UserProcessor.ProcessUser branches on them.
//
// Note: Go has no nil string, so the Java fields/parameters that were nullable Strings
// (SuspensionReason, userType, department) are modelled as plain strings where the empty
// string "" stands in for Java's null.
package complexifs

type User struct {
	Age                int
	Experience         int
	Certifications     []string
	LastLoginDays      int
	FailedLogins       int
	BackgroundCheck    bool
	ComplianceTraining bool
	FinancialClearance bool
	AuditScore         int
	GeneralTraining    bool
	AccountAge         int
	TrialDaysRemaining int
	SuspensionReason   string
	ParentalConsent    bool
}
