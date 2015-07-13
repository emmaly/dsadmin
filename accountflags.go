package dsadmin

import "fmt"

// https://support.microsoft.com/en-us/kb/305144

// AccountFlag is an int
type AccountFlag int // provides type

// AccountFlags provides enum for AccountFlags value in User
type AccountFlags AccountFlag // provides enum

// AccountFlags constants
const (
	AFScript                       AccountFlag = 0x1       // The logon script will be run.
	AFAccountDisable                           = 0x2       // The user account is disabled.
	AFHomeDirRequired                          = 0x8       // The home folder is required.
	AFLockout                                  = 0x10      // The account is currently locked out.
	AFPasswordNotRequired                      = 0x20      // No password is required.
	AFPasswordCantChange                       = 0x40      // see caveat at https://support.microsoft.com/en-us/kb/305144 or https://msdn.microsoft.com/en-us/library/ms680832.aspx; The user cannot change the password. This is a permission on the user's object. For information about how to programmatically set this permission, visit the following Web site: http://msdn2.microsoft.com/en-us/library/aa746398.aspx
	AFEncryptedTextPasswordAllowed             = 0x80      // The user can send an encrypted password.
	AFTemporaryDuplicateAllowed                = 0x100     // This is an account for users whose primary account is in another domain. This account provides user access to this domain, but not to any domain that trusts this domain. This is sometimes referred to as a local user account.
	AFNormalAccount                            = 0x200     // This is a default account type that represents a typical user.
	AFInterdomainTrustAccount                  = 0x800     // This is a permit to trust an account for a system domain that trusts other domains.
	AFWorkstationTrustAccount                  = 0x1000    // This is a computer account for a computer that is running Microsoft Windows NT 4.0 Workstation, Microsoft Windows NT 4.0 Server, Microsoft Windows 2000 Professional, or Windows 2000 Server and is a member of this domain.
	AFServerTrustAccount                       = 0x2000    // This is a computer account for a domain controller that is a member of this domain.
	AFDontExpirePassword                       = 0x10000   // Represents the password, which should never expire on the account.
	AFMNSLogonAccount                          = 0x20000   // This is an MNS logon account.
	AFSmartcardRequired                        = 0x40000   // When this flag is set, it forces the user to log on by using a smart card.
	AFTrustedForDelegation                     = 0x80000   // When this flag is set, the service account (the user or computer account) under which a service runs is trusted for Kerberos delegation. Any such service can impersonate a client requesting the service. To enable a service for Kerberos delegation, you must set this flag on the userAccountControl property of the service account.
	AFNotDelegated                             = 0x100000  // When this flag is set, the security context of the user is not delegated to a service even if the service account is set as trusted for Kerberos delegation.
	AFUseDESKeyOnly                            = 0x200000  // (Windows 2000/Windows Server 2003) Restrict this principal to use only Data Encryption Standard (DES) encryption types for keys.
	AFDontRequirePreauth                       = 0x400000  // (Windows 2000/Windows Server 2003) This account does not require Kerberos pre-authentication for logging on.
	AFPasswordExpired                          = 0x800000  // (Windows 2000/Windows Server 2003) The user's password has expired.
	AFTrustedToAuthForDelegation               = 0x1000000 // (Windows 2000/Windows Server 2003) The account is enabled for delegation. This is a security-sensitive setting. Accounts that have this option enabled should be tightly controlled. This setting lets a service that runs under the account assume a client's identity and authenticate as that user to other remote servers on the network.
	AFPartialSecretsAccount                    = 0x4000000 // (Windows Server 2008/Windows Server 2008 R2) The account is a read-only domain controller (RODC). This is a security-sensitive setting. Removing this setting from an RODC compromises security on that server.
)

// Set will set a flag
func (a *AccountFlags) Set(accountFlag AccountFlag) {
	*a |= AccountFlags(accountFlag) // set the flag
}

// Unset will set a flag
func (a *AccountFlags) Unset(accountFlag AccountFlag) {
	*a &^= AccountFlags(accountFlag) // unset the flag
}

// Has will return true if the flag is enabled
func (a *AccountFlags) Has(accountFlag AccountFlag) bool {
	return AccountFlag(*a)&accountFlag == accountFlag // test for flag
}

// String returns a string
func (a *AccountFlags) String() string {
	return fmt.Sprintf("0x%x", int(*a))
}
