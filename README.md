# getSCSlotData
1.Get single slot data

 ./main --address="0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" --slot=1

 Example
```bash
we have a connection
SC address
0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
signal Slot provided, get slot 1 of the address 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
[87 69 84 72 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 8]
```

2.Get continuous slot data, from loeslot to the highslot data. The default lowslot is 0.

 ./main --address="0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" --highslot=1

 Example
```bash
we have a connection
SC address
0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
highSlot provided, get slot 0 of the address 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
[87 114 97 112 112 101 100 32 69 116 104 101 114 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 26]
highSlot provided, get slot 1 of the address 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
[87 69 84 72 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 8]
``` 

3. Get continuous slot data, from loeslot to the highslot data with target lowslot.

./main --address="0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" --highslot=2 --lowslot=1

 Example
```bash
we have a connection
SC address
0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
highSlot provided, get slot 1 of the address 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
[87 69 84 72 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 8]
highSlot provided, get slot 2 of the address 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 18]
``` 

4. get specific slot data by providing string slot.

 ./main --address="0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" --arrayslot="0 1 2"

 Example
```bash
we have a connection
SC address
0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
arraySlot provided, get slot 0 of the address 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
[87 114 97 112 112 101 100 32 69 116 104 101 114 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 26]
arraySlot provided, get slot 1 of the address 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
[87 69 84 72 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 8]
arraySlot provided, get slot 2 of the address 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 18]
```  