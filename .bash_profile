alias ll="ls -l"
alias ip='ifconfig | grep --color "inet 172"'

[ -r ~/.bashrc ] && source ~/.bashrc
export CLICOLOR=1
export LSCOLORS=GxFxCxDxBxegedabagaced


export PATH=/usr/local/bin:$PATH

code () { VSCODE_CWD="$PWD" open -n -b "com.microsoft.VSCode" --args $* ;}
export PATH="/usr/local/opt/openssl/bin:$PATH"
alias mnmp='/Users/yangchengzhi/yangchengzhi/mnmp/mnmp.sh'
export PATH="/usr/local/opt/mysql@5.7/bin:$PATH"
alias mnmp='/Users/yangchengzhi/yangchengzhi/tools/mnmp//mnmp.sh'

############## golang ############
export GOROOT=/usr/local/Cellar/go/1.10.3/libexec
export GOPATH=$HOME/yangchengzhi/go:$HOME/go:$HOME/yangchengzhi/test
export GOBIN=$HOME/yangchengzhi/go/bin #go源码编译后保存可执行程序的目录。
export PATH=$PATH:$GOROOT/bin:$HOME/yangchengzhi/go/bin:$HOME/go/bin




############### JAVA #############
export JAVA_HOME=$(/usr/libexec/java_home)

export PATH=${JAVA_HOME}/bin:$PATH


############# composer ###########
laravel=$HOME/.composer/vendor/bin
export PATH=$laravel:$PATH

############ vagrant ###############
alias vagrant_run="vagrant ssh 13cd897"

