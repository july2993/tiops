// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package task

import (
	"io"
	"sync"
)

type Task interface {
	Execute(ctx *Context) error
	Rollback(ctx *Context) error
}

type Context struct {
	SSHConnection io.ReadWriter
}

type serialTask struct {
	subtasks []Task
}

func (s *serialTask) Execute(ctx *Context) error {
	panic("implement me")
}

func (s *serialTask) Rollback(ctx *Context) error {
	panic("implement me")
}

type parallelTask struct {
	subtasks []Task
}

func (pt *parallelTask) Execute(ctx *Context) error {
	wg := sync.WaitGroup{}
	for _, t := range pt.subtasks {
		wg.Add(1)
		go func() {
			// TODO: error handling
			_ = t.Execute(ctx)
		}()
	}
	wg.Wait()
	return nil
}

func (pt *parallelTask) Rollback(ctx *Context) error {
	wg := sync.WaitGroup{}
	for _, t := range pt.subtasks {
		wg.Add(1)
		go func() {
			// TODO: error handling
			_ = t.Rollback(ctx)
		}()
	}
	wg.Wait()
	return nil
}
