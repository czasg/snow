# snow
[![LICENSE](https://img.shields.io/github/license/mashape/apistatus.svg?style=flat-square&label=License)](https://github.com/czasg/snow/blob/master/LICENSE)
[![codecov](https://codecov.io/gh/czasg/snow/branch/main/graph/badge.svg?token=J7Y4K906p6)](https://codecov.io/gh/czasg/snow)
[![GitHub Stars](https://img.shields.io/github/stars/czasg/snow.svg?style=flat-square&label=Stars&logo=github)](https://github.com/czasg/snow/stargazers)

Snowflake by Go.

```go
package main

import "github.com/czasg/snow"

func main() {
    // default WorkID & DataCenterID is 0.
    snowID := snow.Next()
    
    // self-define WorkID / DataCenterID
    s := snow.Snow{
        WorkID:       1,
        DataCenterID: 2,
    }
    snowID = s.Next()
}
```
