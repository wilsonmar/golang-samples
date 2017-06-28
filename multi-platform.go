// multi-platform.go
// CAUTION: UNDER CONSTRUCTION: This does not currently work.
// From https://www.youtube.com/watch?v=hsgkdMrEJPs&t=15m40s

package main

func main() {

main_unuxlike.go:

    // Build directive: foo_${GOOS}.go
    // +build linux darwin netbsd openbsd freebsd
    // command means AND.

    func GetCommand() []string {
        return []string{"ifoncfig","-a"}
    }

main_windows.go:

    func GetCommand() []string {
        return []string{"iponcfig","/all"}
    }

main.go:

    cmd := GetCommand()
    out, err := exec.Command(cmd[0], cmd[1:]).Output()

} // main
