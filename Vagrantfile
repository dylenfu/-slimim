Vagrant.configure("2") do |config|
	config.vm.box = "centos7"
	config.ssh.username = "vagrant"
	config.ssh.password = "vagrant"
	config.vm.synced_folder ".", "/home/vagrant/opt/golang/src/slimim"
	config.vm.network "forwarded_port", guest: 22, host: 2200 # SSH port
end
