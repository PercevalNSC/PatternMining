# PatternMining using Go language

This code is developed for the lecture of data sience in UEC. 

## Transaction Data
In this package, transaction data is defined as static struct, a list of transaction. One transaction(tupple) contains ID (int) and transaction([]string). There is example transaction data under this paragraph.


### Example TransactionData
|ID( int ) | transaction ( []string ) |
|:-:|---|
|10 | "A", "C", "D" |
|20 | "A", "B", "C" |
|30 | "A", "B", "C", "E" |
|40 | "B", "E" |

## Use
~~~
min_support := 0.5

trasaction_data := sampleTrasactionData()

fmt.Println(trasaction_data.pickupFrequencyItemset(min_support)) //output frequency itemset: {[{[A]} {[B]} {[C]} {[E]} {[A B]} {[A C]} {[B C]} {[B E]} {[A B C]}]}
~~~

