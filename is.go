package useragent

func (ua *UserAgent) IsMobile() bool {
	return ua.Device.Mobile
}

func (ua *UserAgent) IsTablet() bool {
	return ua.Device.Tablet
}
