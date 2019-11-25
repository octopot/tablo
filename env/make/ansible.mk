_playbook = ansible-playbook -i env/ansible/.private.hosts
_recipes  = ubuntu docker service

.PHONY: ansible-playbook-apply
ansible-playbook-apply: recipes=$(foreach recipe,$(_recipes),env/ansible/$(recipe).yml)
ansible-playbook-apply:
	$(_playbook) $(recipes)

.PHONY: ansible-playbook-check
ansible-playbook-check: recipes=$(foreach recipe,$(_recipes),env/ansible/$(recipe).yml)
ansible-playbook-check:
	$(_playbook) --check $(recipes)

.PHONY: ansible-playbook-diff
ansible-playbook-diff: recipes=$(foreach recipe,$(_recipes),env/ansible/$(recipe).yml)
ansible-playbook-diff:
	$(_playbook) --diff $(recipes)


define recipe_tpl

.PHONY: ansible-playbook-for-$(1)
ansible-playbook-for-$(1):
	$$(eval _recipes = $(1))
	@echo recipe env/ansible/$$(_recipes).yml will be used

endef

render_recipe_tpl = $(eval $(call recipe_tpl,$(recipe)))
$(foreach recipe,$(_recipes),$(render_recipe_tpl))
