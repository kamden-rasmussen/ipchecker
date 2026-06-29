// Package empty implements the EmailProvider interface. All methods intentionally left blank for those who do not want emails to be sent
package empty

type EmptyProvider struct{}

func (e EmptyProvider) SendEmail(newIP string) error {
	return nil
}

func (e EmptyProvider) SendErrorEmail() error {
	return nil
}

func (e EmptyProvider) SendDNSErrorEmail() error {
	return nil
}
