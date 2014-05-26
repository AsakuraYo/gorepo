package sched

type JobBase struct {
    JobId int
    CategoryId string
    ExecName string
    Priority int
    Critical bool
    MaxInstance int
    ExecTimeExpr string
}

//func (jb *JobBase) IsCritical() bool {
//    return jb.critical
//}
//
//func (jb *JobBase) Priority() int {
//    return jb.priority
//}
//
//func (jb *JobBase) SetPriority(priority int) {
//    jb.priority = priority
//}
//
//func (jb *JobBase) JobId() int {
//    return jb.jobId
//}
//
//func (jb *JobBase) SetJobId(jobId int) {
//    jb.jobId = jobId
//}
//
//func (jb *JobBase) MaxInstance() int {
//    return jb.maxInstance
//}
//
//func (jb *JobBase) SetMaxInstance(maxInstance int) {
//    jb.maxInstance = maxInstance
//}
//
//func (jb *JobBase) CategoryId() string {
//    return jb.categoryId
//}
//
//func (jb *JobBase) SetCategoryId(categoryId string) {
//    jb.categoryId = categoryId
//}
//
//func (jb *JobBase) ExecName() string {
//        return jb.execName
//}
//
//func (jb *JobBase) SetExecName(execName string) {
//        jb.execName = execName
//}
//
//func (jb *JobBase) ExecTimeExpr() string {
//            return jb.execTimeExpr
//}
//
//func (jb *JobBase) SetExecTimeExpr(execTimeExpr string) {
//            jb.execTimeExpr = execTimeExpr
//}
