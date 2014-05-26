package main

import (
    "errors"
    "file"
    "fmt"
    "os"
    "syscall"
)

const VERSION = "V14.00.001"

func daemon() (pid uintptr, err error) {
    pid, _, sysErr := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
    if sysErr != 0 {
        err = errors.New("Fork failed")
        return
    }

    if pid == 0 {
        syscall.Setsid()
    } else {
        os.Exit(0)
    }
    return pid, nil
}

var isShutdown bool = false

//var runnableQueue *task.TaskQueue

func main() {
    if len(os.Args) != 2 {
        fmt.Printf("Version: %s\n", VERSION)
        fmt.Printf("Usage: %s <configure file>\n", os.Args[0])
        os.Exit(1)
    }
    var cfgFilePath string = os.Args[1]

    pid, err := daemon()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    config := file.NewConfigure(cfgFilePath)
    if err := config.LoadConfigure(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    pidfile := file.NewPidFile(config.PidFileName())
    if err := pidfile.Init(pid); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer pidfile.Release()

    logfile := file.NewLogFile(config.LogFileName())
    if err := logfile.Init(config.LogMaxSize(), config.LogMaxIndex(), config.LogLevel()); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    finishfile := file.NewFinishFile(config.FinishFile())
    if err := finishfile.Init(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    global := file.NewGlobal()
    if config.TaskCount() < global.ThreadCount() {
        global.SetThreadCount(config.TaskCount())
    }

    //    runnableQueue = task.NewTaskQueue(config.Tasks())

    threadManager := manager.NewThreadManager()
    for i := 0; i < global.ThreadCount(); i++ {
        threadManager.CreateThread()
    }

    // main loop
    for {
        if isShutdown {
            if threadManager.AliveCount() == 0 {
                break
            } else {
                continue
            }
        }
        if current-last > 60 {
            last = current
            isShutdown = finishfile.FinishFlag()
            if isShutdown {
                continue
            }
        }

        printThreadstatistics()
        if _task, isTimeOut := waitQueue.TimePop(); isTimeOut == false {
            //            runnableQueue.Push(_task)
        } else {
            fmt.Println("Pop time out")
        }
    }

    os.Exit(0)
}
