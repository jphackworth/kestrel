# Design Goals

The following are the design goals for zlarkd. There are some differences from cjdroute.

1. Interoperate with existing cjdroute
2. Linux support first. Others platforms are secondary
3. Daemon process can open files. Rely on apparmor/SELinux profiles for security.
4. Stability over speed.
5. Use standard approaches to writing code wherever possible. No hacks without good reason.
