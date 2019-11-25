_playbook = ansible-playbook -i env/ansible/.private.hosts
_recipes  = env/ansible/ubuntu.yml env/ansible/docker.yml env/ansible/service.yml

.PHONY: ansible-playbook-apply
ansible-playbook-apply:
	@$(_playbook) $(_recipes)

.PHONY: ansible-playbook-check
ansible-playbook-check:
	@$(_playbook) --check $(_recipes)

.PHONY: ansible-playbook-diff
ansible-playbook-diff:
	@$(_playbook) --diff $(_recipes)


.PHONY: ansible-playbook-for-ubuntu
ansible-playbook-for-ubuntu:
	$(eval _recipes = env/ansible/ubuntu.yml)
	@echo recipe $(_recipes) will be used

.PHONY: ansible-playbook-for-docker
ansible-playbook-for-docker:
	$(eval _recipes = env/ansible/docker.yml)
	@echo recipe $(_recipes) will be used

.PHONY: ansible-playbook-for-service
ansible-playbook-for-service:
	$(eval _recipes = env/ansible/service.yml)
	@echo recipe $(_recipes) will be used
