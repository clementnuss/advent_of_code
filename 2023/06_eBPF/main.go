package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go aoc aoc.c

func main() {
	log.Println("Starting Advent Of Code - day 06 - eBPF solution")

	inputBytes, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	spl := strings.Split(string(inputBytes), "\n")
	duration := regexp.MustCompile(`\d+`).FindAllString(spl[0], -1)
	record := regexp.MustCompile(`\d+`).FindAllString(spl[1], -1)

	// Remove resource limits for kernels <5.11.
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal("Removing memlock:", err)
	}

	fmt.Println("Loading the eBPF objects into the kernel")
	// Load the compiled eBPF ELF and load it into the kernel.
	var objs aocObjects
	if err := loadAocObjects(&objs, nil); err != nil {
		log.Fatal("Loading eBPF objects:", err)
	}
	defer objs.Close()

	fmt.Println("Updating the eBPF map with the race parameters (the input)")
	var aocKey uint32 = 0
	_ = objs.AocMap.Update(aocKey, uint64(len(duration)), ebpf.UpdateAny) // set the 0-th map value to the count
	for i := 0; i < len(duration); i++ {
		dur, _ := strconv.ParseUint(duration[i], 10, 64)
		rec, _ := strconv.ParseUint(record[i], 10, 64)
		aocKey++
		_ = objs.AocMap.Update(aocKey, dur<<32|rec, ebpf.UpdateAny) // set the 0-th map value to the count
	}
	dbgFile, _ := os.Open("/sys/kernel/debug/tracing/trace_pipe")

	fmt.Println("Attaching the eBPF program to the sys_enter_openat tracepoint")
	kp, err := link.Tracepoint("syscalls", "sys_enter_openat", objs.Aoc06, nil)
	if err != nil {
		log.Fatalf("opening tracepoint: %s", err)
	}
	defer kp.Close()

	fmt.Println("triggering the eBPF program")
	_, _ = os.ReadFile("/tmp/lima/aoc06/trigger")

	var res1, res2 uint64
	if err := objs.AocMap.Lookup(aocKey+1, &res1); err != nil {
		log.Fatalf("reading map: %v", err)
	}
	if err := objs.AocMap.Lookup(aocKey+2, &res2); err != nil {
		log.Fatalf("reading map: %v", err)
	}
	fmt.Printf("res1 %v res2 %v\n", res1, res2)

	fmt.Println("Debug info (/sys/kernel/debug/tracing/trace_pipe):")
	dbg := make([]byte, 16*0x400) // allocate 16KB
	_, _ = dbgFile.Read(dbg)
	fmt.Print(string(dbg))

}
