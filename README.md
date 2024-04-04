# CMU Storage Engine

## Motive
This project serves as a practical application of the knowledge I gained from the "Intro To Database Systems - CMU" course.

## Components

### Buffer Pool Manager

The Buffer Pool Manager plays a crucial role in managing the memory used for caching data pages. It ensures efficient utilization of memory resources and optimizes data retrieval operations.

### Replacer

The Replacer component is responsible for managing the replacement strategy within the buffer pool. It determines which pages should be evicted from memory when additional space is required, I chose to use the LRU-K algorithm taking into consideration both past access timestamps and the frequency of pages.

### Disk Manager

The Disk Manager facilitates interactions between the buffer pool and the disk, It manages the Directory page, Row Pages, and headers stored in the disk.

### Disk Scheduler

The Disk Scheduler optimizes the order of disk operations to minimize seek times and enhance overall disk I/O performance. It aims to efficiently schedule disk access requests to reduce latency.

## Pages Layout

### Directory Page

I changed the design of the directory page from EXTENDIBLE HASH INDEX to a B+ Tree, which compressed storage and allowed for range searches.

### Row Pages

Row pages store actual data records within the database. They organize data in a format suitable for efficient retrieval and manipulation operations.
