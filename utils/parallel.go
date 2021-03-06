package utils;

import (
    "runtime"
);

type empty struct{};
type semaphore chan empty;

func For(begin, end, step uint, f func(uint)) {
    cpus := uint(runtime.GOMAXPROCS(0));
    sem := make(semaphore, cpus);

    for i := uint(0); i < cpus; i++ {
        go func(sem semaphore, cpus, begin, end, step uint, f func(uint)) {
            for i := begin; i < end; i += (cpus * step) {
                f(i);
            }
            sem <- empty{};
        }(sem, cpus, begin+(i*step), end, step, f);
    }

    for i := uint(0); i < cpus; i++ {
        <- sem;
    }
};
