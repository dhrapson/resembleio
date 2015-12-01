# Setting up a vagrant box

Travis CI is used to build resembleio & there might be occasion when, due to a failing build, the developer needs to exactly replicate the Travis build environment. Travis uses Ubuntu Precise at time of writing, which might be different from your own OS.
The most convenient way to do this is probably to use a Vagrant VM.
A high level outline of the setup steps is as below.

## Geting a vagrant VM
1. download vagrant
1. download virtualbox
1. In your user home dir: `vagrant init hashicorp/precise32 # This creates a 'Vagrantfile' in your home directory `
1. `vagrant up`      # this starts the VM
1. `vagrant list`    # this should show you have a single VM named 'default' in the running state
1. `vagrant ssh default`

## Installing the required binaries
1. `sudo apt-get install git`
1. `git config --global user.name "YourName"`
1. `git config --global user.email "YourEmail"`
1. `sudo apt-get install curl`
1. Download go from https://golang.org/doc/install?download=go1.5.1.linux-386.tar.gz
1. On your desktop: `vagrant scp go1.5.1.linux-386.tar.gz default:/home/vagrant/go1.5.1.linux-386.tar.gz`
1. On the vagrant VM: `sudo tar -C /usr/local -xzvf go1.5.1.linux-386.tar.gz`
1. Add to ~/.profile: `export PATH=$PATH:/usr/local/go/bin`
1. Add to ~/.profile: `export GOPATH=$HOME/gopath`
1. Add to ~/.profile: `export PATH=$PATH:$GOPATH/bin`
1. `source ~/.profile`
1. `command curl -sSL https://rvm.io/mpapis.asc | gpg --import -`
1. `\curl -L https://get.rvm.io |    bash -s stable --ruby --autolibs=enable --auto-dotfiles`
1. `rvm install 2.1.7 		# should install ruby v2.1.7`
1. `rvm list 						# should show ruby2.1.7 under the list of installed ruby versions`
1. `gem install bundler`

## Setup working area
1. `cd`
1. `mkdir gopath`
1. `go get github.com/dhrapson/resembleio/resemble`
1. `cd gopath/src/github.com/dhrapson/resembleio`
1. `go get -t -v ./...`
1. `go test -v ./...`
1. `chmod +x run_specs.sh`
1. `./run_specs.sh`

# When finished
1. vagrant suspend default
