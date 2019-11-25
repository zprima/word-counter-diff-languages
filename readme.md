# Word counter in different programing languages


| Language          | Command                  | Time  | Word Count
| ----------------- |:------------------------:| -----:|------------:|
| C                 | time wc -w ./moby.txt    | 0.007 | 214112
| Go (not compiled) | time go run main.go      | 0.265 | 180009
| Go (compiled)     | time ./main              | 0.042 | 180009
| Node.js           | time node main.js        | 0.097 | 194168
| Ruby              | time ruby main.rb        | 0.424 | 165555
