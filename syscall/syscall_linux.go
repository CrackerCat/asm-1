package syscall

const (
	Read = iota
	Write
	Open
	Close
	Stat
	Fstat
	Lstat
	Poll
	Lseek
	Mmap
	Mprotect
	Munmap
	Brk
	Rt_sigaction
	Rt_sigprocmask
	Rt_sigreturn
	Ioctl
	Pread64
	Pwrite64
	Readv
	Writev
	Access
	Pipe
	Select
	Sched_yield
	Mremap
	Msync
	Mincore
	Madvise
	Shmget
	Shmat
	Shmctl
	Dup
	Dup2
	Pause
	Nanosleep
	Getitimer
	Alarm
	Setitimer
	Getpid
	Sendfile
	Socket
	Connect
	Accept
	Sendto
	Recvfrom
	Sendmsg
	Recvmsg
	Shutdown
	Bind
	Listen
	Getsockname
	Getpeername
	Socketpair
	Setsockopt
	Getsockopt
	Clone
	Fork
	Vfork
	Execve
	Exit
	Wait4
	Kill
	Uname
	Semget
	Semop
	Semctl
	Shmdt
	Msgget
	Msgsnd
	Msgrcv
	Msgctl
	Fcntl
	Flock
	Fsync
	Fdatasync
	Truncate
	Ftruncate
	Getdents
	Getcwd
	Chdir
	Fchdir
	Rename
	Mkdir
	Rmdir
	Creat
	Link
	Unlink
	Symlink
	Readlink
	Chmod
	Fchmod
	Chown
	Fchown
	Lchown
	Umask
	Gettimeofday
	Getrlimit
	Getrusage
	Sysinfo
	Times
	Ptrace
	Getuid
	Syslog
	Getgid
	Setuid
	Setgid
	Geteuid
	Getegid
	Setpgid
	Getppid
	Getpgrp
	Setsid
	Setreuid
	Setregid
	Getgroups
	Setgroups
	Setresuid
	Getresuid
	Setresgid
	Getresgid
	Getpgid
	Setfsuid
	Setfsgid
	Getsid
	Capget
	Capset
	Rt_sigpending
	Rt_sigtimedwait
	Rt_sigqueueinfo
	Rt_sigsuspend
	Sigaltstack
	Utime
	Mknod
	Uselib
	Personality
	Ustat
	Statfs
	Fstatfs
	Sysfs
	Getpriority
	Setpriority
	Sched_setparam
	Sched_getparam
	Sched_setscheduler
	Sched_getscheduler
	Sched_get_priority_max
	Sched_get_priority_min
	Sched_rr_get_interval
	Mlock
	Munlock
	Mlockall
	Munlockall
	Vhangup
	Modify_ldt
	Pivot_root
	Sysctl
	Prctl
	Arch_prctl
	Adjtimex
	Setrlimit
	Chroot
	Sync
	Acct
	Settimeofday
	Mount
	Umount2
	Swapon
	Swapoff
	Reboot
	Sethostname
	Setdomainname
	Iopl
	Ioperm
	Create_module
	Init_module
	Delete_module
	Get_kernel_syms
	Query_module
	Quotactl
	Nfsservctl
	Getpmsg
	Putpmsg
	Afs_syscall
	Tuxcall
	Security
	Gettid
	Readahead
	Setxattr
	Lsetxattr
	Fsetxattr
	Getxattr
	Lgetxattr
	Fgetxattr
	Listxattr
	Llistxattr
	Flistxattr
	Removexattr
	Lremovexattr
	Fremovexattr
	Tkill
	Time
	Futex
	Sched_setaffinity
	Sched_getaffinity
	Set_thread_area
	Io_setup
	Io_destroy
	Io_getevents
	Io_submit
	Io_cancel
	Get_thread_area
	Lookup_dcookie
	Epoll_create
	Epoll_ctl_old
	Epoll_wait_old
	Remap_file_pages
	Getdents64
	Set_tid_address
	Restart_syscall
	Semtimedop
	Fadvise64
	Timer_create
	Timer_settime
	Timer_gettime
	Timer_getoverrun
	Timer_delete
	Clock_settime
	Clock_gettime
	Clock_getres
	Clock_nanosleep
	Exit_group
	Epoll_wait
	Epoll_ctl
	Tgkill
	Utimes
	Vserver
	Mbind
	Set_mempolicy
	Get_mempolicy
	Mq_open
	Mq_unlink
	Mq_timedsend
	Mq_timedreceive
	Mq_notify
	Mq_getsetattr
	Kexec_load
	Waitid
	Add_key
	Request_key
	Keyctl
	Ioprio_set
	Ioprio_get
	Inotify_init
	Inotify_add_watch
	Inotify_rm_watch
	Migrate_pages
	Openat
	Mkdirat
	Mknodat
	Fchownat
	Futimesat
	Newfstatat
	Unlinkat
	Renameat
	Linkat
	Symlinkat
	Readlinkat
	Fchmodat
	Faccessat
	Pselect6
	Ppoll
	Unshare
	Set_robust_list
	Get_robust_list
	Splice
	Tee
	Sync_file_range
	Vmsplice
	Move_pages
	Utimensat
	Epoll_pwait
	Signalfd
	Timerfd_create
	Eventfd
	Fallocate
	Timerfd_settime
	Timerfd_gettime
	Accept4
	Signalfd4
	Eventfd2
	Epoll_create1
	Dup3
	Pipe2
	Inotify_init1
	Preadv
	Pwritev
	Rt_tgsigqueueinfo
	Perf_event_open
	Recvmmsg
	Fanotify_init
	Fanotify_mark
	Prlimit64
	Name_to_handle_at
	Open_by_handle_at
	Clock_adjtime
	Syncfs
	Sendmmsg
	Setns
	Getcpu
	Process_vm_readv
	Process_vm_writev
	Kcmp
	Finit_module
	Sched_setattr
	Sched_getattr
	Renameat2
	Seccomp
	Getrandom
	Memfd_create
	Kexec_file_load
	Bpf
	Stub_execveat
	Userfaultfd
	Membarrier
	Mlock2
	Copy_file_range
	Preadv2
	Pwritev2
)
