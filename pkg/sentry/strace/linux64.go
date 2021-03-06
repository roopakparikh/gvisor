// Copyright 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package strace

// linuxAMD64 provides a mapping of the Linux amd64 syscalls and their argument
// types for display / formatting.
var linuxAMD64 = SyscallMap{
	0:   makeSyscallInfo("read", Hex, ReadBuffer, Hex),
	1:   makeSyscallInfo("write", Hex, WriteBuffer, Hex),
	2:   makeSyscallInfo("open", Path, OpenFlags, Mode),
	3:   makeSyscallInfo("close", Hex),
	4:   makeSyscallInfo("stat", Path, Stat),
	5:   makeSyscallInfo("fstat", Hex, Stat),
	6:   makeSyscallInfo("lstat", Path, Stat),
	7:   makeSyscallInfo("poll", Hex, Hex, Hex),
	8:   makeSyscallInfo("lseek", Hex, Hex, Hex),
	9:   makeSyscallInfo("mmap", Hex, Hex, Hex, Hex, Hex, Hex),
	10:  makeSyscallInfo("mprotect", Hex, Hex, Hex),
	11:  makeSyscallInfo("munmap", Hex, Hex),
	12:  makeSyscallInfo("brk", Hex),
	13:  makeSyscallInfo("rt_sigaction", Hex, Hex, Hex),
	14:  makeSyscallInfo("rt_sigprocmask", Hex, Hex, Hex, Hex),
	15:  makeSyscallInfo("rt_sigreturn"),
	16:  makeSyscallInfo("ioctl", Hex, Hex, Hex),
	17:  makeSyscallInfo("pread64", Hex, ReadBuffer, Hex, Hex),
	18:  makeSyscallInfo("pwrite64", Hex, WriteBuffer, Hex, Hex),
	19:  makeSyscallInfo("readv", Hex, ReadIOVec, Hex),
	20:  makeSyscallInfo("writev", Hex, WriteIOVec, Hex),
	21:  makeSyscallInfo("access", Path, Oct),
	22:  makeSyscallInfo("pipe", PipeFDs),
	23:  makeSyscallInfo("select", Hex, Hex, Hex, Hex, Timeval),
	24:  makeSyscallInfo("sched_yield"),
	25:  makeSyscallInfo("mremap", Hex, Hex, Hex, Hex, Hex),
	26:  makeSyscallInfo("msync", Hex, Hex, Hex),
	27:  makeSyscallInfo("mincore", Hex, Hex, Hex),
	28:  makeSyscallInfo("madvise", Hex, Hex, Hex),
	29:  makeSyscallInfo("shmget", Hex, Hex, Hex),
	30:  makeSyscallInfo("shmat", Hex, Hex, Hex),
	31:  makeSyscallInfo("shmctl", Hex, Hex, Hex),
	32:  makeSyscallInfo("dup", Hex),
	33:  makeSyscallInfo("dup2", Hex, Hex),
	34:  makeSyscallInfo("pause"),
	35:  makeSyscallInfo("nanosleep", Timespec, PostTimespec),
	36:  makeSyscallInfo("getitimer", Hex, PostItimerVal),
	37:  makeSyscallInfo("alarm", Hex),
	38:  makeSyscallInfo("setitimer", Hex, ItimerVal, PostItimerVal),
	39:  makeSyscallInfo("getpid"),
	40:  makeSyscallInfo("sendfile", Hex, Hex, Hex, Hex),
	41:  makeSyscallInfo("socket", SockFamily, SockType, SockProtocol),
	42:  makeSyscallInfo("connect", Hex, SockAddr, Hex),
	43:  makeSyscallInfo("accept", Hex, PostSockAddr, SockLen),
	44:  makeSyscallInfo("sendto", Hex, Hex, Hex, Hex, SockAddr, Hex),
	45:  makeSyscallInfo("recvfrom", Hex, Hex, Hex, Hex, PostSockAddr, SockLen),
	46:  makeSyscallInfo("sendmsg", Hex, SendMsgHdr, Hex),
	47:  makeSyscallInfo("recvmsg", Hex, RecvMsgHdr, Hex),
	48:  makeSyscallInfo("shutdown", Hex, Hex),
	49:  makeSyscallInfo("bind", Hex, SockAddr, Hex),
	50:  makeSyscallInfo("listen", Hex, Hex),
	51:  makeSyscallInfo("getsockname", Hex, PostSockAddr, SockLen),
	52:  makeSyscallInfo("getpeername", Hex, PostSockAddr, SockLen),
	53:  makeSyscallInfo("socketpair", SockFamily, SockType, SockProtocol, Hex),
	54:  makeSyscallInfo("setsockopt", Hex, Hex, Hex, Hex, Hex),
	55:  makeSyscallInfo("getsockopt", Hex, Hex, Hex, Hex, Hex),
	56:  makeSyscallInfo("clone", CloneFlags, Hex, Hex, Hex, Hex),
	57:  makeSyscallInfo("fork"),
	58:  makeSyscallInfo("vfork"),
	59:  makeSyscallInfo("execve", Path, ExecveStringVector, ExecveStringVector),
	60:  makeSyscallInfo("exit", Hex),
	61:  makeSyscallInfo("wait4", Hex, Hex, Hex, Rusage),
	62:  makeSyscallInfo("kill", Hex, Hex),
	63:  makeSyscallInfo("uname", Uname),
	64:  makeSyscallInfo("semget", Hex, Hex, Hex),
	65:  makeSyscallInfo("semop", Hex, Hex, Hex),
	66:  makeSyscallInfo("semctl", Hex, Hex, Hex, Hex),
	67:  makeSyscallInfo("shmdt", Hex),
	68:  makeSyscallInfo("msgget", Hex, Hex),
	69:  makeSyscallInfo("msgsnd", Hex, Hex, Hex, Hex),
	70:  makeSyscallInfo("msgrcv", Hex, Hex, Hex, Hex, Hex),
	71:  makeSyscallInfo("msgctl", Hex, Hex, Hex),
	72:  makeSyscallInfo("fcntl", Hex, Hex, Hex),
	73:  makeSyscallInfo("flock", Hex, Hex),
	74:  makeSyscallInfo("fsync", Hex),
	75:  makeSyscallInfo("fdatasync", Hex),
	76:  makeSyscallInfo("truncate", Path, Hex),
	77:  makeSyscallInfo("ftruncate", Hex, Hex),
	78:  makeSyscallInfo("getdents", Hex, Hex, Hex),
	79:  makeSyscallInfo("getcwd", PostPath, Hex),
	80:  makeSyscallInfo("chdir", Path),
	81:  makeSyscallInfo("fchdir", Hex),
	82:  makeSyscallInfo("rename", Path, Path),
	83:  makeSyscallInfo("mkdir", Path, Oct),
	84:  makeSyscallInfo("rmdir", Path),
	85:  makeSyscallInfo("creat", Path, Oct),
	86:  makeSyscallInfo("link", Path, Path),
	87:  makeSyscallInfo("unlink", Path),
	88:  makeSyscallInfo("symlink", Path, Path),
	89:  makeSyscallInfo("readlink", Path, ReadBuffer, Hex),
	90:  makeSyscallInfo("chmod", Path, Mode),
	91:  makeSyscallInfo("fchmod", Hex, Mode),
	92:  makeSyscallInfo("chown", Path, Hex, Hex),
	93:  makeSyscallInfo("fchown", Hex, Hex, Hex),
	94:  makeSyscallInfo("lchown", Hex, Hex, Hex),
	95:  makeSyscallInfo("umask", Hex),
	96:  makeSyscallInfo("gettimeofday", Timeval, Hex),
	97:  makeSyscallInfo("getrlimit", Hex, Hex),
	98:  makeSyscallInfo("getrusage", Hex, Rusage),
	99:  makeSyscallInfo("sysinfo", Hex),
	100: makeSyscallInfo("times", Hex),
	101: makeSyscallInfo("ptrace", PtraceRequest, Hex, Hex, Hex),
	102: makeSyscallInfo("getuid"),
	103: makeSyscallInfo("syslog", Hex, Hex, Hex),
	104: makeSyscallInfo("getgid"),
	105: makeSyscallInfo("setuid", Hex),
	106: makeSyscallInfo("setgid", Hex),
	107: makeSyscallInfo("geteuid"),
	108: makeSyscallInfo("getegid"),
	109: makeSyscallInfo("setpgid", Hex, Hex),
	110: makeSyscallInfo("getppid"),
	111: makeSyscallInfo("getpgrp"),
	112: makeSyscallInfo("setsid"),
	113: makeSyscallInfo("setreuid", Hex, Hex),
	114: makeSyscallInfo("setregid", Hex, Hex),
	115: makeSyscallInfo("getgroups", Hex, Hex),
	116: makeSyscallInfo("setgroups", Hex, Hex),
	117: makeSyscallInfo("setresuid", Hex, Hex, Hex),
	118: makeSyscallInfo("getresuid", Hex, Hex, Hex),
	119: makeSyscallInfo("setresgid", Hex, Hex, Hex),
	120: makeSyscallInfo("getresgid", Hex, Hex, Hex),
	121: makeSyscallInfo("getpgid", Hex),
	122: makeSyscallInfo("setfsuid", Hex),
	123: makeSyscallInfo("setfsgid", Hex),
	124: makeSyscallInfo("getsid", Hex),
	125: makeSyscallInfo("capget", Hex, Hex),
	126: makeSyscallInfo("capset", Hex, Hex),
	127: makeSyscallInfo("rt_sigpending", Hex),
	128: makeSyscallInfo("rt_sigtimedwait", Hex, Hex, Timespec, Hex),
	129: makeSyscallInfo("rt_sigqueueinfo", Hex, Hex, Hex),
	130: makeSyscallInfo("rt_sigsuspend", Hex),
	131: makeSyscallInfo("sigaltstack", Hex, Hex),
	132: makeSyscallInfo("utime", Path, Utimbuf),
	133: makeSyscallInfo("mknod", Path, Mode, Hex),
	134: makeSyscallInfo("uselib", Hex),
	135: makeSyscallInfo("personality", Hex),
	136: makeSyscallInfo("ustat", Hex, Hex),
	137: makeSyscallInfo("statfs", Path, Hex),
	138: makeSyscallInfo("fstatfs", Hex, Hex),
	139: makeSyscallInfo("sysfs", Hex, Hex, Hex),
	140: makeSyscallInfo("getpriority", Hex, Hex),
	141: makeSyscallInfo("setpriority", Hex, Hex, Hex),
	142: makeSyscallInfo("sched_setparam", Hex, Hex),
	143: makeSyscallInfo("sched_getparam", Hex, Hex),
	144: makeSyscallInfo("sched_setscheduler", Hex, Hex, Hex),
	145: makeSyscallInfo("sched_getscheduler", Hex),
	146: makeSyscallInfo("sched_get_priority_max", Hex),
	147: makeSyscallInfo("sched_get_priority_min", Hex),
	148: makeSyscallInfo("sched_rr_get_interval", Hex, Hex),
	149: makeSyscallInfo("mlock", Hex, Hex),
	150: makeSyscallInfo("munlock", Hex, Hex),
	151: makeSyscallInfo("mlockall", Hex),
	152: makeSyscallInfo("munlockall"),
	153: makeSyscallInfo("vhangup"),
	154: makeSyscallInfo("modify_ldt", Hex, Hex, Hex),
	155: makeSyscallInfo("pivot_root", Hex, Hex),
	156: makeSyscallInfo("_sysctl", Hex),
	157: makeSyscallInfo("prctl", Hex, Hex, Hex, Hex, Hex),
	158: makeSyscallInfo("arch_prctl", Hex, Hex),
	159: makeSyscallInfo("adjtimex", Hex),
	160: makeSyscallInfo("setrlimit", Hex, Hex),
	161: makeSyscallInfo("chroot", Path),
	162: makeSyscallInfo("sync"),
	163: makeSyscallInfo("acct", Hex),
	164: makeSyscallInfo("settimeofday", Timeval, Hex),
	165: makeSyscallInfo("mount", Path, Path, Path, Hex, Path),
	166: makeSyscallInfo("umount2", Path, Hex),
	167: makeSyscallInfo("swapon", Hex, Hex),
	168: makeSyscallInfo("swapoff", Hex),
	169: makeSyscallInfo("reboot", Hex, Hex, Hex, Hex),
	170: makeSyscallInfo("sethostname", Hex, Hex),
	171: makeSyscallInfo("setdomainname", Hex, Hex),
	172: makeSyscallInfo("iopl", Hex),
	173: makeSyscallInfo("ioperm", Hex, Hex, Hex),
	174: makeSyscallInfo("create_module", Path, Hex),
	175: makeSyscallInfo("init_module", Hex, Hex, Hex),
	176: makeSyscallInfo("delete_module", Hex, Hex),
	177: makeSyscallInfo("get_kernel_syms", Hex),
	// 178: query_module (only present in Linux < 2.6)
	179: makeSyscallInfo("quotactl", Hex, Hex, Hex, Hex),
	180: makeSyscallInfo("nfsservctl", Hex, Hex, Hex),
	// 181: getpmsg (not implemented in the Linux kernel)
	// 182: putpmsg (not implemented in the Linux kernel)
	// 183: afs_syscall (not implemented in the Linux kernel)
	// 184: tuxcall (not implemented in the Linux kernel)
	// 185: security (not implemented in the Linux kernel)
	186: makeSyscallInfo("gettid"),
	187: makeSyscallInfo("readahead", Hex, Hex, Hex),
	188: makeSyscallInfo("setxattr", Path, Path, Hex, Hex, Hex),
	189: makeSyscallInfo("lsetxattr", Path, Path, Hex, Hex, Hex),
	190: makeSyscallInfo("fsetxattr", Hex, Path, Hex, Hex, Hex),
	191: makeSyscallInfo("getxattr", Path, Path, Hex, Hex),
	192: makeSyscallInfo("lgetxattr", Path, Path, Hex, Hex),
	193: makeSyscallInfo("fgetxattr", Hex, Path, Hex, Hex),
	194: makeSyscallInfo("listxattr", Path, Path, Hex),
	195: makeSyscallInfo("llistxattr", Path, Path, Hex),
	196: makeSyscallInfo("flistxattr", Hex, Path, Hex),
	197: makeSyscallInfo("removexattr", Path, Path),
	198: makeSyscallInfo("lremovexattr", Path, Path),
	199: makeSyscallInfo("fremovexattr", Hex, Path),
	200: makeSyscallInfo("tkill", Hex, Hex),
	201: makeSyscallInfo("time", Hex),
	202: makeSyscallInfo("futex", Hex, FutexOp, Hex, Timespec, Hex, Hex),
	203: makeSyscallInfo("sched_setaffinity", Hex, Hex, Hex),
	204: makeSyscallInfo("sched_getaffinity", Hex, Hex, Hex),
	205: makeSyscallInfo("set_thread_area", Hex),
	206: makeSyscallInfo("io_setup", Hex, Hex),
	207: makeSyscallInfo("io_destroy", Hex),
	208: makeSyscallInfo("io_getevents", Hex, Hex, Hex, Hex, Timespec),
	209: makeSyscallInfo("io_submit", Hex, Hex, Hex),
	210: makeSyscallInfo("io_cancel", Hex, Hex, Hex),
	211: makeSyscallInfo("get_thread_area", Hex),
	212: makeSyscallInfo("lookup_dcookie", Hex, Hex, Hex),
	213: makeSyscallInfo("epoll_create", Hex),
	// 214: epoll_ctl_old (not implemented in the Linux kernel)
	// 215: epoll_wait_old (not implemented in the Linux kernel)
	216: makeSyscallInfo("remap_file_pages", Hex, Hex, Hex, Hex, Hex),
	217: makeSyscallInfo("getdents64", Hex, Hex, Hex),
	218: makeSyscallInfo("set_tid_address", Hex),
	219: makeSyscallInfo("restart_syscall"),
	220: makeSyscallInfo("semtimedop", Hex, Hex, Hex, Hex),
	221: makeSyscallInfo("fadvise64", Hex, Hex, Hex, Hex),
	222: makeSyscallInfo("timer_create", Hex, Hex, Hex),
	223: makeSyscallInfo("timer_settime", Hex, Hex, Hex, Hex),
	224: makeSyscallInfo("timer_gettime", Hex, Hex),
	225: makeSyscallInfo("timer_getoverrun", Hex),
	226: makeSyscallInfo("timer_delete", Hex),
	227: makeSyscallInfo("clock_settime", Hex, Timespec),
	228: makeSyscallInfo("clock_gettime", Hex, PostTimespec),
	229: makeSyscallInfo("clock_getres", Hex, PostTimespec),
	230: makeSyscallInfo("clock_nanosleep", Hex, Hex, Timespec, PostTimespec),
	231: makeSyscallInfo("exit_group", Hex),
	232: makeSyscallInfo("epoll_wait", Hex, Hex, Hex, Hex),
	233: makeSyscallInfo("epoll_ctl", Hex, Hex, Hex, Hex),
	234: makeSyscallInfo("tgkill", Hex, Hex, Hex),
	235: makeSyscallInfo("utimes", Path, Timeval),
	// 236: vserver (not implemented in the Linux kernel)
	237: makeSyscallInfo("mbind", Hex, Hex, Hex, Hex, Hex, Hex),
	238: makeSyscallInfo("set_mempolicy", Hex, Hex, Hex),
	239: makeSyscallInfo("get_mempolicy", Hex, Hex, Hex, Hex, Hex),
	240: makeSyscallInfo("mq_open", Hex, Hex, Hex, Hex),
	241: makeSyscallInfo("mq_unlink", Hex),
	242: makeSyscallInfo("mq_timedsend", Hex, Hex, Hex, Hex, Hex),
	243: makeSyscallInfo("mq_timedreceive", Hex, Hex, Hex, Hex, Hex),
	244: makeSyscallInfo("mq_notify", Hex, Hex),
	245: makeSyscallInfo("mq_getsetattr", Hex, Hex, Hex),
	246: makeSyscallInfo("kexec_load", Hex, Hex, Hex, Hex),
	247: makeSyscallInfo("waitid", Hex, Hex, Hex, Hex, Rusage),
	248: makeSyscallInfo("add_key", Hex, Hex, Hex, Hex, Hex),
	249: makeSyscallInfo("request_key", Hex, Hex, Hex, Hex),
	250: makeSyscallInfo("keyctl", Hex, Hex, Hex, Hex, Hex),
	251: makeSyscallInfo("ioprio_set", Hex, Hex, Hex),
	252: makeSyscallInfo("ioprio_get", Hex, Hex),
	253: makeSyscallInfo("inotify_init"),
	254: makeSyscallInfo("inotify_add_watch", Hex, Hex, Hex),
	255: makeSyscallInfo("inotify_rm_watch", Hex, Hex),
	256: makeSyscallInfo("migrate_pages", Hex, Hex, Hex, Hex),
	257: makeSyscallInfo("openat", Hex, Path, OpenFlags, Mode),
	258: makeSyscallInfo("mkdirat", Hex, Path, Hex),
	259: makeSyscallInfo("mknodat", Hex, Path, Mode, Hex),
	260: makeSyscallInfo("fchownat", Hex, Path, Hex, Hex, Hex),
	261: makeSyscallInfo("futimesat", Hex, Path, Hex),
	262: makeSyscallInfo("newfstatat", Hex, Path, Stat, Hex),
	263: makeSyscallInfo("unlinkat", Hex, Path, Hex),
	264: makeSyscallInfo("renameat", Hex, Path, Hex, Path),
	265: makeSyscallInfo("linkat", Hex, Path, Hex, Path, Hex),
	266: makeSyscallInfo("symlinkat", Path, Hex, Path),
	267: makeSyscallInfo("readlinkat", Hex, Path, ReadBuffer, Hex),
	268: makeSyscallInfo("fchmodat", Hex, Path, Mode),
	269: makeSyscallInfo("faccessat", Hex, Path, Oct, Hex),
	270: makeSyscallInfo("pselect6", Hex, Hex, Hex, Hex, Hex, Hex),
	271: makeSyscallInfo("ppoll", Hex, Hex, Timespec, Hex, Hex),
	272: makeSyscallInfo("unshare", Hex),
	273: makeSyscallInfo("set_robust_list", Hex, Hex),
	274: makeSyscallInfo("get_robust_list", Hex, Hex, Hex),
	275: makeSyscallInfo("splice", Hex, Hex, Hex, Hex, Hex, Hex),
	276: makeSyscallInfo("tee", Hex, Hex, Hex, Hex),
	277: makeSyscallInfo("sync_file_range", Hex, Hex, Hex, Hex),
	278: makeSyscallInfo("vmsplice", Hex, Hex, Hex, Hex),
	279: makeSyscallInfo("move_pages", Hex, Hex, Hex, Hex, Hex, Hex),
	280: makeSyscallInfo("utimensat", Hex, Path, UTimeTimespec, Hex),
	281: makeSyscallInfo("epoll_pwait", Hex, Hex, Hex, Hex, Hex, Hex),
	282: makeSyscallInfo("signalfd", Hex, Hex, Hex),
	283: makeSyscallInfo("timerfd_create", Hex, Hex),
	284: makeSyscallInfo("eventfd", Hex),
	285: makeSyscallInfo("fallocate", Hex, Hex, Hex, Hex),
	286: makeSyscallInfo("timerfd_settime", Hex, Hex, Hex, Hex),
	287: makeSyscallInfo("timerfd_gettime", Hex, Hex),
	288: makeSyscallInfo("accept4", Hex, PostSockAddr, SockLen, SockFlags),
	289: makeSyscallInfo("signalfd4", Hex, Hex, Hex, Hex),
	290: makeSyscallInfo("eventfd2", Hex, Hex),
	291: makeSyscallInfo("epoll_create1", Hex),
	292: makeSyscallInfo("dup3", Hex, Hex, Hex),
	293: makeSyscallInfo("pipe2", PipeFDs, Hex),
	294: makeSyscallInfo("inotify_init1", Hex),
	295: makeSyscallInfo("preadv", Hex, ReadIOVec, Hex, Hex),
	296: makeSyscallInfo("pwritev", Hex, WriteIOVec, Hex, Hex),
	297: makeSyscallInfo("rt_tgsigqueueinfo", Hex, Hex, Hex, Hex),
	298: makeSyscallInfo("perf_event_open", Hex, Hex, Hex, Hex, Hex),
	299: makeSyscallInfo("recvmmsg", Hex, Hex, Hex, Hex, Hex),
	300: makeSyscallInfo("fanotify_init", Hex, Hex),
	301: makeSyscallInfo("fanotify_mark", Hex, Hex, Hex, Hex, Hex),
	302: makeSyscallInfo("prlimit64", Hex, Hex, Hex, Hex),
	303: makeSyscallInfo("name_to_handle_at", Hex, Hex, Hex, Hex, Hex),
	304: makeSyscallInfo("open_by_handle_at", Hex, Hex, Hex),
	305: makeSyscallInfo("clock_adjtime", Hex, Hex),
	306: makeSyscallInfo("syncfs", Hex),
	307: makeSyscallInfo("sendmmsg", Hex, Hex, Hex, Hex),
	308: makeSyscallInfo("setns", Hex, Hex),
	309: makeSyscallInfo("getcpu", Hex, Hex, Hex),
	310: makeSyscallInfo("process_vm_readv", Hex, ReadIOVec, Hex, IOVec, Hex, Hex),
	311: makeSyscallInfo("process_vm_writev", Hex, IOVec, Hex, WriteIOVec, Hex, Hex),
	312: makeSyscallInfo("kcmp", Hex, Hex, Hex, Hex, Hex),
	313: makeSyscallInfo("finit_module", Hex, Hex, Hex),
	314: makeSyscallInfo("sched_setattr", Hex, Hex, Hex),
	315: makeSyscallInfo("sched_getattr", Hex, Hex, Hex),
	316: makeSyscallInfo("renameat2", Hex, Path, Hex, Path, Hex),
	317: makeSyscallInfo("seccomp", Hex, Hex, Hex),
}
