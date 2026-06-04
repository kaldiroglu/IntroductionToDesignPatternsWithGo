package complexifs

type UserProcessor struct{}

// ProcessUser is a ridiculously complicated method that processes user data with deeply
// nested if statements - an example of what NOT to do!
func (UserProcessor) ProcessUser(user *User, isAdmin bool, userType string,
	accessLevel int, hasPermissions bool, department string, isActive bool) string {

	if user != nil {
		if user.Age >= 18 {
			if isActive {
				if userType != "" {
					if userType == "PREMIUM" {
						if isAdmin {
							if accessLevel > 5 {
								if hasPermissions {
									if department != "" {
										if department == "IT" {
											if user.Experience > 2 {
												if len(user.Certifications) > 3 {
													if user.LastLoginDays < 30 {
														if user.FailedLogins < 3 {
															return "FULL_ACCESS_GRANTED"
														} else if user.FailedLogins < 5 {
															return "LIMITED_ACCESS_SECURITY_REVIEW"
														} else {
															return "ACCESS_BLOCKED_TOO_MANY_FAILURES"
														}
													} else if user.LastLoginDays < 90 {
														return "ACCESS_GRANTED_PASSWORD_RESET_REQUIRED"
													} else {
														return "ACCOUNT_DORMANT_REACTIVATION_NEEDED"
													}
												} else if len(user.Certifications) > 1 {
													return "PARTIAL_ACCESS_CERTIFICATION_PENDING"
												} else {
													return "TRAINING_REQUIRED"
												}
											} else if user.Experience > 1 {
												return "SUPERVISED_ACCESS_ONLY"
											} else {
												return "INTERN_ACCESS_BASIC_ONLY"
											}
										} else if department == "HR" {
											if user.BackgroundCheck {
												if user.ComplianceTraining {
													return "HR_FULL_ACCESS"
												}
												return "HR_LIMITED_COMPLIANCE_TRAINING_NEEDED"
											}
											return "HR_ACCESS_DENIED_BACKGROUND_CHECK"
										} else if department == "FINANCE" {
											if user.FinancialClearance {
												if user.AuditScore > 85 {
													return "FINANCE_FULL_ACCESS"
												} else if user.AuditScore > 70 {
													return "FINANCE_RESTRICTED_ACCESS"
												} else {
													return "FINANCE_AUDIT_REQUIRED"
												}
											}
											return "FINANCE_CLEARANCE_PENDING"
										} else {
											if user.GeneralTraining {
												return "GENERAL_DEPARTMENT_ACCESS"
											}
											return "BASIC_ACCESS_TRAINING_REQUIRED"
										}
									}
									return "ACCESS_DENIED_NO_DEPARTMENT"
								} else if accessLevel > 3 {
									return "MODERATE_ACCESS_NO_PERMISSIONS"
								} else {
									return "BASIC_ACCESS_ONLY"
								}
							} else if accessLevel > 2 {
								return "NON_ADMIN_MODERATE_ACCESS"
							} else {
								return "NON_ADMIN_BASIC_ACCESS"
							}
						} else {
							if hasPermissions {
								if accessLevel > 3 {
									return "NON_ADMIN_PREMIUM_HIGH_ACCESS"
								}
								return "NON_ADMIN_PREMIUM_STANDARD_ACCESS"
							}
							return "PREMIUM_USER_LIMITED_ACCESS"
						}
					} else if userType == "STANDARD" {
						if isAdmin {
							if accessLevel > 4 {
								return "ADMIN_STANDARD_HIGH_ACCESS"
							}
							return "ADMIN_STANDARD_NORMAL_ACCESS"
						} else {
							if hasPermissions {
								return "STANDARD_USER_WITH_PERMISSIONS"
							} else if user.AccountAge > 365 {
								return "STANDARD_USER_VETERAN"
							} else {
								return "STANDARD_USER_BASIC"
							}
						}
					} else if userType == "TRIAL" {
						if user.TrialDaysRemaining > 0 {
							if user.TrialDaysRemaining > 7 {
								return "TRIAL_ACTIVE_FULL_FEATURES"
							}
							return "TRIAL_EXPIRING_SOON_LIMITED"
						}
						return "TRIAL_EXPIRED_UPGRADE_REQUIRED"
					} else {
						return "UNKNOWN_USER_TYPE_DEFAULT_ACCESS"
					}
				}
				return "NULL_USER_TYPE_ACCESS_DENIED"
			} else {
				if user.SuspensionReason != "" {
					if user.SuspensionReason == "TEMPORARY" {
						return "ACCOUNT_TEMPORARILY_SUSPENDED"
					}
					return "ACCOUNT_PERMANENTLY_SUSPENDED"
				}
				return "INACTIVE_ACCOUNT_REASON_UNKNOWN"
			}
		} else if user.Age >= 16 {
			if user.ParentalConsent {
				return "MINOR_ACCESS_WITH_CONSENT"
			}
			return "MINOR_ACCESS_PARENTAL_CONSENT_REQUIRED"
		} else {
			return "ACCESS_DENIED_TOO_YOUNG"
		}
	}
	return "NULL_USER_ACCESS_DENIED"
}
