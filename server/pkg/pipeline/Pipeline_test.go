package pipeline

import (
    "testing"
)

func TestPipeLine(t *testing.T) {
    data := `name: dotnetcore # workflow name

on: [push] # event trigger，什么事件触发 build

jobs:
  build:
    runs-on: ubuntu-latest # 指定 runner，使用 Github 提供的 runner

    steps:
    - uses: actions/checkout@v1 # checkout
    - name: Setup .NET Core # 设置 dotnet core 环境
      uses: actions/setup-dotnet@v1
      with:
        dotnet-version: 3.0.100
    - name: dotnet info # 输出 dotnet -info，查看 dotnet 版本信息
      env:
        MY_VAR: Hi there! My name is
        FIRST_NAME: Mona
        MIDDLE_NAME: The
        LAST_NAME: Octocat
      run: dotnet --info
    - name: build
      run: bash build.sh # 在 bash 中运行 build 脚本`
    //data = `name: Hexo Deploy
    //
    //# 只监听 source 分支的改动
    //on:
    // push:
    //   branches:
    //     - source
    //
    //# 自定义环境变量
    //env:
    // POST_ASSET_IMAGE_CDN: true
    //
    //jobs:
    // build-and-deploy:
    //   runs-on: ubuntu-latest
    //
    //   steps:
    //     # 获取博客源码和主题
    //     - name: Checkout
    //       uses: actions/checkout@v2
    //
    //     - name: Checkout theme repo
    //       uses: actions/checkout@v2
    //       with:
    //         repository: printempw/hexo-theme-murasaki
    //         ref: master
    //         path: themes/murasaki
    //
    //     # 这里用的是 Node.js 14.x
    //     - name: Set up Node.js
    //       uses: actions/setup-node@v1
    //       with:
    //         node-version: '14'
    //
    //     # 设置 yarn 缓存，npm 的话可以看 actions/cache@v2 的文档示例
    //     - name: Get yarn cache directory path
    //       id: yarn-cache-dir-path
    //       run: echo "::set-output name=dir::$(yarn cache dir)"
    //
    //     - name: Use yarn cache
    //       uses: actions/cache@v2
    //       id: yarn-cache
    //       with:
    //         path: ${{ steps.yarn-cache-dir-path.outputs.dir }}
    //         key: ${{ runner.os }}-yarn-${{ hashFiles('**/yarn.lock') }}
    //         restore-keys: |
    //           ${{ runner.os }}-yarn-
    //
    //     # 安装依赖
    //     - name: Install dependencies
    //       run: |
    //         yarn install --prefer-offline --frozen-lockfile
    //
    //     # 从之前设置的 secret 获取部署私钥
    //     - name: Set up environment
    //       env:
    //         DEPLOY_KEY: ${{ secrets.DEPLOY_KEY }}
    //       run: |
    //         sudo timedatectl set-timezone "Asia/Shanghai"
    //         mkdir -p ~/.ssh
    //         echo "$DEPLOY_KEY" > ~/.ssh/id_rsa
    //         chmod 600 ~/.ssh/id_rsa
    //         ssh-keyscan github.com >> ~/.ssh/known_hosts
    //
    //     # 生成并部署
    //     - name: Deploy
    //       run: |
    //         npx hexo deploy --generate`
       data = `name: Hello World
on: [push,pull]
jobs:
  my_first_job:
    name: My first job
    runs-on: ubuntu-latest
    steps:
    - name: checkout
      uses: actions/checkout@master
    - name: Run a single-line script
      run: echo "Hello World!"
  my_second_job:
    name: My second job
    runs-on: macos-latest
    steps:
    - name: Run a multi-line script
      env:
        MY_VAR: Hello World!
        MY_NAME: P3TERX
      run: |
        echo $MY_VAR
        echo My name is $MY_NAME`
//       data=`ame: Hexo Deploy
//
//# 只监听 source 分支的改动
//on:
//  push:
//    branches:
//      - source
//
//# 自定义环境变量
//env:
//  POST_ASSET_IMAGE_CDN: true
//
//jobs:
//  build-and-deploy:
//    runs-on: ubuntu-latest
//
//    steps:
//      # 获取博客源码和主题
//      - name: Checkout
//        uses: actions/checkout@v2
//
//      - name: Checkout theme repo
//        uses: actions/checkout@v2
//        with:
//          repository: printempw/hexo-theme-murasaki
//          ref: master
//          path: themes/murasaki
//
//      # 这里用的是 Node.js 14.x
//      - name: Set up Node.js
//        uses: actions/setup-node@v1
//        with:
//          node-version: '14'
//
//      # 设置 yarn 缓存，npm 的话可以看 actions/cache@v2 的文档示例
//      - name: Get yarn cache directory path
//        id: yarn-cache-dir-path
//        run: echo "::set-output name=dir::$(yarn cache dir)"
//
//      - name: Use yarn cache
//        uses: actions/cache@v2
//        id: yarn-cache
//        with:
//          path: ${{ steps.yarn-cache-dir-path.outputs.dir }}
//          key: ${{ runner.os }}-yarn-${{ hashFiles('**/yarn.lock') }}
//          restore-keys: |
//            ${{ runner.os }}-yarn-
//
//      # 安装依赖
//      - name: Install dependencies
//        run: |
//          yarn install --prefer-offline --frozen-lockfile
//
//      # 从之前设置的 secret 获取部署私钥
//      - name: Set up environment
//        env:
//          DEPLOY_KEY: ${{ secrets.DEPLOY_KEY }}
//        run: |
//          sudo timedatectl set-timezone "Asia/Shanghai"
//          mkdir -p ~/.ssh
//          echo "$DEPLOY_KEY" > ~/.ssh/id_rsa
//          chmod 600 ~/.ssh/id_rsa
//          ssh-keyscan github.com >> ~/.ssh/known_hosts
//
//      # 生成并部署
//      - name: Deploy
//        run: |
//          npx hexo deploy --generate`
//       data=`name: Gitee repos mirror periodic job
//on:
//# 如果需要PR触发把push前的#去掉
//# push:
//  schedule:
//    # 每天北京时间9点跑
//    - cron:  '0 1 * * *'
//jobs:
//  build:
//    runs-on: ubuntu-latest
//    steps:
//    - name: Mirror the Github organization repos to Gitee.
//      uses: Yikun/gitee-mirror-action@v0.01
//      with:
//        # 必选，需要同步的Github用户（源）
//        src: github/Yikun
//        # 必选，需要同步到的Gitee的用户（目的）
//        dst: gitee/yikunkero
//        # 必选，Gitee公钥对应的私钥，https://gitee.com/profile/sshkeys
//        dst_key: ${{ secrets.GITEE_PRIVATE_KEY }}
//        # 必选，Gitee对应的用于创建仓库的token，https://gitee.com/profile/personal_access_tokens
//        dst_token:  ${{ secrets.GITEE_TOKEN }}
//        # 如果是组织，指定组织即可，默认为用户user
//        # account_type: org
//        # 还有黑、白名单，静态名单机制，可以用于更新某些指定库
//        # static_list: repo_name
//        # black_list: 'repo_name,repo_name2'
//        # white_list: 'repo_name,repo_name2'`
    pipeline := Pipeline{}
    err := pipeline.Parse(data)
    if err != nil {
       t.Errorf("%+v", err)
       return
    }
    pipeline.Process()
    //var err error
    //m := make(map[interface{}]interface{})
    //err = yaml.Unmarshal([]byte(data), &m)
    //if err != nil {
    //   log.Fatalf("error: %v", err)
    //}
    //fmt.Printf("--- m:\n%v\n\n", m)
    /*
       pipeline := Pipeline{}
       err = yaml.Unmarshal([]byte(data), &pipeline)
       if err != nil {
           log.Fatalf("error: %v", err)
       }
       fmt.Printf("%+v\n", pipeline)
       fmt.Printf("****************************************************************************\n")
       fmt.Printf("%+v\n", pipeline.Name)
       fmt.Printf("****************************************************************************\n")
       fmt.Printf("%+v\n", pipeline.On)
       fmt.Printf("****************************************************************************\n")
       fmt.Println(len(pipeline.Jobs))
       for _, job := range pipeline.Jobs {
           //oo := job.(Job)
           //fmt.Println(oo)
           jj := job.(map[interface{}]interface{})
           fmt.Printf("%+v\n", jj["name"])
           fmt.Printf("%+v\n", jj["runs-on"])
           fmt.Printf("%+v\n", jj["steps"])
           //for _ , step := range jj["steps"].([]map[interface{}]interface{}) {
           //    fmt.Printf("%+v\n", step["name"])
           //}
           fmt.Printf("****************************************************************************\n")
       }*/
}
