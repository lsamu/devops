package pipeline

import (
    "gopkg.in/yaml.v2"
    "log"
    "reflect"
)

type (
    //Pipeline 流水
    Pipeline struct {
        Workflow map[interface{}]interface{} //工作任务
    }
)

//Parse 格式化工作流水
func (p *Pipeline) Parse(data string) (err error) {
    err = yaml.Unmarshal([]byte(data), &p.Workflow)
    if err != nil {
        return err
    }
    return
}

//Process 处理工作
func (p *Pipeline) Process() {
    log.Println("开始任务。。。")
    workflow := p.Workflow
    if workflow["name"] != "" {
        log.Printf("任务名称:%v", workflow["name"])
    }
    on := workflow[true]
    //map[interface {}]interface {}
    //string
    //[]interface {}
    onType := reflect.TypeOf(on)
    log.Printf("仓库事件触发:%v,类型:%v", on, onType)
    if workflow["env"] != "" {
        log.Printf("全局环境变量:%v", workflow["env"])
    }
    jobs := workflow["jobs"]
    for jobKey, jobLoop := range jobs.(map[interface{}]interface{}) {
        log.Printf("任务步骤:%v", jobKey)
        job := jobLoop.(map[interface{}]interface{})
        jobName := job["name"]
        log.Printf("任务名称:%v", jobName)
        jobRunsOn := job["runs-on"]
        log.Printf("运行镜像:%v", jobRunsOn)
        jobNeeds := job["needs"]
        log.Printf("运行依赖:%v", jobNeeds)
        log.Println("****************************************************************************************************")
        for _, stepLoop := range job["steps"].([]interface{}) {
            step := stepLoop.(map[interface{}]interface{})
            stepName := step["name"]
            log.Printf("步骤名称:%v", stepName)
            stepEnv := step["env"]
            log.Printf("环境变量:%v", stepEnv)
            stepWith := step["with"]
            log.Printf("插件参数:%v", stepWith)
            stepUses := step["uses"]
            log.Printf("使用插件:%v", stepUses)
            stepRun := step["run"]
            log.Printf("步骤命令:%v", stepRun)
            log.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
        }
    }
}
