if [ -n "$BASH_VERSION" ]; then
    if [ -f "$HOME/.bashrc" ]; then
    .   "$HOME/.bashrc"
    fi
fi

if [ -d "$HOME/bin" ]; then
    PATH="$HOME/bin:$PATH"
fi

export PATH=$PATH:/usr/local/go/bin

export GOOGLE_APPLICATION_CREDENTIALS="/Users/devtry/Workspace/Golang/golang-firebase-62a8c-firebase-adminsdk-kpvxz-0b68d0c284.json"
