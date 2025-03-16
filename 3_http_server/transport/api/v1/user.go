package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"

	"go_school/3_http_server/service/errors"
	"go_school/3_http_server/service/model"
	"go_school/3_http_server/transport/util"
)                                                                                                                                                                                                            
																																																										
var validate = validator.New()

func init() {                                                                                                                                                                                                                                                                                                                                                                                                                    
	validate.RegisterStructValidation(birthDateValidator, model.BirthDate{})                                                                                                                                                            
}                                                                                                                                                                                                                                       
																																																										
func birthDateValidator(sl validator.StructLevel) {                                                                                                                                                                                     
	birthDate := sl.Current().Interface().(model.BirthDate)                                                                                                                                                                             
																																																										
	// Validate month (1-12)                                                                                                                                                                                                            
	if birthDate.Month < 1 || birthDate.Month > 12 {                                                                                                                                                                                    
		sl.ReportError(birthDate.Month, "Month", "month", "validmonth", "")                                                                                                                                                             
		return                                                                                                                                                                                                                          
	}                                                                                                                                                                                                                                   
																																																										
	// Validate day based on month and year                                                                                                                                                                                             
	maxDay := 31                                                                                                                                                                                                                        
																																																										
	switch birthDate.Month {                                                                                                                                                                                                            
	case 4, 6, 9, 11:                                                                                                                                                                                                                   
		maxDay = 30                                                                                                                                                                                                                     
	case 2:                                                                                                                                                                                                                             
		// Check for leap year                                                                                                                                                                                                          
		if birthDate.Year%4 == 0 && (birthDate.Year%100 != 0 || birthDate.Year%400 == 0) {                                                                                                                                              
			maxDay = 29                                                                                                                                                                                                                 
		} else {                                                                                                                                                                                                                        
			maxDay = 28                                                                                                                                                                                                                 
		}                                                                                                                                                                                                                               
	}                                                                                                                                                                                                                                   
																																																										
	if birthDate.Day < 1 || birthDate.Day > maxDay {                                                                                                                                                                                    
		sl.ReportError(birthDate.Day, "Day", "day", "validday", fmt.Sprintf("must be between 1 and %d for month %d", maxDay, birthDate.Month))                                                                                          
	}                                                                                                                                                                                                                                   
}
																																																										
func getEmailFromURL(r *http.Request) string {                                                                                                                                                                                          
	email := chi.URLParam(r, "email")                                                                                                                                                                                                   
	return email                                                                                                                                                                                                                        
}                                                                                                                                                                                                                                       
																																																										
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {                                                                                                                                                                  
	b, err := io.ReadAll(r.Body)                                                                                                                                                                                                        
	if err != nil {                                                                                                                                                                                                                     
		util.WriteErrResponse(w, http.StatusBadRequest, err)                                                                                                                                                                            
		return                                                                                                                                                                                                                          
	}                                                                                                                                                                                                                                   
																																																										
	var user model.User                                                                                                                                                                                                                 
	if err := json.Unmarshal(b, &user); err != nil {                                                                                                                                                                                    
		util.WriteErrResponse(w, http.StatusBadRequest, err)                                                                                                                                                                            
		return                                                                                                                                                                                                                          
	}                                                                                                                                                                                                                                   
																																																										
	if err := validate.Struct(user); err != nil {                                                                                                                                                                                                                                                                                                                                                                                
		util.WriteErrResponse(w, http.StatusBadRequest, err)                                                                                                                                                           
		return                                                                                                                                                                                                                          
	}                                                                                                                                                                                                                                 
																																																										
	if err := h.service.CreateUser(r.Context(), user); err != nil {                                                                                                                                                                     
		statusCode := http.StatusBadRequest                                                                                                                                                                                             
		if err == errors.ErrUserAlreadyExists {                                                                                                                                                                                         
			statusCode = http.StatusConflict                                                                                                                                                                                            
		}                                                                                                                                                                                                                               
		util.WriteErrResponse(w, statusCode, err)                                                                                                                                                                                       
		return                                                                                                                                                                                                                          
	}                                                                                                                                                                                                                                   
																																																										
	util.WriteResponse(w, http.StatusCreated, user)                                                                                                                                                                                     
}                                                                                                                                                                                                                                       
																																																										
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {                                                                                                                                                                     
	email := getEmailFromURL(r)                                                                                                                                                                                                         
																																																										
	user, err := h.service.GetUser(r.Context(), email)                                                                                                                                                                                  
	if err != nil {                                                                                                                                                                                                                     
		statusCode := http.StatusBadRequest                                                                                                                                                                                             
		if err == errors.ErrUserDoesntExists {                                                                                                                                                                                          
			statusCode = http.StatusNotFound                                                                                                                                                                                            
		}                                                                                                                                                                                                                               
		util.WriteErrResponse(w, statusCode, err)                                                                                                                                                                                       
		return                                                                                                                                                                                                                          
	}                                                                                                                                                                                                                                   
																																																										
	util.WriteResponse(w, http.StatusOK, user)                                                                                                                                                                                          
}                                                                                                                                                                                                                                       
																																																										
