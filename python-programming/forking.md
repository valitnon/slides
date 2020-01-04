# Forking and Multiprocess
{id: forking-processes}

## Performance Monitoring
{id: performance-monitoring}

* Linux, OSX: htop
* Windows: Performance Monitor


## Forking
{id: forking}
{i: fork}
{i: wait}
![](examples/os/fork.py)

```
$ python fork.py
In Parent (common) The child is: 31351
In Parent PID: 31350 PPID: 10025
In Child of common
In Child PID: 31351 PPID: 31350
(31351, 768)
```


## Fork skeleton
{id: fork-skeleton}
![](examples/os/fork_skeleton.py)


## Fork with load
{id: fork-load}
![](examples/os/fork_load.py)


## Fork load results
{id: fork-load-results}

```
$ time python fork_load.py 1
In Parent of 96355
In Child
(96355, 768)

real    0m26.391s
user    0m25.893s
sys 0m0.190s
```

```
$ time python fork_load.py 8
In Parent of 96372
In Parent of 96373
In Parent of 96374
In Child
In Child
In Parent of 96375
In Child
In Child
In Parent of 96376
In Child
In Parent of 96377
In Child
In Child
In Parent of 96378
In Parent of 96379
In Child
(96374, 768)
(96372, 768)
(96375, 768)
(96373, 768)
(96376, 768)
(96377, 768)
(96378, 768)
(96379, 768)

real    0m12.754s
user    0m45.196s
sys 0m0.164s
```


## Marshalling / Serialization
{id: marshalling}


Marshalling (or serialization) is the operation when we take an arbitrary
data structure and convert it into a string in a way that we can convert
the string back to the same data structure.




Marshalling can be used to save data persistent between execution of the same
script, to transfer data between processes, or even between machines.
In some cases it can be used to communicate between two processes written in
different programming languages.




The <a href="http://docs.python.org/library/marshal.html">marshal</a> module
provides such features but it is not recommended as it was built
for internal object serialization for python.




The <a href="http://docs.python.org/library/pickle.html">pickle</a> module was designed for this task.




The json module can be used too.





## Fork
{id: fork}
{i: fork}
{i: getpid}
{i: wait}
![](examples/advanced/fork.py)


## Fork with random
{id: fork-with-random}


When the <emp>random</emp> module is loaded it automatically calls <emp>random.seed()</emp> to initialize the
random generator. When we create a fork this is not called again and thus all the processes
will return the same random numbers. We can fix this by calling <emp>random.seed()</emp>
manually.


![](examples/advanced/forkrand.py)


## Multiprocess file: Pool
{id: multiprocess-file}
{i: multiprocess}
{i: Pool}
![](examples/advanced/multiprocess_files.py)

```
$ python multiprocess_files.py 3 multiprocess_*

Process 22688 analyzing multiprocess_files.py
Process 22689 analyzing multiprocess_load.py
Process 22690 analyzing multiprocess_pool_async.py
Process 22688 analyzing multiprocess_pool.py
{'filename': 'multiprocess_files.py', 'total': 833, 'digits': 10, 'spaces': 275}
{'filename': 'multiprocess_load.py', 'total': 694, 'digits': 14, 'spaces': 163}
{'filename': 'multiprocess_pool_async.py', 'total': 695, 'digits': 8, 'spaces': 161}
{'filename': 'multiprocess_pool.py', 'total': 397, 'digits': 3, 'spaces': 80}
```

{aside}

We asked it to use 3 processes, so looking at the process ID you can see one of them worked twice.
The returned results can be any Python datastructure. A dictionary is usually a good idea.
{/aside}


## Multiprocess load
{id: multiprocess-load}
![](examples/advanced/multiprocess_load.py)


## Multiprocess: Pool
{id: multiprocess-pool}
{i: multiprocess}
{i: Pool}

<b>Pool(3)</b> creates 3 child-processes and let's them compute the values. <b>map</b>
returns the results in the same order as the input came in.

![](examples/advanced/multiprocess_pool.py)

```
python multiprocess_pool.py  11 3
python multiprocess_pool.py  100 5
```



## Multiprocess load async
{id: multiprocess-load-async}
![](examples/advanced/multiprocess_pool_async.py)




## Exercise: fork return data
{id: exercise-fork-return-data}


Allow the chile process to return data to the parent process. Before exiting from the child process, serialize the data-structure you want to send back and save
in a file that corresponds to the parent process and the child process. (eg. created from the PID of the paraent process and the PID of the child process)
In the parent process, when one of the children exits, check if there is a file corresponding to this child process, read the file and de-serialize it.








