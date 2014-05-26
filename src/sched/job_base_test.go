package sched

import "testing"

func TestJobDef(t *testing.T) {
    var job JobBase
    job.JobId = 11326801
    job.CategoryId = "scan"
    job.ExecName = "scan"
    job.Priority = 30001
    job.ExecTimeExpr = "YYYY-MM-DD"
    t.Log(job)
}
