package file

type PidFile struct {
    pid unit
}

func NewPidFile(filePath string) *PidFile {
    return nil
}

func (p *PidFile) Init(pid unit) error {
    return nil
}

func (p *PidFile) Remove() {
}
