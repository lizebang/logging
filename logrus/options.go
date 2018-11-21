/*
 * Revision History:
 *     Initial: 2018/08/12        Feng Yifei
 */

package logrus

import (
	log "github.com/sirupsen/logrus"

	"github.com/TechCatsLab/logging"
	"github.com/TechCatsLab/logging/logrus/hooks"
)

// OptSetDebugLevel set logging level as Debug.
func OptSetDebugLevel(l logging.Logger) error {
	if lg, ok := l.(*logger); ok {
		lg.lg.SetLevel(log.DebugLevel)
	}

	return nil
}

// OptSetInfoLevel set logging level as Info.
func OptSetInfoLevel(l logging.Logger) error {
	if lg, ok := l.(*logger); ok {
		lg.lg.SetLevel(log.InfoLevel)
	}

	return nil
}

// OptSetWarnLevel set logging level as Warn.
func OptSetWarnLevel(l logging.Logger) error {
	if lg, ok := l.(*logger); ok {
		lg.lg.SetLevel(log.WarnLevel)
	}

	return nil
}

// OptSetErrorLevel set logging level as Error.
func OptSetErrorLevel(l logging.Logger) error {
	if lg, ok := l.(*logger); ok {
		lg.lg.SetLevel(log.ErrorLevel)
	}

	return nil
}

// OptShowFileLine set log file line.
func OptShowFileLine(l logging.Logger) error {
	if lg, ok := l.(*logger); ok {
		lg.lg.AddHook(hooks.FileLineHook{})
	}

	return nil
}
