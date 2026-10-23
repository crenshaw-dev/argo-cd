package exec

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
	"unicode"

	"github.com/argoproj/argo-cd/gitops-engine/v3/pkg/utils/tracing"
	"github.com/sirupsen/logrus"

	"github.com/argoproj/argo-cd/v3/util/log"
	"github.com/argoproj/argo-cd/v3/util/rand"
)

var (
	timeout                       time.Duration
	fatalTimeout                  time.Duration
	processGroupIsolationDisabled atomic.Bool
	Unredacted                    = Redact(nil)
)

// DisableProcessGroupIsolation keeps child processes in the caller's process
// group. The argocd CLI uses this so Ctrl-C still reaches helm and kustomize.
func DisableProcessGroupIsolation() {
	processGroupIsolationDisabled.Store(true)
}

// Signal sends sig to cmd. A command started in its own process group is
// signaled as a group, so children receive it too.
func Signal(cmd *exec.Cmd, sig syscall.Signal) error {
	return signalCmd(cmd, sig)
}

var execIDs sync.Map

func rememberExecID(cmd *exec.Cmd, id string) {
	execIDs.Store(cmd, id)
}

func forgetExecID(cmd *exec.Cmd) {
	execIDs.Delete(cmd)
}

func lookupExecID(cmd *exec.Cmd) string {
	id, ok := execIDs.Load(cmd)
	if !ok {
		return ""
	}
	return id.(string)
}

// terminationCause reports why ctx ended. A cause other than cancellation or a
// deadline is a shutdown: the repo server uses its own error for that.
func TerminationCause(ctx context.Context) string {
	err := context.Cause(ctx)
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return "timeout"
	case err != nil && !errors.Is(err, context.Canceled):
		return "shutdown"
	default:
		return "cancellation"
	}
}

// LogTermination records why this package signaled a command.
func LogTermination(ctx context.Context, cmd *exec.Cmd, cause string) {
	pid := 0
	if cmd != nil && cmd.Process != nil {
		pid = cmd.Process.Pid
	}
	contextCause := ""
	if ctx != nil && context.Cause(ctx) != nil {
		contextCause = context.Cause(ctx).Error()
	}
	logrus.WithFields(logrus.Fields{
		"execID":       lookupExecID(cmd),
		"pid":          pid,
		"cause":        cause,
		"contextCause": contextCause,
	}).Info("terminated command")
}

type ExecRunOpts struct {
	// Redactor redacts tokens from the output
	Redactor func(text string) string
	// TimeoutBehavior configures what to do in case of timeout
	TimeoutBehavior TimeoutBehavior
	// SkipErrorLogging determines whether to skip logging of execution errors (rc > 0)
	SkipErrorLogging bool
	// CaptureStderr determines whether to capture stderr in addition to stdout
	CaptureStderr bool
	// StderrWriter receives stderr without adding it to the returned output
	StderrWriter io.Writer
}

func init() {
	initTimeout()
}

func initTimeout() {
	var err error
	timeout, err = time.ParseDuration(os.Getenv("ARGOCD_EXEC_TIMEOUT"))
	if err != nil {
		timeout = 90 * time.Second
	}
	fatalTimeout, err = time.ParseDuration(os.Getenv("ARGOCD_EXEC_FATAL_TIMEOUT"))
	if err != nil {
		fatalTimeout = 10 * time.Second
	}
}

func Run(cmd *exec.Cmd) (string, error) {
	return RunWithRedactor(cmd, nil)
}

func RunWithRedactor(cmd *exec.Cmd, redactor func(text string) string) (string, error) {
	opts := ExecRunOpts{Redactor: redactor}
	return RunWithExecRunOpts(cmd, opts)
}

