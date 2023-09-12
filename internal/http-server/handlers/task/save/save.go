package save

import (
	"encoding/json"
	"errors"
	"log/slog"
	"micriservice/internal/pkg/api/response"
	"micriservice/internal/storage/mysql_db/table"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

//go:generate mockery --name=SaverTask --output=mocks
type SaverTask interface {
	SaveTask(name string, description string) (int64, error)
}

type Request struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
}

type Response struct {
	IDTask int64 `json:"id_task,omitempty"`
	response.Response
}

// TODO: refactoring this handler. To divide by middleware

func New(log *slog.Logger, saverTask SaverTask) echo.HandlerFunc {
	return func(c echo.Context) error {
		const op = "handlers.url.save.new"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", c.Request().Header.Get("request_id")),
		)

		var req Request

		if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
			error_message := "Failed to decode request body"

			log.Error(error_message, slog.String("err", err.Error()))

			return c.JSON(http.StatusBadRequest, response.Error(error_message))
		}

		log.Info("Requesst body deceoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			error_message := "Failed validation request"

			log.Error(error_message, slog.String("err", err.Error()))

			return c.JSON(http.StatusBadRequest, response.Error(error_message))
		}

		id, err := saverTask.SaveTask(req.Name, req.Description)

		if err != nil {
			if errors.Is(err, table.ErrNameTaskExists) {
				error_message := "The task 'name' already exists"

				log.Info(error_message, slog.String("name", req.Name))

				return c.JSON(http.StatusBadRequest, response.Error(error_message))
			} else {
				error_message := "Failed to add Task"

				log.Error(error_message, slog.String("err", err.Error()))

				return c.JSON(http.StatusBadRequest, response.Error(error_message))
			}
		}

		log.Info("Task add", slog.Int64("id", id))

		return c.JSON(http.StatusOK, Response{
			Response: response.Ok(),
			IDTask:   id,
		})
	}
}
