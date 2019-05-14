#How to Contribute

#参与流程

### Git操作流程

**初始化**：

1. 首先fork项目到自己的 GitHub 上
2. clone 自己的项目到本地
3. 创建新分支 `git branch editing`  作为你自己的开发分支
4. 把主仓库添加为远端库  `git remote add upstream git@github.com:1024hub/1024-forum.git`

此时拥有两个项目，一个是自己账户下的项目 `xxx(A)/1024-forum`，称为 A；一个是公共账户下的项目 `1024hub/1024-forum` ，称为 B。

想参与贡献需要在自己的项目 A 下更新修改，然后提交 pull request 到公共项目 B。

你主要操作的分支是自己创建的分支  `editing`，其他分支 dev 、master 只用来同步公共项目的最新的代码，保持最新的代码状态。

###如何同步公共项目的最新代码？



```g
git remote update 
git fetch upstream master dev   #这是拉取B中 master 和 dev 分支的最新代码
git rebase upstream/master master  #将最新代码与你本地代码master分支进行合并
git rebase upstream/dev dev  #将最新代码与你本地代码master分支进行合并
# ... 以及其他你要同步的分支，每个分支操作一遍
# 更新自己项目中的代码
git push origin master 
git push origin dev
# 自己的代码仓库保持了和公共仓库代码相同了
```

---

还有一种方法，让自己的这个分支直接跟踪公共仓库的代码

```
# 切换分支
git checkout dev
git branch -u upstream/dev
git branch -u upstream/master
# 这个时候，你 pull 的代码就是公共项目的最新代码
# 当你想更新自己仓库 master 的代码时，你可以 git push origin dev
```

### 如何将最新代码合并到自己的分支中？

```
# 默认自己的操作分支名是 editing
# 先保证自己分支干净，无未提交的文件，防止冲突
git checkout editing
git rebase master # or dev 将最新的代码应用到自己的分支中
```

### 如何提交 pull request

```
# 首先将最新的代码应用到自己的分支中  参照 #如何将最新代码合并到自己的分支中？
# push 到自己的仓库 A
git push origin editing
# 然后通过页面操作，提交 pull request 到公共项目
```