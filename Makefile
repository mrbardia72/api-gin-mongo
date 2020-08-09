LOGFILE=$(LOGPATH) `date +'%A-%b-%d-%Y-%H-%M-%S'`

.PHONY: help
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help
$(info $v run makefile.....)

.PHONY: commit
commit: ## commit to github exp. run--- make commit
	git add .
	git commit -m "api-mongo-gin-${LOGFILE}"
	git push -u origin master 


.PHONY: git-push
git-push: ## push to github exp. run--- make git-push u="https://github.com/your-username/your-repository.git"
	git init
	git add .
	git commit -m "api-mongo-gin-${LOGFILE}"
	git remote add origin $u
	git pull --rebase origin master
	git push origin master 