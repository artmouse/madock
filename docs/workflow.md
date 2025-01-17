# Workflow

The following guide shows you the normal development workflow using madock.

IMPORTANT: After changing any option in the following files, you should run `madock rebuild`
madock/projects/config.txt
madock/projects/{project name}/env.txt

#### 1. Start containers

```
madock start
```

#### 2. Composer commands

```
madock composer <command>
```

#### 3. Magento commands

```
madock magento <command>
```

#### 4. Working on frontend

```
madock node <command>
madock node grunt exec:<theme>
madock node grunt watch
```

**IMPORTANT:** For the Chrome browser, you can download the LiveReload plugin specifically for madock from the link [Google Chrome plugin](https://chrome.google.com/webstore/detail/livereload-for-madock/cmablbpbnbbgmakinefjgmgpolfahdbo). Then install it and enable it for the site you need.

**NOTE:** You might also need to disable your browser cache. For example in Chrome:

* `Open inspector > Settings > Network > Disable cache (while DevTools is open)`

#### 5. xdebug

* Enable xdebug

  ```
  madock debug on
  ```

* Configure xdebug in PHPStorm (Only first time)

    * [PHPStorm + Xdebug Setup](./xdebug_phpstorm.md)

* Disable xdebug when finish

  ```
  madock debug off
  ```

#### 6. SSL certificates

If you want to manually add an ssl certificate to the browser, you can find it at [path to madock folder]/aruntime/ctx/madockCA.pem
If the SSL certificates do not work, run the `madock ssl:rebuild` command and restart your browser.

#### 7. auth.json

If your project does not have an auth.json file, then when executing `composer` commands, the global auth.json file will be used.

#### 8. Multistores and website codes

Magento uses "base" as the store code by default.
But if you are using multistore, then you need to specify the code of each website along with the website host in the madock configuration. For example: `madock config:set --name=HOSTS --value="website1.test:base website2.test:websitecode"`. You can see site codes in the database table store_website. Or by querying the database `SELECT * FROM store_website`.

#### 9. help
```
  madock help
 ```

This command shows you the following items:

* `bash`    Connect into container using bash

  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`[name of container]` Name of container. Optional. Default container: php. For example: php, node, db, nginx

* `c:f`  Cleaning up static and generated files


* `cli`  Execute any commands inside php container. If you want to run several commands you can cover them in the quotes. For example: `madock cli "php bin/magento setup:upgrade && php bin/magento setup:di:compile"`


* `cloud`  Executing commands to work with Magento Cloud. Also, can be used the long command: magento-cloud)


* `composer`  Execute composer inside php container. For example: `madock composer install`
            
            
* `compress`  Compress a project to archive
            
            
* `config:list`  List all project environment settings


* `config:set`  Set a new value for parameter. For example: `madock config:set --name=HOSTS --value="website1.test:base website2.test:websitecode"`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--name`     Parameter name

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--value`     Parameter value
               
         
* `cron:enable`    Enable cron


* `cron:disable`    Disable cron
              
          
* `db:import`      Import database

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`-f`  Forse mode

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--service-name`  DB container name. Optional. Default container: db. Example: db2


* `db:export`      Export database. For example: `madock db:export --name=fromdevsite`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--name`  Name of the DB export file

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--service-name`  DB container name. Optional. Default container: db. Example: db2


* `db:info`      Information about credentials and remote host and port
 
   
* `debug:enable`   Enable xdebug


* `debug:disable`   Disable xdebug
                     

* `debug:profile:enable`   Enable xdebug profiling


* `debug:profile:disable`   Disable xdebug profiling
                     
   
* `info`   Show information about third-parties modules (name, current version, latest version, status)             
    
    
* `install`   Install Magento. It is a synonym for `madock magento setup:install` with additional actions.            
    
    
* `help`    Displays help for commands
                      
  
* `logs`    View logs of the container. For example: `madock logs php`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`[name of container]`     Container name. Optional. Default container: php. Example: php
                        

* `magento` or `m`   Execute Magento command inside php container. For example: `madock m setup:upgrade`
                        

* `n98`   Execute n98 command inside php container. For example: `madock n98 sys:info`
                        
                        

* `node`    Execute NodeJs command inside php container. For example: `madock node grunt exec:<theme>`
                        

* `patch:create`   Create patch. The patch can be used with the composer plugin cweagans/composer-patches. For example: `madock patch:create --file=vendor/magento/module-analytics/Cron/CollectData.php --name=collect-data-cron.patch --title="Collect data cron patch" --force`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--file`     Path of changed file. For example: vendor/magento/module-analytics/Cron/CollectData.php

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--name`     Name of the patch file

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--title`     Title of the patch

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--force`     Replace patch if it already exists


* `project:remove`   Remove project (project folder, madock project configuration, volumes, images, containers)

* `proxy:start`   Start a proxy server


* `proxy:stop`   Stop a proxy server


* `proxy:restart`   Restart a proxy server


* `proxy:rebuild`   Rebuild a proxy server


* `proxy:prune`   Prune a proxy server
                        

* `prune`   Stop and delete running project containers. For example: `madock prune --with-volumes`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--with-volumes`   Remove volumes, too
                        

* `rebuild` Recreation of all containers in the project. All containers are re-created and the images from the Dockerfile are rebuilt
                        

* `remote:sync:media`  Synchronization media files from remote host. For example: `madock remote:sync:media --images-only --compress`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--images-only`   Synchronization images only

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--compress`      Apply lossy compression. Images will have weight equals 30% of original

* `remote:sync:db`  Create and download dump of DB from remote host. For example: `madock remote:sync:db --name=local`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--name`  Name of the DB export file


* `remote:sync:file`  Create and download dump of DB from remote host

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--path`   Path to file on server (from Magento root)
                        

* `restart` Restarting all containers and services. Stop all containers and start them again
                        

* `service:list`   Show all services


* `service:enable`   Enable the service. For example: `madock service:enable phpmyadmin`

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`[service name]`  Service name
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--global`  Enable the service globally


* `service:disable`   Disable the service. For example: `madock service:disable phpmyadmin` 

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`[service name]`  Service name
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--global`  Disable the service globally
                        

* `setup`   Initial the project setup

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--download`   Download the specific Magento version from Composer to the container

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--install`   Install Magento from the source code

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--sample-data`   Install Magento Sample Data                    

* `setup:env`   Generate app/etc/env.php

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`-f`   Force re-create the file

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`--host`   Default host
                        

* `ssl:rebuild`   Rebuild SSL Certificates  
                        

* `start`   Starting all containers and services
                        

* `status`   Display the status of the project
                        

* `stop`    Stopping all containers and services
                        

* `uncompress`  Uncompress the project from archive