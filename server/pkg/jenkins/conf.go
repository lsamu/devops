package jenkins

import (
    "context"
    "github.com/bndr/gojenkins"
)

func InitJenkins()  {
    ctx := context.Background()
    jenkins := gojenkins.CreateJenkins(nil, "http://localhost:8080/", "admin", "admin")
    // Provide CA certificate if server is using self-signed certificate
    // caCert, _ := ioutil.ReadFile("/tmp/ca.crt")
    // jenkins.Requester.CACert = caCert
    _, err := jenkins.Init(ctx)


    if err != nil {
        panic("Something Went Wrong")
    }

    jenkins.GetJob(ctx, "#jobname")
    //queueid, err := jenkins.InvokeSimple(ctx, params)
    //if err != nil {
    //    panic(err)
    //}
    //build, err := jenkins.GetBuildFromQueueID(ctx, job, queueid)
    //if err != nil {
    //    panic(err)
    //}

    // Wait for build to finish
    //for build.IsRunning(ctx) {
    //    time.Sleep(5000 * time.Millisecond)
    //    build.Poll(ctx)
    //}
    //
    //fmt.Printf("build number %d with result: %v\n", build.GetBuildNumber(), build.GetResult())
}

