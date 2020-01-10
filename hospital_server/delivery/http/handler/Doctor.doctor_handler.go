package handler

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"github.com/monday271/hospital_server/Doctor"
	//"github.com/yaredsolomon/webProgram1/hospital/request"
	"net/http"
	"strconv"

	//"github.com/betsegawlemma/restaurant-rest/comment"
	"github.com/julienschmidt/httprouter"
	_ "github.com/yaredsolomon/webProgram1/hospital/entity"
	//"github.com/yaredsolomon/webProgram1/hospital/request"
)

//"github.com/yaredsolomon/webProgram1/sathurday18/entity"

// DoctorAppointmentHandler handles appointment related http requests
type DoctorAppointmentHandler struct {
	appointmentService Doctor.AppointmentService
}

// NewDoctorAppointmentHandler returns new DoctorAppointmentHandler object
func NewDoctorAppointmentHandler(aptService Doctor.AppointmentService) *DoctorAppointmentHandler {
	return &DoctorAppointmentHandler{appointmentService: aptService}
}

// GetAppointments handles GET /v1/doctor/appointments request
func (dah *DoctorAppointmentHandler) GetAppointments(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	appointments, errs := dah.appointmentService.Appointments()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(appointments, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

// GetSingleAppointment handles GET /v1/doctor/appointments/:id request
func (dah *DoctorAppointmentHandler) GetSingleAppointment(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	fmt.Println(" i am about to get single Doctor value")
	fmt.Println(ps.ByName("id"))
	id, err := strconv.Atoi(ps.ByName("id"))
	//id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println(id)

	if err != nil {
		fmt.Println("ERRRRRRRRRRRR")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println(id)

	appointment, errs := dah.appointmentService.Appointment(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(appointment, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// PutAppointment handles PUT /v1/doctor/appointments/:id request
func (dah *DoctorAppointmentHandler) PutAppointment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" I am at the put method")
	id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	appointment, errs := dah.appointmentService.Appointment(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println(" i have get single value")

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &appointment)

	appointment, errs = dah.appointmentService.UpdateAppointment(appointment)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(appointment, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// DeleteComment handles DELETE /v1/doctor/appointments/:id request
func (dah *DoctorAppointmentHandler) DeleteAppointment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := dah.appointmentService.DeleteAppointment(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
