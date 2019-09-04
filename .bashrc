function git_branch {
  branch="`git branch 2>/dev/null | grep "^\*" | sed -e "s/^\*\ //"`"
  if [ "${branch}" != "" ];then
      if [ "${branch}" = "(no branch)" ];then
          branch="(`git rev-parse --short HEAD`...)"
      fi
      echo " ($branch)"
  fi
}

export PS1='\033[33m\]Orange\033[01\] @ \033[01;36m\]\W\[\033[01;32m\]$(git_branch)\[\033[00m\] \$ '

alias eshop_mysql_local='mysql -uroot -proot eshop'

# 登录开发机配置
alias r='~/bin/relay_auto_login/relay'

# 快读跳转
projectPath=/Users/yangchengzhi/yangchengzhi/go/src
alias goto='func() { cd ${projectPath}/$1;}; func'

alias run="go run main.go"


######################### tencent cloud ############
alias root_cloud="ssh root@188.131.132.143"
alias deployer_cloud="ssh deployer@188.131.132.143"


######################### local mysql ##############
alias eshop_mysql_local="mysql -uroot -proot eshop"
alias eshop_mysql_local_3307="mysql -uroot -proot -h'127.0.0.1' -P3307  eshop"
alias eshop_mysql_local_3308="mysql -uroot -proot -h'127.0.0.1' -P3308 eshop"

export CGO_ENABLED=1
export CGO_LDFLAGS_ALLOW=".*"

alias grep="grep --color"

alias multi_mysql='mysqld_multi --defaults-file=/usr/local/etc/my.cnf start'
alias eshop_vagrant="mysql -uroot -proot -h'eshop.com' eshop"
alias vagrant_start="cd /Users/yangchengzhi/Homestead && vagrant up && vagrant ssh"
alias pkillEshopapi="pkill eshop-api"
