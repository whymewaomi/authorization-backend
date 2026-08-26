package core_logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"
)

func NewLogger() (*slog.Logger, *os.File, error) {
	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, nil, fmt.Errorf("create logs directory: %w", err)
	}

	cwd, _ := os.Getwd()
	fmt.Println("WORKING DIRECTORY:", cwd)

	fileName := fmt.Sprintf(
		"logs/app-%s.log",
		time.Now().Format("2006-01-02_15-04-05.000000000"),
	)

	file, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_EXCL,
		0644,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("open log file: %w", err)
	}

	writer := io.MultiWriter(os.Stdout, file)

	handler := slog.NewJSONHandler(writer, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return slog.New(handler), file, nil
}
