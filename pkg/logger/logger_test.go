package logger

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"
)

func Test_localHandler(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(h *localHandler) *slog.Logger
		logFn    func(l *slog.Logger)
		contains []string
		absent   []string
	}{
		{
			name:     "basic INFO log has no key names for built-in fields",
			setup:    func(h *localHandler) *slog.Logger { return slog.New(&customHandler{inner: h}) },
			logFn:    func(l *slog.Logger) { l.Info("hello world") },
			contains: []string{"INFO", "hello world"},
			absent:   []string{"time=", "level=", "source=", "msg="},
		},
		{
			name:     "custom attr uses key=value format",
			setup:    func(h *localHandler) *slog.Logger { return slog.New(&customHandler{inner: h}) },
			logFn:    func(l *slog.Logger) { l.Info("msg", slog.String("env", "dev")) },
			contains: []string{"env=dev"},
		},
		{
			name:     "attr value with spaces is quoted",
			setup:    func(h *localHandler) *slog.Logger { return slog.New(&customHandler{inner: h}) },
			logFn:    func(l *slog.Logger) { l.Info("msg", slog.String("k", "v a")) },
			contains: []string{`k="v a"`},
		},
		{
			name:     "ERROR log contains stacktrace",
			setup:    func(h *localHandler) *slog.Logger { return slog.New(&customHandler{inner: h}) },
			logFn:    func(l *slog.Logger) { l.Error("boom") },
			contains: []string{"ERROR", "boom", "stacktrace="},
		},
		{
			name: "DEBUG log is filtered when level is INFO",
			setup: func(h *localHandler) *slog.Logger {
				h.level = slog.LevelInfo
				return slog.New(&customHandler{inner: h})
			},
			logFn:  func(l *slog.Logger) { l.Debug("hidden") },
			absent: []string{"hidden"},
		},
		{
			name: "With() pre-attrs appear in output",
			setup: func(h *localHandler) *slog.Logger {
				return slog.New(&customHandler{inner: h}).With(slog.String("app", "svc"))
			},
			logFn:    func(l *slog.Logger) { l.Info("msg") },
			contains: []string{"app=svc"},
		},
		{
			name: "WithGroup() prefixes attr keys",
			setup: func(h *localHandler) *slog.Logger {
				return slog.New(&customHandler{inner: h}).WithGroup("grp")
			},
			logFn:    func(l *slog.Logger) { l.Info("msg", slog.String("k", "v")) },
			contains: []string{"grp.k=v"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			h := newLocalHandler(&buf, slog.LevelDebug)
			logger := tt.setup(h)
			tt.logFn(logger)
			out := buf.String()
			t.Log(out)

			for _, want := range tt.contains {
				if !strings.Contains(out, want) {
					t.Errorf("output %q does not contain %q", out, want)
				}
			}
			for _, absent := range tt.absent {
				if strings.Contains(out, absent) {
					t.Errorf("output %q should not contain %q", out, absent)
				}
			}
		})
	}
}
