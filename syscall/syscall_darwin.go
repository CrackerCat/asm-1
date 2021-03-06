package syscall

const (
	Nosys                          = 0
	Exit                           = 1
	Fork                           = 2
	Read                           = 3
	Write                          = 4
	Open                           = 5
	Close                          = 6
	Wait4                          = 7
	Link                           = 9
	Unlink                         = 10
	Chdir                          = 12
	Fchdir                         = 13
	Mknod                          = 14
	Chmod                          = 15
	Chown                          = 16
	Getfsstat                      = 18
	Getpid                         = 20
	Setuid                         = 23
	Getuid                         = 24
	Geteuid                        = 25
	Ptrace                         = 26
	Recvmsg                        = 27
	Sendmsg                        = 28
	Recvfrom                       = 29
	Accept                         = 30
	Getpeername                    = 31
	Getsockname                    = 32
	Access                         = 33
	Chflags                        = 34
	Fchflags                       = 35
	Sync                           = 36
	Kill                           = 37
	Getppid                        = 39
	Dup                            = 41
	Pipe                           = 42
	Getegid                        = 43
	Sigaction                      = 46
	Getgid                         = 47
	Sigprocmask                    = 48
	Getlogin                       = 49
	Setlogin                       = 50
	Acct                           = 51
	Sigpending                     = 52
	Sigaltstack                    = 53
	Ioctl                          = 54
	Reboot                         = 55
	Revoke                         = 56
	Symlink                        = 57
	Readlink                       = 58
	Execve                         = 59
	Umask                          = 60
	Chroot                         = 61
	Msync                          = 65
	Vfork                          = 66
	Munmap                         = 73
	Mprotect                       = 74
	Madvise                        = 75
	Mincore                        = 78
	Getgroups                      = 79
	Setgroups                      = 80
	Getpgrp                        = 81
	Setpgid                        = 82
	Setitimer                      = 83
	Swapon                         = 85
	Getitimer                      = 86
	Getdtablesize                  = 89
	Dup2                           = 90
	Fcntl                          = 92
	Select                         = 93
	Fsync                          = 95
	Setpriority                    = 96
	Socket                         = 97
	Connect                        = 98
	Getpriority                    = 100
	Bind                           = 104
	Setsockopt                     = 105
	Listen                         = 106
	Sigsuspend                     = 111
	Gettimeofday                   = 116
	Getrusage                      = 117
	Getsockopt                     = 118
	Readv                          = 120
	Writev                         = 121
	Settimeofday                   = 122
	Fchown                         = 123
	Fchmod                         = 124
	Setreuid                       = 126
	Setregid                       = 127
	Rename                         = 128
	Flock                          = 131
	Mkfifo                         = 132
	Sendto                         = 133
	Shutdown                       = 134
	Socketpair                     = 135
	Mkdir                          = 136
	Rmdir                          = 137
	Utimes                         = 138
	Futimes                        = 139
	Adjtime                        = 140
	Gethostuuid                    = 142
	Setsid                         = 147
	Getpgid                        = 151
	Setprivexec                    = 152
	Pread                          = 153
	Pwrite                         = 154
	Nfssvc                         = 155
	Statfs                         = 157
	Fstatfs                        = 158
	Unmount                        = 159
	Getfh                          = 161
	Quotactl                       = 165
	Mount                          = 167
	Csops                          = 169
	Csops_audittoken               = 170
	Waitid                         = 173
	Kdebug_typefilter              = 177
	Kdebug_trace_string            = 178
	Kdebug_trace64                 = 179
	Kdebug_trace                   = 180
	Setgid                         = 181
	Setegid                        = 182
	Seteuid                        = 183
	Sigreturn                      = 184
	Thread_selfcounts              = 186
	Fdatasync                      = 187
	Stat                           = 188
	Fstat                          = 189
	Lstat                          = 190
	Pathconf                       = 191
	Fpathconf                      = 192
	Getrlimit                      = 194
	Setrlimit                      = 195
	Getdirentries                  = 196
	Mmap                           = 197
	Lseek                          = 199
	Truncate                       = 200
	Ftruncate                      = 201
	Sysctl                         = 202
	Mlock                          = 203
	Munlock                        = 204
	Undelete                       = 205
	Open_dprotected_np             = 216
	Getattrlist                    = 220
	Setattrlist                    = 221
	Getdirentriesattr              = 222
	Exchangedata                   = 223
	Searchfs                       = 225
	Delete                         = 226
	Copyfile                       = 227
	Fgetattrlist                   = 228
	Fsetattrlist                   = 229
	Poll                           = 230
	Watchevent                     = 231
	Waitevent                      = 232
	Modwatch                       = 233
	Getxattr                       = 234
	Fgetxattr                      = 235
	Setxattr                       = 236
	Fsetxattr                      = 237
	Removexattr                    = 238
	Fremovexattr                   = 239
	Listxattr                      = 240
	Flistxattr                     = 241
	Fsctl                          = 242
	Initgroups                     = 243
	Posix_spawn                    = 244
	Ffsctl                         = 245
	Nfsclnt                        = 247
	Fhopen                         = 248
	Minherit                       = 250
	Semsys                         = 251
	Msgsys                         = 252
	Shmsys                         = 253
	Semctl                         = 254
	Semget                         = 255
	Semop                          = 256
	Msgctl                         = 258
	Msgget                         = 259
	Msgsnd                         = 260
	Msgrcv                         = 261
	Shmat                          = 262
	Shmctl                         = 263
	Shmdt                          = 264
	Shmget                         = 265
	Shm_open                       = 266
	Shm_unlink                     = 267
	Sem_open                       = 268
	Sem_close                      = 269
	Sem_unlink                     = 270
	Sem_wait                       = 271
	Sem_trywait                    = 272
	Sem_post                       = 273
	Sysctlbyname                   = 274
	Open_extended                  = 277
	Umask_extended                 = 278
	Stat_extended                  = 279
	Lstat_extended                 = 280
	Fstat_extended                 = 281
	Chmod_extended                 = 282
	Fchmod_extended                = 283
	Access_extended                = 284
	Settid                         = 285
	Gettid                         = 286
	Setsgroups                     = 287
	Getsgroups                     = 288
	Setwgroups                     = 289
	Getwgroups                     = 290
	Mkfifo_extended                = 291
	Mkdir_extended                 = 292
	Identitysvc                    = 293
	Shared_region_check_np         = 294
	Vm_pressure_monitor            = 296
	Psynch_rw_longrdlock           = 297
	Psynch_rw_yieldwrlock          = 298
	Psynch_rw_downgrade            = 299
	Psynch_rw_upgrade              = 300
	Psynch_mutexwait               = 301
	Psynch_mutexdrop               = 302
	Psynch_cvbroad                 = 303
	Psynch_cvsignal                = 304
	Psynch_cvwait                  = 305
	Psynch_rw_rdlock               = 306
	Psynch_rw_wrlock               = 307
	Psynch_rw_unlock               = 308
	Psynch_rw_unlock2              = 309
	Getsid                         = 310
	Settid_with_pid                = 311
	Psynch_cvclrprepost            = 312
	Aio_fsync                      = 313
	Aio_return                     = 314
	Aio_suspend                    = 315
	Aio_cancel                     = 316
	Aio_error                      = 317
	Aio_read                       = 318
	Aio_write                      = 319
	Lio_listio                     = 320
	Iopolicysys                    = 322
	Process_policy                 = 323
	Mlockall                       = 324
	Munlockall                     = 325
	Issetugid                      = 327
	Proc_info                      = 336
	Sendfile                       = 337
	Stat64                         = 338
	Fstat64                        = 339
	Lstat64                        = 340
	Stat64_extended                = 341
	Lstat64_extended               = 342
	Fstat64_extended               = 343
	Getdirentries64                = 344
	Statfs64                       = 345
	Fstatfs64                      = 346
	Getfsstat64                    = 347
	Audit                          = 350
	Auditon                        = 351
	Getauid                        = 353
	Setauid                        = 354
	Getaudit_addr                  = 357
	Setaudit_addr                  = 358
	Auditctl                       = 359
	Bsdthread_create               = 360
	Bsdthread_terminate            = 361
	Kqueue                         = 362
	Kevent                         = 363
	Lchown                         = 364
	Bsdthread_register             = 366
	Workq_open                     = 367
	Workq_kernreturn               = 368
	Kevent64                       = 369
	Thread_selfid                  = 372
	Ledger                         = 373
	Kevent_qos                     = 374
	Kevent_id                      = 375
	Pselect                        = 394
	Pselect_nocancel               = 395
	Read_nocancel                  = 396
	Write_nocancel                 = 397
	Open_nocancel                  = 398
	Close_nocancel                 = 399
	Wait4_nocancel                 = 400
	Recvmsg_nocancel               = 401
	Sendmsg_nocancel               = 402
	Recvfrom_nocancel              = 403
	Accept_nocancel                = 404
	Msync_nocancel                 = 405
	Fcntl_nocancel                 = 406
	Select_nocancel                = 407
	Fsync_nocancel                 = 408
	Connect_nocancel               = 409
	Sigsuspend_nocancel            = 410
	Readv_nocancel                 = 411
	Writev_nocancel                = 412
	Sendto_nocancel                = 413
	Pread_nocancel                 = 414
	Pwrite_nocancel                = 415
	Waitid_nocancel                = 416
	Poll_nocancel                  = 417
	Msgsnd_nocancel                = 418
	Msgrcv_nocancel                = 419
	Sem_wait_nocancel              = 420
	Aio_suspend_nocancel           = 421
	Fsgetpath                      = 427
	Audit_session_self             = 428
	Audit_session_join             = 429
	Fileport_makeport              = 430
	Fileport_makefd                = 431
	Audit_session_port             = 432
	Pid_suspend                    = 433
	Pid_resume                     = 434
	Pid_hibernate                  = 435
	Pid_shutdown_sockets           = 436
	Shared_region_map_and_slide_np = 438
	Kas_info                       = 439
	Memorystatus_control           = 440
	Guarded_open_np                = 441
	Guarded_close_np               = 442
	Guarded_kqueue_np              = 443
	Change_fdguard_np              = 444
	Usrctl                         = 445
	Proc_rlimit_control            = 446
	Connectx                       = 447
	Disconnectx                    = 448
	Peeloff                        = 449
	Socket_delegate                = 450
	Telemetry                      = 451
	Proc_uuid_policy               = 452
	Memorystatus_get_level         = 453
	System_override                = 454
	Vfs_purge                      = 455
	Sfi_ctl                        = 456
	Sfi_pidctl                     = 457
	Coalition                      = 458
	Coalition_info                 = 459
	Necp_match_policy              = 460
	Getattrlistbulk                = 461
	Clonefileat                    = 462
	Openat                         = 463
	Openat_nocancel                = 464
	Renameat                       = 465
	Faccessat                      = 466
	Fchmodat                       = 467
	Fchownat                       = 468
	Fstatat                        = 469
	Fstatat64                      = 470
	Linkat                         = 471
	Unlinkat                       = 472
	Readlinkat                     = 473
	Symlinkat                      = 474
	Mkdirat                        = 475
	Getattrlistat                  = 476
	Proc_trace_log                 = 477
	Bsdthread_ctl                  = 478
	Openbyid_np                    = 479
	Recvmsg_x                      = 480
	Sendmsg_x                      = 481
	Thread_selfusage               = 482
	Csrctl                         = 483
	Guarded_open_dprotected_np     = 484
	Guarded_write_np               = 485
	Guarded_pwrite_np              = 486
	Guarded_writev_np              = 487
	Renameatx_np                   = 488
	Mremap_encrypted               = 489
	Netagent_trigger               = 490
	Stack_snapshot_with_config     = 491
	Microstackshot                 = 492
	Grab_pgo_data                  = 493
	Persona                        = 494
	Work_interval_ctl              = 499
	Getentropy                     = 500
	Necp_open                      = 501
	Necp_client_action             = 502
	Ulock_wait                     = 515
	Ulock_wake                     = 516
	Fclonefileat                   = 517
	Fs_snapshot                    = 518
	Terminate_with_payload         = 520
	Abort_with_payload             = 521
	Necp_session_open              = 522
	Necp_session_action            = 523
	Setattrlistat                  = 524
	Net_qos_guideline              = 525
	Fmount                         = 526
	Ntp_adjtime                    = 527
	Ntp_gettime                    = 528
	Os_fault_with_payload          = 529
)
