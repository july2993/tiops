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

import "sync"

type Builder struct {
	tasks []Task
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Ping(host string) *Builder {
	b.tasks = append(b.tasks, Ping{host: host})
	return b
}

func (b *Builder) SSH(host, keypath string) *Builder {
	b.tasks = append(b.tasks, SSH{
		host:    host,
		keypath: keypath,
	})
	return b
}

func (b *Builder) CopyFile(src, dstHost, dstPath string) *Builder {
	b.tasks = append(b.tasks, &CopyFile{
		src:     src,
		dstHost: dstHost,
		dstPath: dstPath,
	})
	return b
}

func (b *Builder) RenderConfig(tmplPath, dstHost, dstPath string) *Builder {
	b.tasks = append(b.tasks, &RenderConfig{
		tmplPath: tmplPath,
		dstHost:  dstHost,
		dstPath:  dstPath,
	})
	return b
}

func (b *Builder) Wait(wg sync.WaitGroup) *Builder {
	b.tasks = append(b.tasks, Wait{wg: wg})
	return b
}

func (b *Builder) Parallel(tasks ...Task) *Builder {
	b.tasks = append(b.tasks, &parallelTask{subtasks: tasks})
	return b
}

func (b *Builder) Build() Task {
	return &serialTask{subtasks: b.tasks}
}
