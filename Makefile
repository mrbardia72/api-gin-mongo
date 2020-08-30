LOGFILE=$(LOGPATH) `date +'%A-%b-%d-%Y-%H-%M-%S'`
dir_form = forms/
main_form = package forms \n \n\n type name struct { \n\n\n }
dir_model = models/
main_model = package models\n\nimport ('gopkg.in/mgo.v2/bson') \n\ntype name struct { \n\n\n }
dir_helper = helpers/
main_helper = package helpers
dir_services = services/
main_services = package services
controllers= controllers/auth
.PHONY: help
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help
$(info $v run makefile.....)

.PHONY: commit
commit: ## commit to github exp. run---> make commit
	git add .
	git commit -m "api-mongo-gin-${LOGFILE}"
	git push


.PHONY: push
push: ## push to github exp. run---> make push u="https://github.com/your-username/your-repository.git"
	git init
	git add .
	git commit -m "api-mongo-gin-${LOGFILE}"
	git remote add origin $u
	git pull --rebase origin master
	git push origin master

.PHONY: form
form: ## create a new form ---> make form a="your-name-form"
	@mkdir -vp '$(dir_form)'
	echo "$(main_form)" >> $(dir_form)/$a.go

.PHONY: model
model: ## create a new model ---> make model a="your-name-model"
	@mkdir -vp '$(dir_model)'
	echo "$(main_model)" >> $(dir_model)/$a.go

.PHONY: helper
helper: ## create a new helper ---> make helper a="your-name-helper"
	@mkdir -vp '$(dir_helper)'
	echo "$(main_helper)" >> $(dir_helper)/$a.go

.PHONY: services
services: ## create a new services ---> make services a="your-name-services"
	@mkdir -vp '$(dir_services)'
	echo "$(main_services)" >> $(dir_services)/$a.go

.PHONY: controller
controller: ## create a new controller ---> make controller d="dir-name" a="your-name-controller"
	@mkdir -pv controllers/$d
	echo "package $d" >> controllers/$d/$a.go

.PHONY: repo
repo: ## create a new repo ---> make repo d="dir-name" a="your-name-repository"
	@mkdir -pv repository/$d
	echo "package $d" >> repository/$d/$a.go