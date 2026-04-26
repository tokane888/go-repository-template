package logger

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Config struct {
	Level      string // debug, info, warn, error
	Format     string // local(見やすさ重視), cloud(CloudWatch等で解析可能であることを重視)
	Env        string // 環境名(cloudでのみログ出力)
	AppName    string // アプリ名(cloudでのみログ出力)
	AppVersion string // アプリのバージョン(cloudでのみログ出力)
}

// localHandler outputs logs as space-separated values without key names for
// built-in fields (time, level, source, msg). Custom attributes use key=value.
type localHandler struct {
	mu       sync.Mutex
	out      io.Writer
	level    slog.Level
	wd       string // working directory for relative source paths
	preAttrs []slog.Attr
	groups   []string
}

func newLocalHandler(out io.Writer, level slog.Level) *localHandler {
	wd, _ := os.Getwd()
	return &localHandler{out: out, level: level, wd: wd}
}

func (h *localHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

// Handle output log with simple format
// e.g. "2026-04-26T14:47:49.654+09:00 INFO cmd/sample/main.go:23 hello"
func (h *localHandler) Handle(_ context.Context, r slog.Record) error {
	var buf bytes.Buffer

	if !r.Time.IsZero() {
		buf.WriteString(r.Time.In(time.Local).Format("2006-01-02T15:04:05.000Z07:00"))
		buf.WriteByte(' ')
	}

	buf.WriteString(r.Level.String())

	if r.PC != 0 {
		frames := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := frames.Next()
		src := f.File
		if h.wd != "" {
			if rel, err := filepath.Rel(h.wd, f.File); err == nil {
				src = rel
			}
		}
		fmt.Fprintf(&buf, " %s:%d", src, f.Line)
	}

	buf.WriteByte(' ')
	buf.WriteString(r.Message)

	for _, a := range h.preAttrs {
		h.appendAttr(&buf, a)
	}
	r.Attrs(func(a slog.Attr) bool {
		h.appendAttr(&buf, a)
		return true
	})

	buf.WriteByte('\n')

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf.Bytes())
	return err
}

func (h *localHandler) appendAttr(buf *bytes.Buffer, a slog.Attr) {
	a.Value = a.Value.Resolve()
	if a.Equal(slog.Attr{}) {
		return
	}
	if a.Value.Kind() == slog.KindGroup {
		for _, ga := range a.Value.Group() {
			if a.Key != "" {
				ga.Key = a.Key + "." + ga.Key
			}
			h.appendAttr(buf, ga)
		}
		return
	}
	key := a.Key
	if len(h.groups) > 0 {
		key = strings.Join(h.groups, ".") + "." + key
	}
	val := a.Value.String()
	if strings.ContainsAny(val, " \t=") {
		fmt.Fprintf(buf, " %s=%q", key, val)
	} else {
		fmt.Fprintf(buf, " %s=%s", key, val)
	}
}

func (h *localHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newAttrs := make([]slog.Attr, len(h.preAttrs)+len(attrs))
	copy(newAttrs, h.preAttrs)
	copy(newAttrs[len(h.preAttrs):], attrs)
	return &localHandler{
		out:      h.out,
		level:    h.level,
		wd:       h.wd,
		preAttrs: newAttrs,
		groups:   h.groups,
	}
}

func (h *localHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	newGroups := make([]string, len(h.groups)+1)
	copy(newGroups, h.groups)
	newGroups[len(h.groups)] = name
	return &localHandler{
		out:      h.out,
		level:    h.level,
		wd:       h.wd,
		preAttrs: h.preAttrs,
		groups:   newGroups,
	}
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

	var inner slog.Handler
	switch cfg.Format {
	case "local":
		// local環境では読みやすさ重視
		// (非構造化ログ、ローカルタイムゾーン、ミリ秒精度)
		inner = newLocalHandler(os.Stderr, level)
	case "cloud":
		// cloud環境ではcloud watch等で読まれる前提で解析重視
		// (構造化ログ、UTC、ナノ秒精度)
		opts := &slog.HandlerOptions{Level: level, AddSource: true, ReplaceAttr: cloudTimeReplacer}
		inner = slog.NewJSONHandler(os.Stderr, opts)
	default:
		// LOG_FORMATが不正の場合、cloud向けフォーマットで出力
		fmt.Fprintf(os.Stderr, "invalid LOG_FORMAT %q, fallback to 'cloud'\n", cfg.Format)
		opts := &slog.HandlerOptions{Level: level, AddSource: true, ReplaceAttr: cloudTimeReplacer}
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
