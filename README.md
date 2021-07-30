# utils
It collects some small useful functions comes from below projects, and remove some functions which I think are useless. 
Except that, it contains other useful help utility.

- [godash](https://github.com/alioygur/godash)
- [is](https://github.com/alioygur/is)

## alidts
help parsing aliyun DTS messages which come from kafka.

## parallel
help run some long-running process parallelling
```
import log 

var group parallel.Group
group.Add(
    func() error{
        // do something
    },
    func(error){
        // when receive error, gracefully clean and shutdown
    }
)
err := group.Run()
if err != nil {
    log.Printf("run stops, received error: %v", err)
}
```


