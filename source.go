// Copyright (c) 2026 Ty Bennett. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package main

func prepareLocalSource(path string) (absPath string, err error)

// Returns the cloned dir, a cleanup func, and any error
func cloneRepository(url, ref string) (dir string, cleanup func(), err error)
