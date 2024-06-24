package jamfprointegration

// GetFqdn returns the base domain for the Jamf Pro integration.
func (j *Integration) GetFqdn() string {
	return j.Fqdn
}
