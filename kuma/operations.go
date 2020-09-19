package kuma

import (
	"context"
	"fmt"

	cfg "github.com/layer5io/meshery-kuma/internal/config"
	"github.com/mgfeller/common-adapter-library/adapter"
)

// ApplyOperation applies the operation on kuma
func (h *KumaAdapter) ApplyOperation(ctx context.Context, op string, id string, del bool) error {

	operations := make(adapter.Operations, 0)
	err := h.Config.Operations(&operations)
	if err != nil {
		return err
	}

	status := "deploying"
	e := &adapter.Event{
		Operationid: id,
		Summary:     "Deploying",
		Details:     "None",
	}

	switch op {
	case cfg.InstallKumav071, cfg.InstallKumav070, cfg.InstallKumav060:
		go func(hh *KumaAdapter, ee *adapter.Event) {
			if status, err := hh.installKuma(del, operations[op].Properties["version"]); err != nil {
				e.Summary = fmt.Sprintf("Error while %s Kuma service mesh", status)
				e.Details = err.Error()
				hh.StreamErr(e, err)
				return
			}
			ee.Summary = fmt.Sprintf("Kuma service mesh %s successfully", status)
			ee.Details = fmt.Sprintf("The Kuma service mesh is now %s.", status)
			hh.StreamInfo(e)
		}(h, e)
	case cfg.InstallSampleBookInfo:
		go func(hh *KumaAdapter, ee *adapter.Event) {
			if status, err := hh.installSampleApp(operations[op].Properties["description"]); err != nil {
				e.Summary = fmt.Sprintf("Error while %s Sample %s application", status, operations[op].Properties["description"])
				e.Details = err.Error()
				hh.StreamErr(e, err)
				return
			}
			ee.Summary = fmt.Sprintf("Sample %s application %s successfully", operations[op].Properties["description"], status)
			ee.Details = fmt.Sprintf("The Sample %s application is now %s.", operations[op].Properties["description"], status)
			hh.StreamInfo(e)
		}(h, e)
	case cfg.ValidateSmiConformance:
		go func(hh *KumaAdapter, ee *adapter.Event) {
			err := hh.validateSMIConformance(ee.Operationid)
			if err != nil {
				return
			}
		}(h, e)
	default:
		h.StreamErr(e, adapter.ErrOpInvalid)
	}

	return nil
}
