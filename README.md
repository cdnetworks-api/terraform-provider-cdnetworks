# Terraform Provider For Cdnetworks

[![stars](https://img.shields.io/github/stars/cdnetworks-api/terraform-provider-cdnetworks)](https://img.shields.io/github/stars/cdnetworks-api/terraform-provider-cdnetworks)
[![Forks](https://img.shields.io/github/forks/cdnetworks-api/terraform-provider-cdnetworks)](https://img.shields.io/github/forks/cdnetworks-api/terraform-provider-cdnetworks)
[![Go Report Card](https://goreportcard.com/badge/github.com/cdnetworks-api/terraform-provider-cdnetworks)](https://goreportcard.com/report/github.com/cdnetworks-api/terraform-provider-cdnetworks)
[![Releases](https://img.shields.io/github/release/cdnetworks-api/terraform-provider-cdnetworks.svg?style=flat-square)](https://github.com/cdnetworks-api/terraform-provider-cdnetworks/releases)
[![License](https://img.shields.io/github/license/cdnetworks-api/terraform-provider-cdnetworks)](https://img.shields.io/github/license/cdnetworks-api/terraform-provider-cdnetworks)
[![Issues](https://img.shields.io/github/issues/cdnetworks-api/terraform-provider-cdnetworks)](https://img.shields.io/github/issues/cdnetworks-api/terraform-provider-cdnetworks)

<div>
  <p>
    <a href="https://www.cdnetworks.com">
        <img src="https://www.cdnetworks.com/wp-content/uploads/2020/11/cdnetworks-logo-svg.svg" alt="logo" title="Terraform" width="300" height="45">
    </a>
    <br>
    <i>Cdnetworks Infrastructure for Terraform.</i>
    <br>
  </p>
</div>



* Tutorials: https://www.terraform.io

* [![Documentation](https://img.shields.io/badge/documentation-blue)](https://registry.terraform.io/providers/cdnetworks-api/cdnetworks/latest/docs)

* [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)

* Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

    

## Requirements

* [Terraform](https://www.terraform.io/downloads.html) 0.13.x
* [Go](https://golang.org/doc/install) 1.17.x (to build the provider plugin)

## Usage

### Build from source code

Clone repository to: `$GOPATH/src/github.com/cdnetworks-api/terraform-provider-cdnetworks`

```sh
$ mkdir -p $GOPATH/src/github.com/cdnetworks-api
$ cd $GOPATH/src/github.com/cdnetworks-api
$ git clone https://github.com/cdnetworks-api/terraform-provider-cdnetworks.git
$ cd terraform-provider-cdnetworks
$ go build .
```

If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it.

### Configure proxy info (optional)

If you are beind a proxy, for example, in a corporate network, you must set the proxy environment variables correctly. For example:

```
export http_proxy=http://your-proxy-host:your-proxy-port  # This is just an example, use your real proxy settings!
export https_proxy=$http_proxy
export HTTP_PROXY=$http_proxy
export HTTPS_PROXY=$http_proxy
```

## Run demo

You can edit your own terraform configuration files. Learn examples from examples directory.

Now you can try your terraform demo:

```
terraform init
terraform plan
terraform apply
```

If you want to destroy the resource, make sure the instance is already in ``running`` status, otherwise the destroy might fail.

```
terraform destroy
```

## Developer Guide

### DEBUG

You will need to set an environment variable named ``TF_LOG``, for more info please refer to [Terraform official doc](https://www.terraform.io/docs/internals/debugging.html):

```
export debuggable=true
```

In your source file, import the standard package ``log`` and print the message such as:

```
log.Println("[DEBUG] the message and some import values: %v", importantValues)

```

### License

Terraform-Provider-Cdnetworks is under the Mozilla Public License 2.0. See the [LICENSE](LICENSE.txt) file for details.