func RunWithExecRunOpts(cmd *exec.Cmd, opts ExecRunOpts) (string, error) {
	cmdOpts := CmdOpts{
		Timeout:          timeout,
		FatalTimeout:     fatalTimeout,
		Redactor:         opts.Redactor,
		TimeoutBehavior:  opts.TimeoutBehavior,
		SkipErrorLogging: opts.SkipErrorLogging,
		CaptureStderr:    opts.CaptureStderr,
		StderrWriter:     opts.StderrWriter,
	}
	span := tracing.NewLoggingTracer(log.NewLogrusLogger(log.NewWithCurrentConfig())).StartSpan(fmt.Sprintf("exec %v", cmd.Args[0]))
	span.SetBaggageItem("dir", cmd.Dir)
	if cmdOpts.Redactor != nil {
		span.SetBaggageItem("args", opts.Redactor(fmt.Sprintf("%v", cmd.Args)))
	} else {
		span.SetBaggageItem("args", fmt.Sprintf("%v", cmd.Args))
	}
	defer span.Finish()
	return RunCommandExt(cmd, cmdOpts)
}

// GetCommandArgsToLog represents the given command in a way that we can copy-and-paste into a terminal
func GetCommandArgsToLog(cmd *exec.Cmd) string {
	var argsToLog []string
	for _, arg := range cmd.Args {
		if arg == "" {
			argsToLog = append(argsToLog, `""`)
			continue
		}

		containsSpace := false
		for _, r := range arg {
			if unicode.IsSpace(r) {
				containsSpace = true
				break
			}
		}
		if containsSpace {
			// add quotes and escape any internal quotes
			argsToLog = append(argsToLog, strconv.Quote(arg))
		} else {
			argsToLog = append(argsToLog, arg)
		}
	}
	args := strings.Join(argsToLog, " ")
	return args
}

type CmdError struct {
	Args   string
	Stderr string
	Cause  error
}

func (ce *CmdError) Error() string {
	res := fmt.Sprintf("`%v` failed %v", ce.Args, ce.Cause)
	if ce.Stderr != "" {
		res = fmt.Sprintf("%s: %s", res, ce.Stderr)
	}
	return res
}

func (ce *CmdError) String() string {
	return ce.Error()
}

func newCmdError(args string, cause error, stderr string) *CmdError {
	return &CmdError{Args: args, Stderr: stderr, Cause: cause}
}

// TimeoutBehavior defines behavior for when the command takes longer than the passed in timeout to exit
// By default, SIGKILL is sent to the process and it is not waited upon
type TimeoutBehavior struct {
	// Signal determines the signal to send to the process
	Signal syscall.Signal
	// ShouldWait determines whether to wait for the command to exit once timeout is reached
	ShouldWait bool
}

type CmdOpts struct {
	// Timeout determines how long to wait for the command to exit
	Timeout time.Duration
	// FatalTimeout is the amount of additional time to wait after Timeout before fatal SIGKILL
	FatalTimeout time.Duration
	// Redactor redacts tokens from the output
	Redactor func(text string) string
	// TimeoutBehavior configures what to do in case of timeout
	TimeoutBehavior TimeoutBehavior
	// SkipErrorLogging defines whether to skip logging of execution errors (rc > 0)
	SkipErrorLogging bool
	// CaptureStderr defines whether to capture stderr in addition to stdout
	CaptureStderr bool
	// StderrWriter receives stderr without adding it to the returned output
	StderrWriter io.Writer
}

var DefaultCmdOpts = CmdOpts{
	Timeout:          time.Duration(0),
	FatalTimeout:     time.Duration(0),
	Redactor:         Unredacted,
	TimeoutBehavior:  TimeoutBehavior{syscall.SIGKILL, false},
	SkipErrorLogging: false,
	CaptureStderr:    false,
}

func Redact(items []string) func(text string) string {
	return func(text string) string {
		for _, item := range items {
			text = strings.ReplaceAll(text, item, "******")
		}
		return text
	}
}

