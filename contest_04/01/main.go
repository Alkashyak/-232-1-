package main
import (
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
)
func main() {
    bj, err := ioutil.ReadFile("data.json")
    if err != nil {
        log.Fatal(err)
    }
    count := 0
    var num int
    fmt.Scan(&num)
    var data []map[string]interface{}
    if err := json.Unmarshal(bj, &data); err != nil {
        log.Fatal(err)
    }
    for _, pr_i := range data {
        tasks, ok := pr_i["tasks"].([]interface{})
        if ok {
            for _, tsk := range tasks {
                task, ok := tsk.(map[string]interface{})
                if ok {
                    user_id, ok := task["user_id"].(float64)
                    if ok && int(user_id) == num {
                        completed, ok := task["completed"].(bool)
                        if ok && completed {
                            count++
                        }
                    }
                }
            }
        }
    }
    fmt.Print(count)
}





