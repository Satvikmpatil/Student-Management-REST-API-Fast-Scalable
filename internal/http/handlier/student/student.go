package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Satvikmpatil/Student-Management-REST-API-Fast-Scalable/internal/Utils/response"
	"github.com/Satvikmpatil/Student-Management-REST-API-Fast-Scalable/internal/storage"
	"github.com/Satvikmpatil/Student-Management-REST-API-Fast-Scalable/internal/types"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GerneralError(err))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GerneralError(err))
			return
		}
		slog.Info("Creating a student")

		//request valid

		if err := validator.New().Struct(student); err != nil {
			validateError := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationErrors(validateError))
			return
		}

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GerneralError(err))
			return
		}
		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("Getting a Student",slog.String("id",id))
		studentId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GerneralError(err))
			return
		}
		student, err := storage.GetStudentById(studentId)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GerneralError(err))
			return
		}
		response.WriteJson(w, http.StatusOK, student)
	}
}