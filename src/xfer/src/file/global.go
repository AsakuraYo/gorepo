package file

type Global struct {
    workPath string
}

var g *Global

func NewGlobal() *Global {
    if g == nil {
        g = &Global{}
    }
    return g
}

func (g *Global) WorkPath() string {
    return g.workPath
}

func (g *Global) SetWorkPath(workPath string) {
    g.workPath = workPath
}

func (g *Global) ThreadCount() int {
    return 0
}

func (g *Global) SetThreadCount(threadCount int) {
}
