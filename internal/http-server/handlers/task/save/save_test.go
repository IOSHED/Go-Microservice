package save_test

import (
	"log/slog"
	"micriservice/internal/config"
	"micriservice/internal/http-server/handlers/task/save"
	"micriservice/internal/http-server/handlers/task/save/mocks"
	"micriservice/internal/logger"
	http_server "micriservice/internal/pkg/http-server/echo"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Сall at the beginning of the test.
func setupTest(json_request string) (echo.Context, *slog.Logger, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodPost, "/test/create", strings.NewReader(json_request))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	cfg := &config.Config{
		Env: "local",
		HttpServer: config.HttpServer{
			Address:     "localhost:8080",
			Timeout:     time.Duration(4 * time.Second),
			IdleTimeout: time.Duration(60 * time.Second),
		},
	}

	server := http_server.New(cfg)

	context := server.Echo.NewContext(req, rec)

	log := logger.SetupLogger(cfg.Env)

	return context, log, rec
}

func TestHandlerBadRequest(t *testing.T) {

	context, log, rec := setupTest(`{"invalid": "json"}`)

	taskSaverMock := mocks.NewSaverTask(t)

	handler := save.New(log, taskSaverMock)
	handler(context)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandlerValidationFailed(t *testing.T) {

	context, log, rec := setupTest(`{"name": ""}`)

	taskSaverMock := mocks.NewSaverTask(t)

	handler := save.New(log, taskSaverMock)
	handler(context)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// TODO: error
func TestHandlerSaveSuccess(t *testing.T) {

	context, log, rec := setupTest(`{"name": "todo1"}`)

	taskSaverMock := mocks.NewSaverTask(t)

	handler := save.New(log, taskSaverMock)
	handler(context)

	assert.Equal(t, http.StatusOK, rec.Code)
}
