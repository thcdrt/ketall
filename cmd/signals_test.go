// +build !windows

/*
Copyright 2019 Cornelius Weig

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"context"
	"sync"
	"syscall"
	"testing"
)

func TestCatchCtrlC(t *testing.T) {
	var group sync.WaitGroup
	group.Add(1)

	ctx, cancel := context.WithCancel(context.Background())
	catchedSignal := syscall.SIGINT
	catchSigs(cancel, catchedSignal)

	go func() {
		<-ctx.Done()
		group.Done()
	}()

	syscall.Kill(syscall.Getpid(), catchedSignal)
	group.Wait()
}
