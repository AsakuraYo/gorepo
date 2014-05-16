package main

import (
    "os"
    "fmt"
    "errors"
    "bufio"
    "io"
)

// ==Context start==
var (
    InitailizeError = errors.New("Initialize error")
)

type Context struct {
    sourceDir           string
    sourceFile          string
    checkDir            string
    checkFile           string
    targetDir           string
    targetFile          string
    separator           string
    requestUserName     string
    requestUserNumber   string
    processField        string
    cryptoMethod        string
    encryptKey          string
    encryptVersion      int
    encryptMethodIndex  int
    ifReplaceSource     bool
    isDecrypt           bool
    processFields       []int
}

func NewContext() *Context {
    return &Context{}
}

func (c *Context) Init(args []string) error {
    c.sourceDir     = args[1]
    c.sourceFile    = args[2]
    c.checkDir      = args[3]
    c.checkFile     = args[4]
    c.targetDir     = args[5]
    c.targetFile    = args[6]
    return nil
}

func (c *Context) SourceFile() string {
    return c.sourceDir + "/" + c.sourceFile
}

func (c *Context) CheckFile() string {
    return c.checkDir + "/" + c.checkFile
}

func (c *Context) TargetFile() string {
    return c.targetDir + "/" + c.targetFile
}
// ==Context end==

// ==CheckFile start==
type CheckFile struct {
}

func NewCheckFile(filename string) *CheckFile {
    return &CheckFile{}
}

func (cf *CheckFile) LoadInfo() error {
    return nil
}
// ==CheckFile end==

// ==Processor start==
type ProcessFunc func(string) (string, int)
func lineProcess(source *os.File, target *os.File, handler ProcessFunc) int {
    totalLines := 0
    reader := bufio.NewReader(source)
    for {
        inLine, err := reader.ReadString('\n')
        if err == io.EOF {
            break
        } else {
            outLine, lines := handler(inLine)
            totalLines += lines
            if _, err := target.WriteString(outLine); err != nil {
                fmt.Println(err)
                os.Exit(1)
            }
        }
    }
    return totalLines
}
// ==Processor end==


const PACKAGE = "file_cryptor"
const VERSION = "V14.00.001"

func usage() {
    fmt.Println("Version:", VERSION)
    fmt.Println(PACKAGE, "<sourceDir> <sourceFile> <checkDir> <checkFile> <targetDir> <targetFile> <ifReplaceSource> <encryptMethod>")
    fmt.Println("")
    fmt.Println("ifReplaceSource:   Set 1 to remove sourceFile and move the targetFile to replace it. Set 0 to do nothing.")
    fmt.Println("encryptMethod  :   For example set 'AES' to use AES to encrypt, set 'AES-d' to use AES to decrypt and so on.")
}

func main () {
    if len(os.Args) != 9 {
        usage()
        os.Exit(1)
    }

    ctx := NewContext()
    if err := ctx.Init(os.Args); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    chkFile := NewCheckFile(ctx.CheckFile())
    if err := chkFile.LoadInfo(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    source, err := os.OpenFile(ctx.SourceFile(), os.O_RDONLY, 0664)
    if err != nil {
        fmt.Println(ctx.SourceFile(), ":", err)
        os.Exit(1)
    }
    defer source.Close()

    target, err := os.OpenFile(ctx.TargetFile(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
    if err != nil {
        fmt.Println(ctx.TargetFile(), ":", err)
        os.Exit(1)
    }
    defer target.Close()

    processCnt := lineProcess(source, target, process)

    // write log

    fmt.Printf("Deal with: %d Lines\n", processCnt)
    fmt.Printf("Runtime: %.2f Sec\n", 0.0)
}

func process(inLine string) (string, int) {
    return ">>" + inLine, 1
}

