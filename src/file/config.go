package file

import (
    "errors"
)

var NotImplementErr = errors.New("Not Implement.")

type Configure struct {
    filePath string
}

type PidFile struct {
    pid uintptr
    filePath string
}

type LogFile struct {
    filePath string
}

type FinishFile struct {
    filePath string
}

func NewConfigure(filePath string) *Configure {
    return &Configure{filePath}
}

func (*Configure) LoadConfigure() error {
    // return NotImplementErr
    return nil
}

func (*Configure) PidFileName() string {
    return ""
}

func (*Configure) LogFileName() string {
    return ""
}

func (*Configure) LogMaxSize() int {
    return 0
}

func (*Configure) LogMaxIndex() int {
    return 0
}

func (*Configure) LogLevel() string {
    return ""
}

func (*Configure) FinishFile() string {
    return ""
}

func (*Configure) TaskCount() int {
    return 0
}

func NewPidFile(filePath string) *PidFile {
    return &PidFile{0, filePath}
}

func (p *PidFile) Init(pid uintptr) error {
    // TODO Examine the file is empty or not. If not, test whether the pid is running.
    // If running, then throw error. If not running, write the pid to the file.
    return nil
}

func NewLogFile(filePath string) *LogFile {
    return &LogFile{filePath}
}

func (l *LogFile) Init(maxSize, maxIndex int, logLevel string) error {
    // return NotImplementErr
    return nil
}

func NewFinishFile(filePath string) *FinishFile {
    return &FinishFile{filePath}
}

func (*FinishFile) Init() error {
    return nil
}
