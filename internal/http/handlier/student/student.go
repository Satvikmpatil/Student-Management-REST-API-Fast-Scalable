package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Satvikmpatil/Student-Management-REST-API-Fast-Scalable/internal/Utils/response"
	"github.com/Satvikmpatil/Student-Management-REST-API-Fast-Scalable/internal/types"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		
		err:=json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err,io.EOF){
			response.WriteJson(w,http.StatusBadRequest,response.GerneralError(err))
			return 
		}

		if err != nil{
			response.WriteJson(w,http.StatusBadRequest,response.GerneralError(err))
			return 
		}
		slog.Info("Creating a student")

		//request valid

		if err :=validator.New().Struct(student); err != nil{
			validateError := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest,response.ValidationErrors(validateError))
		}


		response.WriteJson(w, http.StatusCreated, map[string]string{"success":"OK"})

		w.Write([]byte("WELCOME STUDENT"))
	}
}
