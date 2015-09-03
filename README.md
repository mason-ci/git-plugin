git tool for Mason CI
======
This tool enables Mason CI to use git repositories

Capabilities
------------
* Poll repository
* Remote trigger
* Checkout
* Tag version
* Merge

Settings
--------
| Name             | Value                                 |
| ---------------- | ------------------------------------- |
| url              | (Optional) The url to the repository  |
| branches.build   | A list of branches to build           |
| branches.ignore  | A list of branches to igore           |

Variables
---------
These variables are exported and can be used by other tools.

| Name    | Value               |
| ------- | ------------------- |
| branch  | The current branch  |
