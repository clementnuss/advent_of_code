package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go aoc aoc.c

func main() {

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

	// Load the compiled eBPF ELF and load it into the kernel.
	var objs aocObjects
	if err := loadAocObjects(&objs, nil); err != nil {
		log.Fatal("Loading eBPF objects:", err)
	}
	defer objs.Close()

	var aocKey uint32 = 0
	_ = objs.AocMap.Update(aocKey, uint64(len(duration)), ebpf.UpdateAny) // set the 0-th map value to the count
	for i := 0; i < len(duration); i++ {
		dur, _ := strconv.ParseUint(duration[i], 10, 64)
		rec, _ := strconv.ParseUint(record[i], 10, 64)

		aocKey++
		_ = objs.AocMap.Update(aocKey, dur<<32|rec, ebpf.UpdateAny) // set the 0-th map value to the count
	}

	kp, err := link.Tracepoint("syscalls", "sys_enter_openat", objs.Aoc06, nil)
	if err != nil {
		log.Fatalf("opening tracepoint: %s", err)
	}
	defer kp.Close()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	log.Println("Waiting for events..")
	for range ticker.C {
		var res1, res2 uint64
		if err := objs.AocMap.Lookup(aocKey+1, &res1); err != nil {
			log.Fatalf("reading map: %v", err)
		}
		if err := objs.AocMap.Lookup(aocKey+2, &res2); err != nil {
			log.Fatalf("reading map: %v", err)
		}
		log.Printf("res1 %v res2 %v", res1, res2)
	}
}
