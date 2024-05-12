# Watch My IP

Send out an alert when my IP changes.

The code uses [2ip.io](https://2ip.io) to discover the current external IP address.  If it changes an alert will be send using the [Pushover API](https://pushover.net).

The site to determine external IP is configurable.

Alerting is currently only done via the [Pushover API](https://pushover.net). Other notifiers may be available - please see [lib-core-go](https://github.com/skeletonkey/lib-core-go/tree/main/notify). If a specific one does not exist feel free to create an issue or submit a pull request with the desired notifier.

When the IP changes there is only 1 alert sent out without regards if there are any failures. Thoughts on how to improve this are welcome - please create a new issue for this project with the label `enhancement`.

## Environmental Variables

### Execution

#### PROJECT_CONFIG_FILE

This is the location of the config file for the program.  An example is provided with config_sample.json.

### Deployment

Makefile.deploy is really a bash script for deploying the program on to a server running Docker, builds and starts the container.

It is recommended to use [direnv](https://direnv.net/) (or something similar) to place these variable in to a local file that is not checked in.  Eventually, some type of Secret Management will be created and replace this.

#### BRANCH

The git branch that has the code that is being deployed.  This should usually be 'main'.

#### DOCKER_ARGS

Any additional arguments that need to be supplied to the `docker run` command.  For this program this is an empty string.

#### IMAGE_NAME

Name of the image that will be created and run. Please follow Docker standards for the name. `watchmyip` is recommended for this variable.

#### REMOTE_DIR

The directory on the remote server where the Dockerfile is placed.

#### REMOTE_SERVER

The server name/address which is running Docker. The program will be built and run on this server.

#### USER

The user that is used during ssh/scp communications with REMOTE_SERVER.