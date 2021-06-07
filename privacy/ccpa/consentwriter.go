package ccpa

import (
	"github.com/mxmCherry/openrtb/v15/openrtb2"
	"github.com/prebid/prebid-server/openrtb_ext"
)

// ConsentWriter implements the PolicyWriter interface for CCPA.
type ConsentWriter struct {
	Consent string
}

// Write mutates an OpenRTB bid request with the CCPA consent string.
func (c ConsentWriter) Write(req *openrtb_ext.RequestWrapper) {
	if req == nil || req.Request == nil {
		return
	}
	req.RegExt.SetUSPrivacy(c.Consent)
}

// ConsentWriterLegacy implements the old PolicyWriter interface for CCPA.
// This is used where we have not converted to RequestWrapper yet
type ConsentWriterLegacy struct {
	Consent string
}

// Write mutates an OpenRTB bid request with the CCPA consent string.
func (c ConsentWriterLegacy) Write(req *openrtb2.BidRequest) error {
	if req == nil {
		return nil
	}
	reqWrap := &openrtb_ext.RequestWrapper{Request: req}
	if err := reqWrap.ExtractRegExt(); err == nil {
		reqWrap.RegExt.SetUSPrivacy(c.Consent)
	} else {
		return err
	}
	return reqWrap.Sync()
}