func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {                                                                                                                                                                   
	users := h.service.ListUsers(r.Context())                                                                                                                                                                                           
	util.WriteResponse(w, http.StatusOK, users)                                                                                                                                                                                         
}                                                                                                                                                                                                                                       
																																																										
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {                                                                                                                                                                  
	email := getEmailFromURL(r)                                                                                                                                                                                                         
																																																										
	b, err := io.ReadAll(r.Body)                                                                                                                                                                                                        
	if err != nil {                                                                                                                                                                                                                     
		util.WriteErrResponse(w, http.StatusBadRequest, err)                                                                                                                                                                            
		return                                                                                                                                                                                                                          
	}                                                                                                                                                                                                                                   
																																																										
	var user model.User                                                                                                                                                                                                                 
	if err := json.Unmarshal(b, &user); err != nil {                                                                                                                                                                                    
		util.WriteErrResponse(w, http.StatusBadRequest, err)                                                                                                                                                                            
		return                                                                                                                                                                                                                          
	}                                                                                                                                                                                                                                   
																																																										
	if err := validate.Struct(user); err != nil {                                                                                                                                                                                                                                                                                                                                                                               
		util.WriteErrResponse(w, http.StatusBadRequest, err)                                                                                                                                                           
		return                                                                                                                                                                                                                          
	}                                                                                                                                                                                                                        
																																																										
	updatedUser, err := h.service.UpdateUser(r.Context(), email, user)                                                                                                                                                                  
	if err != nil {                                                                                                                                                                                                                     
		statusCode := http.StatusBadRequest                                                                                                                                                                                             
		if err == errors.ErrUserDoesntExists {                                                                                                                                                                                          
			statusCode = http.StatusNotFound                                                                                                                                                                                            
		}                                                                                                                                                                                                                               
		util.WriteErrResponse(w, statusCode, err)                                                                                                                                                                                       
		return                                                                                                                                                                                                                          
	}                                                                                                                                                                                                                                   
																																																										
	util.WriteResponse(w, http.StatusOK, updatedUser)                                                                                                                                                                                   
}                                                                                                                                                                                                                                       
																																																										
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {                                                                                                                                                                  
	email := getEmailFromURL(r)                                                                                                                                                                                                         
																																																										
	err := h.service.DeleteUser(r.Context(), email)                                                                                                                                                                                     
	if err != nil {                                                                                                                                                                                                                     
		statusCode := http.StatusBadRequest                                                                                                                                                                                             
		if err == errors.ErrUserDoesntExists {                                                                                                                                                                                          
			statusCode = http.StatusNotFound                                                                                                                                                                                            
		}                                                                                                                                                                                                                               
		util.WriteErrResponse(w, statusCode, err)                                                                                                                                                                                       
		return                                                                                                                                                                                                                          
	}                                                                                                                                                                                                                                   
																																																										
	w.WriteHeader(http.StatusNoContent)                                                                                                                                                                                                 
}