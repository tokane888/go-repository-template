package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"
)

type Config struct {
	Level      string // debug, info, warn, error
	Format     string // local(見やすさ重視), cloud(CloudWatch等で解析可能であることを重視)
	Env        string // 環境名(cloudでのみログ出力)
	AppName    string // アプリ名(cloudでのみログ出力)
	AppVersion string // アプリのバージョン(cloudでのみログ出力)
}

type customHandler struct {
	inner slog.Handler
}

func (h *customHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.inner.Enabled(ctx, level)
}

func (h *customHandler) Handle(ctx context.Context, r slog.Record) error {
	if r.Level >= slog.LevelError {
		buf := make([]byte, 4096)
		n := runtime.Stack(buf, false)
		r.AddAttrs(slog.String("stacktrace", string(buf[:n])))
	}
	return h.inner.Handle(ctx, r)
}

func (h *customHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &customHandler{inner: h.inner.WithAttrs(attrs)}
}

func (h *customHandler) WithGroup(name string) slog.Handler {
	return &customHandler{inner: h.inner.WithGroup(name)}
}

func localTimeReplacer(_ []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey {
		return slog.String(slog.TimeKey, a.Value.Time().In(time.Local).Format("2006-01-02T15:04:05.000Z07:00"))
	}
	return a
}

func cloudTimeReplacer(_ []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey {
		return slog.String(slog.TimeKey, a.Value.Time().UTC().Format(time.RFC3339Nano))
	}
	return a
}

func NewLogger(cfg Config) *slog.Logger {
	var level slog.Level
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		fmt.Fprintf(os.Stderr, "invalid LOG_LEVEL %q, fallback to 'info'\n", cfg.Level)
		level = slog.LevelInfo
	}
	opts := &slog.HandlerOptions{Level: level, AddSource: true}

	var inner slog.Handler
	switch cfg.Format {
	case "local":
		// local環境では読みやすさ重視
		// (非構造化ログ、ローカルタイムゾーン、ミリ秒精度)
		opts.ReplaceAttr = localTimeReplacer
		inner = slog.NewTextHandler(os.Stderr, opts)
	case "cloud":
		// cloud環境ではcloud watch等で読まれる前提で解析重視
		// (構造化ログ、UTC、ナノ秒精度)
		opts.ReplaceAttr = cloudTimeReplacer
		inner = slog.NewJSONHandler(os.Stderr, opts)
	default:
		// LOG_FORMATが不正の場合、cloud向けフォーマットで出力
		fmt.Fprintf(os.Stderr, "invalid LOG_FORMAT %q, fallback to 'cloud'\n", cfg.Format)
		opts.ReplaceAttr = cloudTimeReplacer
		inner = slog.NewJSONHandler(os.Stderr, opts)
	}

	// Errorレベルにスタックトレースを付与するカスタムHandler
	handler := &customHandler{inner: inner}
	logger := slog.New(handler)

	// cloud向けはフィールド追加
	if cfg.Format == "cloud" {
		logger = logger.With(
			slog.String("app", cfg.AppName),
			slog.String("env", cfg.Env),
			slog.String("ver", cfg.AppVersion),
		)
	}

	return logger
}