// RunCommandExt is a convenience function to run/log a command and return/log stderr in an error upon
// failure.
func RunCommandExt(cmd *exec.Cmd, opts CmdOpts) (string, error) {
	execId, err := rand.RandHex(5)
	if err != nil {
		return "", err
	}
	logCtx := logrus.WithFields(logrus.Fields{"execID": execId})
	rememberExecID(cmd, execId)
	defer forgetExecID(cmd)

	redactor := DefaultCmdOpts.Redactor
	if opts.Redactor != nil {
		redactor = opts.Redactor
	}

	// log in a way we can copy-and-paste into a terminal
	args := strings.Join(cmd.Args, " ")
	logCtx.WithFields(logrus.Fields{"dir": cmd.Dir}).Info(redactor(args))

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if opts.StderrWriter != nil {
		cmd.Stderr = io.MultiWriter(&stderr, opts.StderrWriter)
	}

	// Start timers for timeout
	timeout := DefaultCmdOpts.Timeout
	fatalTimeout := DefaultCmdOpts.FatalTimeout

	if opts.Timeout != time.Duration(0) {
		timeout = opts.Timeout
	}

	if opts.FatalTimeout != time.Duration(0) {
		fatalTimeout = opts.FatalTimeout
	}
	// Bound the wait for inherited pipes. A grandchild that outlives the command
	// would otherwise keep Wait blocked until that grandchild exits. A caller that
	// already set WaitDelay, such as the git client, keeps its own value.
	// WaitDelay has to be set before Start: for a command created with a
	// cancelable context, Start launches a goroutine that reads it.
	if cmd.WaitDelay == 0 && fatalTimeout != 0 {
		cmd.WaitDelay = fatalTimeout
	}

	start := time.Now()
	setProcessGroup(cmd)
	err = cmd.Start()
	if err != nil {
		return "", err
	}

	timeoutBehavior := DefaultCmdOpts.TimeoutBehavior
	if opts.TimeoutBehavior.Signal != syscall.Signal(0) {
		timeoutBehavior = opts.TimeoutBehavior
	}

	// The timeout is a context deadline. When it fires, the command is signaled
	// and, if it should be given a grace period, SIGKILL follows fatalTimeout
	// later. Wait itself is the only thing the caller blocks on: the deadline
	// goroutine exits when this function cancels the context after Wait returns.
	var timedOut atomic.Bool
	var fatalHit atomic.Bool
	stopFatal := make(chan struct{})
	var timeoutWG sync.WaitGroup
	timeoutCtx, cancelTimeout := context.WithCancel(context.Background())
	if timeout > 0 {
		timeoutCtx, cancelTimeout = context.WithTimeout(context.Background(), timeout)
	}
	defer cancelTimeout()
	if timeout > 0 {
		timeoutWG.Add(1)
		go func() {
			defer timeoutWG.Done()
			<-timeoutCtx.Done()
			if !errors.Is(timeoutCtx.Err(), context.DeadlineExceeded) {
				return
			}
			timedOut.Store(true)
			_ = signalCmd(cmd, timeoutBehavior.Signal)
			if !timeoutBehavior.ShouldWait || fatalTimeout == 0 {
				return
			}
			timer := time.NewTimer(fatalTimeout)
			select {
			case <-timer.C:
				fatalHit.Store(true)
				_ = signalCmd(cmd, syscall.SIGKILL)
			case <-stopFatal:
				timer.Stop()
			}
		}()
	}

	err = cmd.Wait()
	cancelTimeout()
	close(stopFatal)
	timeoutWG.Wait()

	output := stdout.String()
	if opts.CaptureStderr {
		output += stderr.String()
	}
	logCtx.WithFields(logrus.Fields{"duration": time.Since(start)}).Debug(redactor(output))
	trimmed := strings.TrimSuffix(output, "\n")

	if fatalHit.Load() {
		LogTermination(timeoutCtx, cmd, "timeout")
		err = newCmdError(redactor(args), fmt.Errorf("fatal timeout after %v", timeout+fatalTimeout), "")
		logCtx.Error(err.Error())
		return trimmed, err
	}
	if timedOut.Load() {
		LogTermination(timeoutCtx, cmd, "timeout")
		err = newCmdError(redactor(args), fmt.Errorf("timeout after %v", timeout), "")
		logCtx.Error(err.Error())
		return trimmed, err
	}
	if err != nil {
		err = newCmdError(redactor(args), errors.New(redactor(err.Error())), strings.TrimSpace(redactor(stderr.String())))
		if !opts.SkipErrorLogging {
			logCtx.Error(err.Error())
		}
		return trimmed, err
	}
	return trimmed, nil
}

func RunCommand(name string, opts CmdOpts, arg ...string) (string, error) {
	return RunCommandExt(exec.CommandContext(context.Background(), name, arg...), opts)
}
