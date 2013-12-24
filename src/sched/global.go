// Package of scheduler core.
package sched

import "os"
import "errors"

const (
    HOME_VAR = "SCHED_HOME"
    LOG_DIR = "log"
    VAR_DIR = "var"
    ETC_DIR = "etc"
    CKPT_DIR = "ckpt"
    TRANSLOG_DIR = "var"
    DB_DIR = "db"
    DATA_DIR = "data"
    LIB_DIR = "lib"
)

var ErrEnvironmentNotSet = errors.New(HOME_VAR + " environment variable not set!")

func HomeDir() (homeDir string, err error) {
    homeDir = os.Getenv(HOME_VAR)
    if len(homeDir) == 0 {
        err = ErrEnvironmentNotSet
        return
    }
    return
}

func LogDir() (logDir string, err error) {
    logDir, err = HomeDir()
    if err != nil {
        return
    }
    logDir += "/" + LOG_DIR
    return
}

func VarDir() (varDir string, err error) {
    varDir, err = HomeDir()
    if err != nil {
        return
    }
    varDir += "/" + VAR_DIR
    return
}

func ConfigDir() (etcDir string, err error) {
    etcDir, err = HomeDir()
    if err != nil {
        return
    }
    etcDir += "/" + ETC_DIR
    return
}

func CkptDir() (ckptDir string, err error) {
    ckptDir, err = HomeDir()
    if err != nil {
        return
    }
    ckptDir += "/" + CKPT_DIR
    return
}

func TranslogDir() (translogDir string, err error) {
    translogDir, err = HomeDir()
    if err != nil {
        return
    }
    translogDir += "/" + TRANSLOG_DIR
    return
}

func DbDir() (dbDir string, err error) {
    dbDir, err = HomeDir()
    if err != nil {
        return
    }
    dbDir += "/" + DB_DIR
    return
}

func DataDir() (dataDir string, err error) {
    dataDir, err = HomeDir()
    if err != nil {
        return
    }
    dataDir += "/" + DATA_DIR
    return
}

func LibDir() (libDir string, err error) {
    libDir, err = HomeDir()
    if err != nil {
        return
    }
    libDir += "/" + LIB_DIR
    return
}

