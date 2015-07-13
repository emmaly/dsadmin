package dsadmin

// User is a Windows Domain user
type User struct {
	DN           string       `ldif:"dn"`                          // CN=blah,CN=blah,...
	SAMID        string       `ldif:"sAMAccountName,omitempty"`    // legacy username (not UPN)
	UPN          string       `ldif:"userPrincipalName,omitempty"` // the full UPN, not just the suffix
	NameDisplay  string       `ldif:"displayName,omitempty"`
	NameFirst    string       `ldif:"givenName,omitempty"`
	NameLast     string       `ldif:"sn,omitempty"`
	Description  string       `ldif:"description,omitempty"`
	Title        string       `ldif:"title,omitempty"`
	EmployeeID   string       `ldif:"employeeID,omitempty"`
	Company      string       `ldif:"company,omitempty"`
	Office       string       `ldif:"physicalDeliveryOfficeName,omitempty"`
	Department   string       `ldif:"department,omitempty"`
	Manager      string       `ldif:"manager,omitempty"` // this is the manager's DN
	PhoneOffice  string       `ldif:"telephoneNumber,omitempty"`
	PhoneHome    string       `ldif:"homePhone,omitempty"`
	PhonePager   string       `ldif:"pager,omitempty"`
	PhoneMobile  string       `ldif:"mobile,omitempty"`
	PhoneFax     string       `ldif:"facsimileTelephoneNumber,omitempty"`
	PhoneIPPhone string       `ldif:"ipPhone,omitempty"`
	Email        string       `ldif:"mail,omitempty"`
	URL          string       `ldif:"wWWHomePage,omitempty"`
	HomeDrive    string       `ldif:"homeDrive,omitempty"`     // drive letter for home drive; used with HomeDir
	HomeDir      string       `ldif:"homeDirectory,omitempty"` // this is the path to their home directory, probably UNC; used for HomeDrive
	Profile      string       `ldif:"profilePath,omitempty"`   // user's profile path, probably UNC (for roaming profiles?)
	LoginScript  string       `ldif:"scriptPath,omitempty"`
	AccountFlags AccountFlags `ldif:"userAccountControl"`       // password policy, etc; see also https://support.microsoft.com/en-us/kb/305144
	AcctExpires  int64        `ldif:"accountExpires,omitempty"` // nanoseconds?
	MemberOf     []string     // set of DNs
}

// AddUser creates a new user in the domain
func AddUser(user User, password string) {
	//cmd := exec.Command("dsadd", "user", username, password, "")
	// sAMAccountType: 805306368
}
