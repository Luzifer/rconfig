# 2.6.0 / 2025-04-13

  * Update pflag, require Go 1.23
  * Lint: Update linter config, disable struct-tag on validate test
  * CI: Pin checkout action
  * CI: Switch to Github hosted image

# 2.5.2 / 2024-08-27

  * Fix: Add support for int16 flags
  * Fix: Ensure int/uint do not overflow
  * Lint: Replace deprecated linter

# 2.5.1 / 2024-08-18

  * CI: Update linter config, fix linter errors, update CI config

# 2.5.0 / 2023-12-20

  * Update deps, fix linter errors, bump min Go version
  * [CI] Replace old CI with Github Actions

# 2.4.0 / 2021-09-06

  * Switch to repo-runner for tests
  * Port tests to pure-Go-tests

# 2.3.0 / 2021-08-06

  * Drop pre-1.15 support, update dependencies
  * Add hint for go modules to v2

# 2.2.1 / 2019-02-04

  * Add go module information

# 2.2.0 / 2018-09-18

  * Add support for time.Time flags

# 2.1.0 / 2018-08-02

  * Add AutoEnv feature

# 2.0.0 / 2018-08-02

  * Breaking: Ensure an empty default string does not yield a slice with 1 element  
    Though this is a just a tiny change it does change the default behaviour, so I'm marking this as a breaking change. You should ensure your code is fine with the changes.

# 1.2.0 / 2017-06-19

  * Add ParseAndValidate method

# 1.1.0 / 2016-06-28

  * Support time.Duration config parameters
  * Added goreportcard badge
  * Added testcase for using bool with ENV and default
