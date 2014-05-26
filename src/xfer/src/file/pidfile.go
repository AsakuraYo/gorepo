package file

type PidFile struct {
    pid uintptr
}

func NewPidFile(filePath string) *PidFile {
    return nil
}

func (p *PidFile) Init(pid uintptr) error {
    return nil
}

func (p *PidFile) Remove() {
}
