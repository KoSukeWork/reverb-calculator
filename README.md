# reverb-calculator
Calculate delay times for reverb and delays efx by BPM.

## Usage
```
Usage of ./reverb-calc:
  -bars int
    	The amount of bars to display (default 4)
  -bpm float
    	Beats per Minute (default 120)
  -depth int
    	The lowest divider of one note to show. (default 11)
```
## Examples
```
$ ./reverb-calc -bpm 120 -bars 1 -depth 5
 Tempo:     120.00 BPM
   1/1:    2000.00 ms      2.000 s
   1/2:    1000.00 ms      1.000 s
   1/4:     500.00 ms      0.500 s
   1/8:     250.00 ms      0.250 s
  1/16:     125.00 ms      0.125 s

./reverb-calc -bpm 120 -bars 8 -depth 1
 Tempo:     120.00 BPM
   8/1:   16000.00 ms     16.000 s
   7/1:   14000.00 ms     14.000 s
   6/1:   12000.00 ms     12.000 s
   5/1:   10000.00 ms     10.000 s
   4/1:    8000.00 ms      8.000 s
   3/1:    6000.00 ms      6.000 s
   2/1:    4000.00 ms      4.000 s
   1/1:    2000.00 ms      2.000 s
```
