package http

import (
	"bytes"
	"encoding/json"
	ps "service-hf-orch-p5/external/strings"
	"service-hf-orch-p5/internal/core/domain/entity"
	"service-hf-orch-p5/internal/core/domain/entity/dto"
	"fmt"
	"net/http"
	"time"
)

func (h *handler) getVoucherByID(rw http.ResponseWriter, req *http.Request) {
	id := getID("voucher", req.URL.Path)

	v, err := h.app.GetVoucherByID(id)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save voucher: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(ps.MarshalString(v)))
}

func (h *handler) saveVoucher(rw http.ResponseWriter, req *http.Request) {
	var buff bytes.Buffer

	var reqVoucher dto.RequestVoucher

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqVoucher); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}
	voucher := entity.Voucher{
		Code:       reqVoucher.Code,
		Percentage: reqVoucher.Percentage,
	}

	if len(reqVoucher.ExpiresAt) > 0 {
		voucher.ExpiresAt.Value = new(time.Time)
		if err := voucher.ExpiresAt.SetTimeFromString(reqVoucher.ExpiresAt); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, `{"error": "error to save voucher: %v"} `, err)
			return
		}
	}

	reqVoucher.ExpiresAt = voucher.ExpiresAt.Format()

	v, err := h.app.SaveVoucher(reqVoucher)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save voucher: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(ps.MarshalString(v)))
}

func (h *handler) updateVoucherByID(rw http.ResponseWriter, req *http.Request) {
	id := getID("voucher", req.URL.Path)

	var buff bytes.Buffer

	var reqVoucher dto.RequestVoucher

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqVoucher); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	v, err := h.app.UpdateVoucherByID(id, reqVoucher)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to update voucher: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(ps.MarshalString(v)))
}
