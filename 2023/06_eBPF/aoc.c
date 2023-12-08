// go:build ignore

#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

char __license[] SEC("license") = "Dual MIT/GPL";
const char aoc_path[] = "/tmp/lima/aoc";

struct
{
    __uint(type, BPF_MAP_TYPE_ARRAY);
    __type(key, __u32);
    __type(value, __u64);
    __uint(max_entries, 1);
} counting_map SEC(".maps");

struct
{
    __uint(type, BPF_MAP_TYPE_ARRAY);
    __type(key, __u32);
    __type(value, __u64);
    __uint(max_entries, 10);
} aoc_map SEC(".maps");
// This struct is defined according to the following format file:
// /sys/kernel/tracing/events/syscalls/sys_enter_openat/format
struct openat_info
{
    /* The first 8 bytes is not allowed to read */
    __u64 pad;

    __s32 __syscall_nr;
    __s32 pad2;
    __u64 dfd;
    const char *filename;
    int flags;
    __u64 mode;
};

// Square root of integer
__u64 int_sqrt(__u64 s)
{
    // Zero yields zero
    // One yields one
    if (s <= 1)
        return s;

    // Initial estimate (must be too high)
    __u64 x0 = s / 2;

    // Update
    __u64 x1 = (x0 + s / x0) / 2;

    __u32 safeguard = 0;
    while (x1 < x0 && safeguard < 300) // Bound check
    {
        x0 = x1;
        x1 = (x0 + s / x0) / 2;
        safeguard++;
    }
    return x0;
}

__u64 compute_ways(__u64 duration, __u64 record)
{
    record += 1; // we must be strictly greater than the last record
    __u64 h_min, h_max;
    __u64 h_min_100 = int_sqrt(10000 * (duration * duration - 4 * record));
    h_min_100 = 100 * duration - h_min_100;
    h_min_100 /= 2;
    if (h_min_100 % 100 > 0) // manually implement the ceiling function
    {
        h_min = h_min_100 / 100 + 1;
    }
    else
    {
        h_min = h_min_100 / 100;
    }
    h_max = (duration + int_sqrt(duration * duration - 4 * record)) / 2 + 1;
    bpf_printk("h_min: %8d h_max: %8d", h_min, h_max);
    return h_max - h_min;
}

// This tracepoint is defined in mm/page_alloc.c:__alloc_pages_nodemask()
// Userspace pathname: /sys/kernel/tracing/events/syscalls/sys_enter_openat
SEC("tracepoint/syscalls/sys_enter_openat")
int aoc06(struct openat_info *info)
{
    __u32 key = 0;

    char filename[256];
    bpf_probe_read_str(filename, 256, info->filename);
    if (bpf_strncmp(filename, sizeof(aoc_path) - 1, aoc_path) != 0)
    {
        return 0; // we only want to be run when a specific path is opened
    }
    __u64 *count = bpf_map_lookup_elem(&aoc_map, &key);
    if (!count)
    {
        return 0;
    }

    __u64 res1 = 1, res2 = 0;
    __u64 dur2 = 0;
    __u64 rec2 = 0;

    for (__u32 i = 0; i < *count && i < sizeof(aoc_map.max_entries); i++)
    {
        key++;
        __u64 *dur_rec = bpf_map_lookup_elem(&aoc_map, &key);
        if (!dur_rec)
        {
            return 0;
        }

        __u64 tuple = *dur_rec;
        __u64 duration = tuple >> 32;
        __u64 record = tuple & (((__u64)1 << 32) - 1);
        bpf_printk("duration: %8d record: %8d", duration, record);
        res1 *= compute_ways(duration, record);

        // parse part2 number
        __u64 sfgrd = 0, tmp = duration;
        while (tmp > 0 && sfgrd++ < 10)
        {
            dur2 *= 10;
            tmp /= 10;
        }
        dur2 += duration;

        sfgrd = 0;
        tmp = record;
        while (tmp > 0 && sfgrd++ < 10)
        {
            rec2 *= 10;
            tmp /= 10;
        }
        rec2 += record;

        bpf_printk("dur2: %llu rec2: %llu", dur2, rec2);
    }
    res2 = compute_ways(dur2, rec2);
    bpf_printk("res1: %llu res2: %llu", res1, res2);

    key++;
    bpf_map_update_elem(&aoc_map, &key, &res1, BPF_ANY);
    key++;
    bpf_map_update_elem(&aoc_map, &key, &res2, BPF_ANY);

    return 0;
}
