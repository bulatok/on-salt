package salt

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Logger struct {
	bot     *bot
	level   Level
	options *Options
	buf     []byte
	mu      *sync.Mutex
}

// NewProduction will log only InfoLevel and higher Level entries
func NewProduction(token string, options *Options, dsts ...Dst) (*Logger, error) {
	return newLogger(token, options, InfoLevel, dsts...)
}

// NewDev will log on any Level
func NewDev(token string, options *Options, dsts ...Dst) (*Logger, error) {
	return newLogger(token, options, TraceLevel, dsts...)
}

func newLogger(token string, options *Options, level Level, dsts ...Dst) (*Logger, error) {
	for _, dst := range dsts {
		if !dst.Valid() {
			return nil, fmt.Errorf("%s is not correct identifier", dst.ID())
		}
	}
	return &Logger{
		options: options,
		level:   level,
		buf:     []byte{},
		mu:      new(sync.Mutex),
		bot: &bot{
			Token: token,
			Dst:   dsts,
		},
	}, nil
}

func (l *Logger) prepareMsg(msg string, level Level) string {
	// LEVEL [TIME]: MSG
	defer l.mu.Unlock()
	l.mu.Lock()
	l.buf = l.buf[:0]
	if l.options.WithoutLevels {
		l.buf = append(l.buf, msg...)
		return string(l.buf)
	}

	l.buf = append(l.buf, level.String()...)
	l.buf = append(l.buf, ' ')
	l.buf = append(l.buf, '[')
	now := time.Now()
	if !l.options.TimeUnix {
		l.buf = append(l.buf, now.Format("2006-01-02T15:04:05.999-MST")...)
	} else {
		u := strconv.FormatInt(now.Unix(), 10)
		l.buf = append(l.buf, u...)
	}
	l.buf = append(l.buf, "]: "...)
	l.buf = append(l.buf, msg...)
	return string(l.buf)
}

func (l *Logger) Error(msg string) error {
	return l.bot.send(l.prepareMsg(msg, ErrorLevel))
}

func (l *Logger) Info(msg string) error {
	return l.bot.send(l.prepareMsg(msg, InfoLevel))
}

func (l *Logger) Debug(msg string) error {
	if l.level >= InfoLevel {
		return nil
	}
	return l.bot.send(l.prepareMsg(msg, DebugLevel))
}

func (l *Logger) Trace(msg string) error {
	if l.level >= InfoLevel {
		return nil
	}
	return l.bot.send(l.prepareMsg(msg, TraceLevel))
}

func (l *Logger) Warn(msg string) error {
	return l.bot.send(l.prepareMsg(msg, WarnLevel))
}

func (l *Logger) Fatal(msg string) error {
	defer os.Exit(1)
	return l.bot.send(l.prepareMsg(msg, FatalLevel))
}

func (l *Logger) Panic(msg string) error {
	defer func() {
		panic(msg)
	}()
	return l.bot.send(l.prepareMsg(msg, PanicLevel))
}
