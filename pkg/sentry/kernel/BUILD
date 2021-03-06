package(licenses = ["notice"])  # Apache 2.0

load("//tools/go_generics:defs.bzl", "go_template_instance")
load("//tools/go_stateify:defs.bzl", "go_library", "go_test")

go_template_instance(
    name = "pending_signals_list",
    out = "pending_signals_list.go",
    package = "kernel",
    prefix = "pendingSignal",
    template = "//pkg/ilist:generic_list",
    types = {
        "Linker": "*pendingSignal",
    },
)

go_template_instance(
    name = "process_group_list",
    out = "process_group_list.go",
    package = "kernel",
    prefix = "processGroup",
    template = "//pkg/ilist:generic_list",
    types = {
        "Linker": "*ProcessGroup",
    },
)

go_template_instance(
    name = "seqatomic_taskgoroutineschedinfo",
    out = "seqatomic_taskgoroutineschedinfo.go",
    package = "kernel",
    suffix = "TaskGoroutineSchedInfo",
    template = "//pkg/sync:generic_seqatomic",
    types = {
        "Value": "TaskGoroutineSchedInfo",
    },
)

go_template_instance(
    name = "session_list",
    out = "session_list.go",
    package = "kernel",
    prefix = "session",
    template = "//pkg/ilist:generic_list",
    types = {
        "Linker": "*Session",
    },
)

go_template_instance(
    name = "task_list",
    out = "task_list.go",
    package = "kernel",
    prefix = "task",
    template = "//pkg/ilist:generic_list",
    types = {
        "Linker": "*Task",
    },
)

go_library(
    name = "kernel",
    srcs = [
        "abstract_socket_namespace.go",
        "context.go",
        "fd_map.go",
        "fs_context.go",
        "ipc_namespace.go",
        "kernel.go",
        "kernel_state.go",
        "pending_signals.go",
        "pending_signals_list.go",
        "pending_signals_state.go",
        "process_group_list.go",
        "ptrace.go",
        "rseq.go",
        "seccomp.go",
        "seqatomic_taskgoroutineschedinfo.go",
        "session_list.go",
        "sessions.go",
        "signal.go",
        "signal_handlers.go",
        "syscalls.go",
        "syscalls_state.go",
        "syslog.go",
        "task.go",
        "task_acct.go",
        "task_block.go",
        "task_clone.go",
        "task_context.go",
        "task_exec.go",
        "task_exit.go",
        "task_identity.go",
        "task_list.go",
        "task_log.go",
        "task_net.go",
        "task_resources.go",
        "task_run.go",
        "task_sched.go",
        "task_signals.go",
        "task_start.go",
        "task_stop.go",
        "task_syscall.go",
        "task_usermem.go",
        "thread_group.go",
        "threads.go",
        "timekeeper.go",
        "timekeeper_state.go",
        "timer.go",
        "uts_namespace.go",
        "vdso.go",
        "version.go",
    ],
    importpath = "gvisor.googlesource.com/gvisor/pkg/sentry/kernel",
    imports = [
        "gvisor.googlesource.com/gvisor/pkg/sentry/arch",
        # "gvisor.googlesource.com/gvisor/pkg/sentry/kernel/kdefs",
        "gvisor.googlesource.com/gvisor/pkg/tcpip",
    ],
    visibility = ["//:sandbox"],
    deps = [
        "//pkg/abi",
        "//pkg/abi/linux",
        "//pkg/amutex",
        "//pkg/binary",
        "//pkg/bits",
        "//pkg/bpf",
        "//pkg/cpuid",
        "//pkg/eventchannel",
        "//pkg/log",
        "//pkg/refs",
        "//pkg/secio",
        "//pkg/sentry/arch",
        "//pkg/sentry/context",
        "//pkg/sentry/fs",
        "//pkg/sentry/fs/lock",
        "//pkg/sentry/fs/timerfd",
        "//pkg/sentry/hostcpu",
        "//pkg/sentry/inet",
        "//pkg/sentry/kernel/auth",
        "//pkg/sentry/kernel/epoll",
        "//pkg/sentry/kernel/futex",
        "//pkg/sentry/kernel/kdefs",
        "//pkg/sentry/kernel/sched",
        "//pkg/sentry/kernel/semaphore",
        "//pkg/sentry/kernel/shm",
        "//pkg/sentry/kernel/time",
        "//pkg/sentry/limits",
        "//pkg/sentry/loader",
        "//pkg/sentry/memmap",
        "//pkg/sentry/mm",
        "//pkg/sentry/platform",
        "//pkg/sentry/safemem",
        "//pkg/sentry/socket/netlink/port",
        "//pkg/sentry/time",
        "//pkg/sentry/uniqueid",
        "//pkg/sentry/usage",
        "//pkg/sentry/usermem",
        "//pkg/state",
        "//pkg/state/statefile",
        "//pkg/sync",
        "//pkg/syserror",
        "//pkg/tcpip",
        "//pkg/tcpip/stack",
        "//pkg/tcpip/transport/unix",
        "//pkg/waiter",
    ],
)

go_test(
    name = "kernel_test",
    size = "small",
    srcs = [
        "fd_map_test.go",
        "table_test.go",
        "task_test.go",
        "timekeeper_test.go",
    ],
    embed = [":kernel"],
    deps = [
        "//pkg/abi",
        "//pkg/sentry/arch",
        "//pkg/sentry/context/contexttest",
        "//pkg/sentry/fs/filetest",
        "//pkg/sentry/kernel/kdefs",
        "//pkg/sentry/kernel/sched",
        "//pkg/sentry/limits",
        "//pkg/sentry/platform",
        "//pkg/sentry/time",
        "//pkg/sentry/usage",
        "//pkg/sentry/usermem",
        "//pkg/syserror",
    ],
)
