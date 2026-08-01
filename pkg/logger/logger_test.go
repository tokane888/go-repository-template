package logger

import (
	"bytes"
	"io"
	"log/slog"
	"os"
	"strings"
	"testing"
)

func Test_NewLogger(t *testing.T) {
	tests := []struct {
		name     string
		cfg      Config
		logFn    func(l *slog.Logger)
		contains []string
		absent   []string
	}{
		{
			name:     "local format: basic INFO log has no key names for built-in fields",
			cfg:      Config{Level: "debug", Format: "local"},
			logFn:    func(l *slog.Logger) { l.Info("hello world") },
			contains: []string{"INFO", "hello world"},
			absent:   []string{"time=", "level=", "source=", "msg="},
		},
		{
			name:     "local format: custom attr uses key=value format",
			cfg:      Config{Level: "debug", Format: "local"},
			logFn:    func(l *slog.Logger) { l.Info("msg", slog.String("env", "dev")) },
			contains: []string{"env=dev"},
		},
		{
			name:     "local format: attr value with spaces is quoted",
			cfg:      Config{Level: "debug", Format: "local"},
			logFn:    func(l *slog.Logger) { l.Info("msg", slog.String("k", "v a")) },
			contains: []string{`k="v a"`},
		},
		{
			name:     "local format: ERROR log contains stacktrace",
			cfg:      Config{Level: "debug", Format: "local"},
			logFn:    func(l *slog.Logger) { l.Error("boom") },
			contains: []string{"ERROR", "boom", "stacktrace="},
			absent:   []string{"/pkg/logger/"},
		},
		{
			name:   "local format: DEBUG log is filtered when level is INFO",
			cfg:    Config{Level: "info", Format: "local"},
			logFn:  func(l *slog.Logger) { l.Debug("hidden") },
			absent: []string{"hidden"},
		},
		{
			name:     "local format: With() pre-attrs appear in output",
			cfg:      Config{Level: "debug", Format: "local"},
			logFn:    func(l *slog.Logger) { l.With(slog.String("app", "svc")).Info("msg") },
			contains: []string{"app=svc"},
		},
		{
			name:     "local format: WithGroup() prefixes attr keys",
			cfg:      Config{Level: "debug", Format: "local"},
			logFn:    func(l *slog.Logger) { l.WithGroup("grp").Info("msg", slog.String("k", "v")) },
			contains: []string{"grp.k=v"},
		},
		{
			name:     "cloud format: emits JSON with env/ver fields",
			cfg:      Config{Level: "info", Format: "cloud", Env: "prod", AppVersion: "1.2.3"},
			logFn:    func(l *slog.Logger) { l.Info("hello") },
			contains: []string{`"msg":"hello"`, `"env":"prod"`, `"ver":"1.2.3"`},
		},
		{
			name:     "invalid format falls back to cloud JSON with warning",
			cfg:      Config{Level: "info", Format: "bogus"},
			logFn:    func(l *slog.Logger) { l.Info("hello") },
			contains: []string{"invalid LOG_FORMAT", `"msg":"hello"`},
		},
		{
			name:     "invalid level falls back to info with warning",
			cfg:      Config{Level: "bogus", Format: "cloud"},
			logFn:    func(l *slog.Logger) { l.Debug("hidden"); l.Info("shown") },
			contains: []string{"invalid LOG_LEVEL", `"msg":"shown"`},
			absent:   []string{"hidden"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}
			orig := os.Stderr
			os.Stderr = w

			logger := NewLogger(tt.cfg)
			tt.logFn(logger)

			os.Stderr = orig
			if err := w.Close(); err != nil {
				t.Fatal(err)
			}

			var buf bytes.Buffer
			if _, err := io.Copy(&buf, r); err != nil {
				t.Fatal(err)
			}
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
