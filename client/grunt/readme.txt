GruntJS is a web development task runner that can handle many of the tedious jobs during development. For more information visit http://gruntjs.com/

This project is setup to use Karma and SASS. Other plugins can be found at the gruntjs website.

-- requirements

	-- Minimum
	1. NodeJs
		http://nodejs.org/
	2. GruntJs
		http://gruntjs.com/getting-started
		-- The Gruntfile comes setup for Sass (see SASS below)

	-- Karma
	1. Karma
		http://karma-runner.github.io/0.10/index.html
		> npm install -g karma

	-- SASS
	1. Ruby
		http://www.ruby-lang.org/en/downloads/
	2. Sass
		http://sass-lang.com/download.html

-- setup

	install the node_modules. this will create a node_modules folder with the dependencies addded automatically.
	> npm install

-- use

	the default tasks are Karma and Sass used via the watcher.
	> grunt